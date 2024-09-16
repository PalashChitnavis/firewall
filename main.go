package main

import (
	"firewall/networking"
	"firewall/portprocess"
	"firewall/types"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Initialize the process-port mapping system
	// portprocess.Init()

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

	// Start a goroutine to listen on each device for both incoming and outgoing packets
	for _, device := range devices {
		fmt.Printf("Starting capture on device: %s\n", device.Description)
		wg.Add(2) // Adding 2 for incoming and outgoing packet capture

		// Goroutine for capturing incoming packets
		go func(device pcap.Interface) {
			defer wg.Done()
			capturePackets(device, pcap.DirectionIn)
		}(device)

		// Goroutine for capturing outgoing packets
		go func(device pcap.Interface) {
			defer wg.Done()
			capturePackets(device, pcap.DirectionOut)
		}(device)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

// Function to capture packets on a given device for a specific direction (incoming/outgoing)
func capturePackets(device pcap.Interface, direction pcap.Direction) {
	// Open device for live capture
	handle, err := pcap.OpenLive(device.Name, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Printf("Error opening device %s: %v\n", device.Description, err)
		return
	}
	defer handle.Close()

	// Set the capture direction (incoming/outgoing)
	handle.SetDirection(direction)

	// Create a packet source to process packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Determine capture type based on direction
	captureType := types.PacketDirectionIn
	if direction == pcap.DirectionOut {
		captureType = types.PacketDirectionOut
	}

	//fmt.Printf("Listening for %s packets on %s...\n", captureType, device.Description)

	// Channel to signal when to stop capturing (after 10 seconds)
	stop := time.After(5 * time.Second)

	// Capture and print packet details
	for {
		select {
		case packet := <-packetSource.Packets():
			// Process packet
			packetInfo := networking.GetPacketInfo(packet, device)
			// ##TODO
			if packetInfo.LinkLayerType != "" && packetInfo.NetworkLayerProtocol != "" && packetInfo.TransportProtocol != "" {
				packetInfo.Direction = captureType // Set direction as incoming or outgoing
				if packetInfo.SourcePort != 0 && packetInfo.DestinationPort != 0 {
					if captureType == types.PacketDirectionIn {
						packetInfo.ProcessPID = portprocess.GetPIDByPort(packetInfo.DestinationPort)
						packetInfo.ProcessName = portprocess.PrintProcessInfo(packetInfo.ProcessPID)
					}else{
						packetInfo.ProcessPID = portprocess.GetPIDByPort(packetInfo.SourcePort)
						packetInfo.ProcessName = portprocess.PrintProcessInfo(packetInfo.ProcessPID)
					}
				}
				types.PrintPacketInfo(packetInfo)  // Print packet info
			}

		case <-stop:
			fmt.Printf("Stopping %s packet capture on device: %s\n", captureType, device.Description)
			return
		}
	}
}