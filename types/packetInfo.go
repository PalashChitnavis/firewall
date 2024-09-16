package types

import "fmt"

// PrintPacketInfo prints the PacketInfo struct in a structured format
func PrintPacketInfo(p PacketInfo) {
	fmt.Println("========== Packet Information ==========")

	// Device Info
	fmt.Println("Device Information:")
	fmt.Printf("  Device Name: %s\n", p.DeviceName)
	fmt.Printf("  Device ID: %s\n", p.DeviceID)

	// Link Layer Info
	fmt.Println("\nLink Layer Information:")
	fmt.Printf("  Source MAC: %s\n", p.SourceMAC)
	fmt.Printf("  Destination MAC: %s\n", p.DestinationMAC)
	fmt.Printf("  Link Layer Type: %s\n", p.LinkLayerType)

	// Network Layer Info
	fmt.Println("\nNetwork Layer Information:")
	fmt.Printf("  Source IP: %s\n", p.SourceIP)
	fmt.Printf("  Destination IP: %s\n", p.DestinationIP)
	fmt.Printf("  Network Layer Protocol: %s\n", p.NetworkLayerProtocol)

	// Transport Layer Info
	fmt.Println("\nTransport Layer Information:")
	fmt.Printf("  Source Port: %d\n", p.SourcePort)
	fmt.Printf("  Destination Port: %d\n", p.DestinationPort)
	fmt.Printf("  Transport Protocol: %s\n", p.TransportProtocol)

	// Application Layer Info
	fmt.Println("\nApplication Layer Information:")
	fmt.Printf("  Application Protocol: %s\n", p.ApplicationProtocol)
	fmt.Printf("  Payload: %s\n", string(p.Payload))

	// Process Information
	fmt.Println("\nProcess Information:")
	fmt.Printf("  Process Name: %s\n", p.ProcessName)
	fmt.Printf("  Process PID: %d\n", p.ProcessPID)

	// Additional Information
	fmt.Println("\nAdditional Information:")
	fmt.Printf("  Packet Size: %d bytes\n", p.PacketSize)
	fmt.Printf("  Timestamp: %s\n", p.Timestamp)
	fmt.Printf("  Direction: %s\n", p.Direction)

	fmt.Println("=========================================")
}

type PacketInfo struct {

	// Device Info
	DeviceName string
	DeviceID   string

	// Link Layer
	SourceMAC      string
	DestinationMAC string
	LinkLayerType  string // ethernet/wifi

	// Network Layer
	SourceIP             string
	DestinationIP        string
	NetworkLayerProtocol string // ip/arp/icmp

	// Transport Layer
	SourcePort        int
	DestinationPort   int
	TransportProtocol string // tcp/udp

	// Application Layer
	Payload             []byte
	ApplicationProtocol string // http/ftp/smtp

	// Process Information
	ProcessName string
	ProcessPID  int32

	// Additional Information
	PacketSize int
	Timestamp  string
	Direction  string // incoming/outgoing
}