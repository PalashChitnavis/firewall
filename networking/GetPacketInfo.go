package networking

import (
	"firewall/types"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

//Extract all info from the packet
func GetPacketInfo(packet gopacket.Packet , device pcap.Interface) (types.PacketInfo){
	var packetInfo types.PacketInfo 
	packetInfo.DeviceName = device.Description
	packetInfo.DeviceID = device.Name

	// Extract link layer information
	GetLinkLayerInfo(packet, &packetInfo)

	// Extract network layer information only if link layer information is available
	if packetInfo.LinkLayerType != "" {
		GetNetworkLayerInfo(packet, &packetInfo)
	}

	if packetInfo.NetworkLayerProtocol != "" {
		GetTransportLayerInfo(packet,&packetInfo)
	}

	if packetInfo.TransportProtocol != "" {
		GetApplicationLayerInfo(packet,&packetInfo)
	}

	return packetInfo
}