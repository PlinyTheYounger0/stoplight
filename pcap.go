package main

import (
	"fmt"
	//"log"
	"time"

	"github.com/google/gopacket"

	//"github.com/google/gopacket/layers"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type PacketInfo struct {
	timeStamp time.Time
	truncated bool
	srcIP string
	dstIP string
	srcMAC string
	dstMAC string
}

func openPcapReader(device string, promiscuous bool) error {
	fmt.Printf("%v: PCAP Reader Opened Successfully\n", time.Now())

	handle, err := pcap.OpenLive(device, int32(65535), promiscuous, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("Failed to Open PCAP Handler: %w\n", err)
	}
	defer handle.Close()

	fmt.Printf("%v: PCAP Handler Opened Successfully\n", time.Now())

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	err = packetParser(packetSource)
	if err != nil {
		return fmt.Errorf("Error Parsing Packets - %w\n", err)
	}

	return nil
}

func packetParser(packetSource *gopacket.PacketSource) error {
	fmt.Printf("%v: Packet Parser Began\n", time.Now())
	defer fmt.Printf("%v: Packet Parser Closed Successfully\n", time.Now())
	for packet := range packetSource.Packets() {
		if err := packet.ErrorLayer(); err != nil {
			return fmt.Errorf("Failed to Decode Packet: %v", err)
		}
	
		for _, layer := range packet.Layers() {
			switch layer.LayerType() {
			case layers.LayerTypeEthernet:
				eth := layer.(*layers.Ethernet)
				fmt.Printf("Layer 2: %v -> %v\n", eth.SrcMAC, eth.DstMAC)
			case layers.LayerTypeIPv4:
				ip := layer.(*layers.IPv4)
				fmt.Printf("Layer 3: %v -> %v\n", ip.SrcIP, ip.DstIP)
			case layers.LayerTypeTCP:
				tcp := layer.(*layers.TCP)
				fmt.Printf("Layer 4: %v -> %v\n", tcp.SrcPort, tcp.DstPort)
			case gopacket.LayerTypePayload:
				app := layer.(*gopacket.Payload)
				fmt.Println(string(app.Payload()))
			default:
				fmt.Printf("Layer: %v Not Implemented", layer)
			}
		}
	}

	return nil
}
