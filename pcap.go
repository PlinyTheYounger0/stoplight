package main

import (
	"fmt"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func openPcapReader(device string, promiscuous bool) error {
	fmt.Printf("%v: PCAP Reader Opened Successfully\n", time.Now())

	handle, err := pcap.OpenLive(device, int32(65535), promiscuous, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("Failed to Open PCAP Handler: %w\n", err)
	}
	defer handle.Close()

	fmt.Printf("%v: PCAP Handler Opened Successfully\n", time.Now())

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())


	for packet := range packetSource.Packets() {
		packetParser(packet)
	}

	return nil
}

func packetParser(packet gopacket.Packet) {
	if err := packet.ErrorLayer(); err != nil {
		fmt.Printf("Failed to Decode Packet: %v\n", err)
	}


	fmt.Printf("\n%s\n", packet.Metadata().Timestamp)

	for _, layer := range packet.Layers() {
		switch layer.LayerType() {

		case layers.LayerTypeEthernet:
			eth := layer.(*layers.Ethernet)
			fmt.Printf("Eth: %v -> %v\n", eth.SrcMAC, eth.DstMAC)

		case layers.LayerTypeIPv4:
			ipv4 := layer.(*layers.IPv4)
			fmt.Printf("IPv4: %v -> %v\n", ipv4.SrcIP, ipv4.DstIP)
		case layers.LayerTypeIPv6:
			ipv6 := layer.(*layers.IPv6)
			fmt.Printf("IPv6: %v -> %v\n", ipv6.SrcIP, ipv6.DstIP)

		case layers.LayerTypeTCP:
			tcp := layer.(*layers.TCP)
			fmt.Printf("TCP: %v -> %v\n", tcp.SrcPort, tcp.DstPort)
		case layers.LayerTypeUDP:
			udp := layer.(*layers.UDP)
			fmt.Printf("UDP: %v -> %v\n", udp.SrcPort, udp.DstPort)

		case layers.LayerTypeNTP:
			ntp := layer.(*layers.NTP)
			fmt.Println("NTP:")
			fmt.Printf("	Vesion: %d\n", ntp.Version)
			fmt.Printf("	Mode: %d\n", ntp.Mode)
			fmt.Printf("	Stratum: %d\n", ntp.Stratum)

		case layers.LayerTypeICMPv4:
			icmpv4 := layer.(*layers.ICMPv4)
			fmt.Println("ICMPv4:")
			fmt.Printf("	Typecode: %v\n", icmpv4.TypeCode)
			fmt.Printf("	Checksum: %v\n", icmpv4.Checksum)
			fmt.Printf("	ID: %v\n", icmpv4.Id)
			fmt.Printf("	Seq: %v\n", icmpv4.Seq)

		case layers.LayerTypeICMPv6:
			imcpv6 := layer.(*layers.ICMPv6)
			fmt.Println("ICMPv6: ")
			fmt.Printf("	Checksum: %d\n", imcpv6.Checksum)
		case layers.LayerTypeICMPv6NeighborSolicitation:
			icmpv6NS := layer.(*layers.ICMPv6NeighborSolicitation)
			fmt.Println("ICMPv6 Neighbor Solicitation: ")
			fmt.Printf("	Target Address: %v\n", icmpv6NS.TargetAddress)
			for _, option := range icmpv6NS.Options{
				if option.Type.String() == "SourceAddress" {
					srcMAC := net.HardwareAddr(option.Data)
					fmt.Printf("	Source Address: %v\n", srcMAC)
				} else {
					fmt.Printf("	%s: %v\n", option.Type.String(), option.Data)
				}
			}

//		case gopacket.LayerTypePayload:
//			app := layer.(*gopacket.Payload)
//			fmt.Printf("Application Data: %s\n", string(app.Payload()))

		default:
			fmt.Printf("Layer: %v Not Implemented\n", layer.LayerType())
		}
	}
}

