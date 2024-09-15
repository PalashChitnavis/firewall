package networking

import (
	"firewall/types"

	"github.com/google/gopacket"
)

// printNetworkLayerInfo extracts and prints network layer information from a packet
func GetApplicationLayerInfo(packet gopacket.Packet , packetInfo *types.PacketInfo) {
	// Network Layer
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		packetInfo.Payload = applicationLayer.Payload()
		// perform deep inspection using dpi to get application layer protocol
	}
}
