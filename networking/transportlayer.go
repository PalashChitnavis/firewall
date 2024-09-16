package networking

import (
	"firewall/types"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// printNetworkLayerInfo extracts and prints network layer information from a packet
func GetTransportLayerInfo(packet gopacket.Packet , packetInfo *types.PacketInfo) {
	// Network Layer
	transportLayer := packet.TransportLayer()
	if transportLayer != nil {
		switch transportLayer.LayerType() {
		case layers.LayerTypeTCP:
			tcpPacket, _ := packet.Layer(layers.LayerTypeTCP).(*layers.TCP)
			packetInfo.TransportProtocol = types.TransportLayerTCP
			packetInfo.SourcePort = int(tcpPacket.SrcPort)
			packetInfo.DestinationPort = int(tcpPacket.DstPort)
		case layers.LayerTypeUDP:
			udpPacket, _ := packet.Layer(layers.LayerTypeUDP).(*layers.UDP)
			packetInfo.TransportProtocol = types.TransportLayerUDP
			packetInfo.SourcePort = int(udpPacket.SrcPort)
			packetInfo.DestinationPort = int(udpPacket.DstPort)
		}
	}
}
