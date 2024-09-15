package networking

import (
	"firewall/types"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// Extracts and prints link layer information from a gopacket.Packet
func GetLinkLayerInfo(packet gopacket.Packet, packetInfo *types.PacketInfo) {
    // Get the link layer (Ethernet) from the packet
    linkLayer := packet.LinkLayer()
    if linkLayer == nil {
        return
    }

    // Check if the link layer is Ethernet
    if ethernetLayer, ok := packet.Layer(layers.LayerTypeEthernet).(*layers.Ethernet); ok {
        packetInfo.SourceMAC = ethernetLayer.SrcMAC.String()
        packetInfo.DestinationMAC = ethernetLayer.DstMAC.String()
        packetInfo.LinkLayerType = types.LinkLayerEthernet
    }

    // Check if the link layer is WiFi (802.11)
    // if wifiLayer, ok := packet.Layer(layers.LayerTypeDot11).(*layers.Dot11) ; ok {
    //     packetInfo.SourceMAC = wifiLayer.Address1.String() // Adapt as needed
    //     packetInfo.DestinationMAC = wifiLayer.Address2.String() // Adapt as needed
    //     packetInfo.LinkLayerType = types.LinkLayerWifi
    //     return
    // }
}
