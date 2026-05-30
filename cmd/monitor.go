/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
)

var verbose bool
var promiscuous bool

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Begin monitoring traffic on your network via your network interface",
	Long: `The monitor command allows the user to begin a monioring session on their network interface.

	The monitoring session is started by the user who is currently logged in and will be logged as 
	`,

	Run: func(cmd *cobra.Command, args []string) {

	if verbose {
		fmt.Println("monitor called")
		fmt.Printf("%v: PCAP Reader Opened Successfully\n", time.Now())
	}

	device := "wlx00c0cab5102c"

	handle, err := pcap.OpenLive(device, int32(65535), promiscuous, pcap.BlockForever)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to Open PCAP Reader: %v", err)
		os.Exit(1)
	}
	defer handle.Close()

	if verbose {
		fmt.Printf("%v: PCAP Handler Opened Successfully\n", time.Now())
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		packetParser(packet)
	}

	},
}

func init() {
	monitorCmd.Flags().BoolP("promiscuous", "p", false, "enable promiscuous mode for the network interface")
	monitorCmd.Flags().BoolP("verbose", "v", false, "enable verbose output")
	rootCmd.AddCommand(monitorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// monitorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// monitorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

