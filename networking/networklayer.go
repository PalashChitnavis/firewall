package networking

import (
	"firewall/types"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// printNetworkLayerInfo extracts and prints network layer information from a packet
func GetNetworkLayerInfo(packet gopacket.Packet , packetInfo *types.PacketInfo) {
	// Network Layer
	networkLayer := packet.NetworkLayer()
	if networkLayer != nil {
		switch networkLayer.LayerType() {
		case layers.LayerTypeIPv4:
			ipPacket, _ := packet.Layer(layers.LayerTypeIPv4).(*layers.IPv4)
			packetInfo.NetworkLayerProtocol = types.NetworkLayerIPv4
			packetInfo.SourceIP = ipPacket.SrcIP.String()
			packetInfo.DestinationIP = ipPacket.DstIP.String()
		case layers.LayerTypeIPv6:
			ipPacket, _ := packet.Layer(layers.LayerTypeIPv6).(*layers.IPv6)
			packetInfo.NetworkLayerProtocol = types.NetworkLayerIPv6
			packetInfo.SourceIP = ipPacket.SrcIP.String()
			packetInfo.DestinationIP = ipPacket.DstIP.String()
		// case layers.LayerTypeARP:
		// 	arpPacket, _ := packet.Layer(layers.LayerTypeARP).(*layers.ARP)
		// 	packetInfo.NetworkLayerProtocol = types.NetworkLayerARP
		// 	packetInfo.SourceIP = string(arpPacket.SourceProtAddress)
		// 	packetInfo.DestinationIP = arpPacket.DstProtAddress
		// case layers.LayerTypeICMPv4:
		// 	//icmpPacket, _ := packet.Layer(layers.LayerTypeICMPv4).(*layers.ICMPv4)
		// 	packetInfo.NetworkLayerProtocol = types.NetworkLayerICMPv4
		// case layers.LayerTypeICMPv6:
		// 	//icmpPacket, _ := packet.Layer(layers.LayerTypeICMPv6).(*layers.ICMPv6)
		// 	packetInfo.NetworkLayerProtocol = types.NetworkLayerICMPv6
		}
	}
}
