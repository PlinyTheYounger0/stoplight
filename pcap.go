package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type PacketInfo struct {
	captureTime time.Time
	truncated bool
	srcIP string
	dstIP string
	srcMAC string
	dstMAC string
}

func openPcapReader(device string, promiscuous bool) error {
	handle, err := pcap.OpenLive(device, int32(65535), promiscuous, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("Failed to Open PCAP Handler: %w", err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	err = packetParser(packetSource)
	if err != nil {
		return fmt.Errorf("Error Parsing Packets - %w", err)
	}

	return nil
}

func packetParser(packetSource *gopacket.PacketSource) error {
	for packet := range packetSource.Packets() {
		if err := packet.ErrorLayer(); err != nil {
			return fmt.Errorf("Failed to Decode Packet: %v", err)
		}

		packetInfo := PacketInfo{}
		app := packet.ApplicationLayer()
		if app != nil {
			log.Printf("Application Layer Data: %v", app.Payload())
		}

		network := packet.NetworkLayer()
		if network != nil {
			packetInfo.srcIP = network.NetworkFlow().Src().String()
			packetInfo.dstIP = network.NetworkFlow().Dst().String()
		} else {
			packetInfo.srcIP = ""
			packetInfo.dstIP = ""
		}

		link := packet.LinkLayer()
		if link != nil {
			packetInfo.srcMAC = link.LinkFlow().Src().String()
			packetInfo.dstMAC = link. LinkFlow().Dst().String()
		} else {
			packetInfo.srcMAC = ""
			packetInfo.dstMAC = ""
		}

		log.Printf(
			"Soruce IP: %s\nDestination IP: %s\nSource MAC: %s\nDestination MAC: %s\n",
			packetInfo.srcIP,
			packetInfo.dstIP,
			packetInfo.srcMAC,
			packetInfo.dstMAC,
		)
	}

	return nil
}
