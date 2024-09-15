package types

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
	ProcessPID  string

	// Additional Information
	PacketSize int
	Timestamp  string
	Direction  string // incoming/outgoing
}