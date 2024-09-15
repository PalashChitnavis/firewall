package main

import (
	"firewall/networking"
	"firewall/portprocess"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Global variables
var macAddressSet map[string]struct{}
var localAddresses []string

func main() {
	portprocess.Init()

	// Initialize the MAC address set
	macAddressSet = make(map[string]struct{})

	// Initialize local IP addresses and ports
	//initializeLocalAddressesAndPorts()

	// Find all network devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	if len(devices) == 0 {
		log.Fatal("No devices found.")
	}

	// WaitGroup to handle multiple goroutines
	var wg sync.WaitGroup

	// Start a goroutine to listen on each device
	for _, device := range devices {
		fmt.Printf("Starting capture on device: %s\n", device.Description)
		wg.Add(1)
		go func(deviceName string, deviceDesc string) {
			defer wg.Done()
			capturePackets(device)
		}(device.Name, device.Description)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

// Initialize local IP addresses and ports
// func initializeLocalAddressesAndPorts() {
// 	interfaces, err := net.Interfaces()
// 	if err != nil {
// 		fmt.Println("Error retrieving network interfaces:", err)
// 		return
// 	}

// 	for _, iface := range interfaces {
// 		if len(iface.HardwareAddr) > 0 {
// 			macAddress := iface.HardwareAddr.String()
// 			macAddressSet[macAddress] = struct{}{}
// 		}
// 	}

// 	fmt.Println("MAC address set:", macAddressSet)

// 	// Initialize local IP addresses
// 	for _, iface := range interfaces {
// 		addrs, _ := iface.Addrs()
// 		for _, addr := range addrs {
// 			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
// 				localAddresses = append(localAddresses, ipnet.IP.String())
// 			}
// 		}
// 	}
// }

// Function to capture packets on a given device for 10 seconds
func capturePackets(device pcap.Interface) {
	handle, err := pcap.OpenLive(device.Name, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Printf("Error opening device %s: %v\n", device.Description, err)
		return
	}
	defer handle.Close()

	// Create a packet source to process packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	fmt.Printf("Listening for packets on %s...\n", device.Description)

	// Channel to signal when to stop capturing
	stop := time.After(10 * time.Second)

	// Capture and print packet details
	for {
		select {
		case packet := <-packetSource.Packets():
			//printPacketInfo(deviceDesc, packet)
			networking.GetPacketInfo(packet,device)
		case <-stop:
			fmt.Printf("Stopping capture on device: %s\n", device.Description)
			return
		}
	}
}

// Function to print packet details
// func printPacketInfo(deviceDesc string, packet gopacket.Packet) {
// 	fmt.Printf("Device: %s\n", deviceDesc)
// 	fmt.Println("----- Packet Details -----")

// 	// Determine packet direction
// 	direction := checkDirection(packet)
// 	fmt.Println("Packet direction:", direction)

// 	// Ethernet Layer
// 	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
// 	if ethernetLayer != nil {
// 		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
// 		srcMAC := ethernetPacket.SrcMAC.String()
// 		dstMAC := ethernetPacket.DstMAC.String()

// 		fmt.Printf("Source MAC: %s\n", srcMAC)
// 		fmt.Printf("Destination MAC: %s\n", dstMAC)
// 	}

// 	// Network Layer
// 	networkLayer := packet.NetworkLayer()
// 	if networkLayer != nil {
// 		switch networkLayer.LayerType() {
// 		case layers.LayerTypeIPv4:
// 			ipPacket, _ := networkLayer.(*layers.IPv4)
// 			fmt.Printf("Source IP: %s\n", ipPacket.SrcIP)
// 			fmt.Printf("Destination IP: %s\n", ipPacket.DstIP)
// 			fmt.Printf("Network Layer Protocol: IPv4\n")
// 		case layers.LayerTypeIPv6:
// 			ipPacket, _ := networkLayer.(*layers.IPv6)
// 			fmt.Printf("Source IP: %s\n", ipPacket.SrcIP)
// 			fmt.Printf("Destination IP: %s\n", ipPacket.DstIP)
// 			fmt.Printf("Network Layer Protocol: IPv6\n")
// 		case layers.LayerTypeARP:
// 			arpPacket, _ := packet.Layer(layers.LayerTypeARP).(*layers.ARP)
// 			fmt.Printf("Source MAC: %s\n", arpPacket.SourceHwAddress)
// 			fmt.Printf("Source IP: %s\n", arpPacket.SourceProtAddress)
// 			fmt.Printf("Destination MAC: %s\n", arpPacket.DstHwAddress)
// 			fmt.Printf("Destination IP: %s\n", arpPacket.DstProtAddress)
// 			fmt.Printf("Network Layer Protocol: ARP\n")
// 		case layers.LayerTypeICMPv4:
// 			icmpPacket, _ := packet.Layer(layers.LayerTypeICMPv4).(*layers.ICMPv4)
// 			fmt.Printf("ICMPv4 Type: %d\n", icmpPacket.TypeCode.Type())
// 			fmt.Printf("ICMPv4 Code: %d\n", icmpPacket.TypeCode.Code())
// 			fmt.Printf("Network Layer Protocol: ICMPv4\n")
// 		case layers.LayerTypeICMPv6:
// 			icmpPacket, _ := packet.Layer(layers.LayerTypeICMPv6).(*layers.ICMPv6)
// 			fmt.Printf("ICMPv6 Type: %d\n", icmpPacket.TypeCode.Type())
// 			fmt.Printf("ICMPv6 Code: %d\n", icmpPacket.TypeCode.Code())
// 			fmt.Printf("Network Layer Protocol: ICMPv6\n")
// 		default:
// 			fmt.Printf("Network Layer Protocol Default Case %s\n", networkLayer.LayerType())
// 		}
// 	}

// 	//var port int

// 	// TCP Layer
// 	tcpLayer := packet.Layer(layers.LayerTypeTCP)
// 	if tcpLayer != nil {
// 		tcpPacket, _ := tcpLayer.(*layers.TCP)
// 		// if direction == "Incoming" {
// 		// 	port = int(tcpPacket.DstPort)
// 		// } else if direction == "Outgoing" {
// 		// 	port = int(tcpPacket.SrcPort)
// 		// } else {
// 		// 	port = -1
// 		// }
// 		fmt.Printf("Source Port: %d\n", tcpPacket.SrcPort)
// 		fmt.Printf("Destination Port: %d\n", tcpPacket.DstPort)
// 		fmt.Printf("Transport Layer Protocol: TCP\n")
// 	}

// 	// UDP Layer
// 	udpLayer := packet.Layer(layers.LayerTypeUDP)
// 	if udpLayer != nil {
// 		udpPacket, _ := udpLayer.(*layers.UDP)
// 		// if direction == "Incoming" {
// 		// 	port = int(udpPacket.DstPort)
// 		// } else if direction == "Outgoing" {
// 		// 	port = int(udpPacket.SrcPort)
// 		// } else {
// 		// 	port = -1
// 		// }
// 		fmt.Printf("Source Port: %d\n", udpPacket.SrcPort)
// 		fmt.Printf("Destination Port: %d\n", udpPacket.DstPort)
// 		fmt.Printf("Transport Layer Protocol: UDP\n")
// 	}

// 	// if direction == "Incoming" {
// 	// 	fmt.Println("Packet is incoming")
// 	// } else if direction == "Outgoing" {
// 	// 	fmt.Println("Packet is outgoing")
// 	// } else {
// 	// 	fmt.Println("Packet is fucked -------------------------------------------------------")
// 	// }

// 	// if port != -1 {
// 	// 	processInfo, found := portprocess.GetProcess(port)
// 	// 	if found {
// 	// 		fmt.Printf("Port: %d\nProcess Name: %s\nProcess PID: %s\n", port, processInfo.Name, processInfo.PID)
// 	// 	} else {
// 	// 		fmt.Printf("Port: %d is not associated with any process\n", port)
// 	// 	}
// 	// } else {
// 	// 	fmt.Println("Port is mapped to -1, please check")
// 	// }

// 	fmt.Println("--------------------------")
// }

// // Function to determine if a packet is incoming or outgoing
// func checkDirection(packet gopacket.Packet) string {
// 	var incoming bool
// 	var outgoing bool

// 	// Ethernet Layer
// 	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
// 	if ethernetLayer != nil {
// 		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
// 		srcMAC := ethernetPacket.SrcMAC.String()
// 		dstMAC := ethernetPacket.DstMAC.String()

// 		if _, found := macAddressSet[srcMAC]; found {
// 			outgoing = true
// 		}
// 		if _, found := macAddressSet[dstMAC]; found {
// 			incoming = true
// 		}
// 	}

// 	if incoming && !outgoing {
// 		return "Incoming"
// 	}
// 	if outgoing && !incoming {
// 		return "Outgoing"
// 	}
// 	return "Unknown"
// }
