package dpi

/*
#cgo CFLAGS: -IC:/msys64/home/Palash/nDPI/src/include
#cgo LDFLAGS: -LC:/msys64/home/Palash/nDPI/src/lib -lndpi -lws2_32
#include <ndpi_main.h>
#include <ndpi_api.h>
#include <stdlib.h>
#include <string.h>
struct ndpi_detection_module_struct *initialize_ndpi() {
    struct ndpi_global_context *g_ctx = NULL;
    struct ndpi_detection_module_struct *ndpi_struct;
    ndpi_struct = ndpi_init_detection_module(g_ctx);
    if (ndpi_struct == NULL) {
        return NULL;
    }
    return ndpi_struct;
}

typedef unsigned char u_char;

const char* detect_protocol(struct ndpi_detection_module_struct *ndpi_struct, const u_char *packet) {
    // Calculate packet length from the input packet data
    unsigned short packetlen = strlen((const char *)packet); // For simplicity, assume the packet is a valid string
    // In case the packet is binary data, you may want to adjust this logic.
    // Initialize an empty flow structure
    struct ndpi_flow_struct *flow = (struct ndpi_flow_struct *)calloc(1, sizeof(struct ndpi_flow_struct));

    // Assume packet time is 0 (can be replaced with actual timestamp if necessary)
    u_int64_t packet_time_ms = 0;

    // Detect protocol using the nDPI function
    ndpi_protocol detected_protocol = ndpi_detection_process_packet(ndpi_struct, flow, packet, packetlen, packet_time_ms, NULL);

    // Extract the protocol ID from the detected protocol
    u_int16_t proto_id = detected_protocol.proto.master_protocol; // Use master protocol for the main application protocol

    // Free the flow structure
    free(flow);

     const char* proto_name = ndpi_get_proto_name(ndpi_struct, proto_id);

    // Print the protocol ID and name
    printf("Detected Protocol ID: %u\n", proto_id);
    printf("Detected Protocol Name: %s\n", proto_name);

    // Return the protocol name
    return proto_name;
}

// Convert char* to string
char* char_ptr_to_string(const char* cstr) {
    if (cstr == NULL) return NULL;
    size_t len = strlen(cstr);
    char* str = (char*)malloc(len + 1);
    if (str) {
        strcpy(str, cstr);
    }
    return str;
}


*/
import "C"
import (
	"unsafe"

	"github.com/google/gopacket"
)

type NDPIContext *C.struct_ndpi_detection_module_struct

// InitializeNDPI initializes the nDPI detection module and returns a pointer to it.
func InitializeNDPI() *C.struct_ndpi_detection_module_struct {
    return C.initialize_ndpi()
}

func PacketToCUChar(packet gopacket.Packet) (*C.u_char) {
	// Get the raw packet data as a byte slice
	packetData := packet.Data()

	// Convert the byte slice to a C pointer
	packetPtr := (*C.u_char)(unsafe.Pointer(&packetData[0]))

	// Return the C pointer and the length of the packet data
	return packetPtr
}

func DetectProtocol (ndpi_struct NDPIContext , packet gopacket.Packet) string{
	 
	packetPtr := PacketToCUChar(packet)
	protocolNameC := C.detect_protocol(ndpi_struct , packetPtr)
	s := C.GoString(protocolNameC)
	return s
}




