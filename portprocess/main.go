package portprocess

import (
	"log"

	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

func PrintNetworkConnections() {
	connections, err := net.Connections("all") // "all" includes all types of connections: TCP, UDP, etc.
	if err != nil {
		log.Fatalf("Error retrieving network connections: %v", err)
	}

	for _, conn := range connections {
		log.Printf("PID: %d, Local Address: %s, Remote Address: %s, Status: %s",
			conn.Pid, conn.Laddr, conn.Raddr, conn.Status)
	}
}

func GetPIDByPort(port int) int32 {
	connections, err := net.Connections("all")
	if err != nil {
		log.Fatalf("Error retrieving network connections: %v", err)
		return -1
	}

	for _, conn := range connections {
		// Check if the connection is using the specified port
		if conn.Laddr.Port == uint32(port) {
			//fmt.Println(conn.Pid)
			return conn.Pid
		}
	}

	// Return -1 if no process was found using the specified port
	return -1
}

func PrintProcessInfo(pid int32) string {
    proc, err := process.NewProcess(pid)
    if err != nil {
        //log.Printf("Error retrieving process info for PID %d: %v", pid, err)
        return ""
    }

    name, err := proc.Name()
    if err != nil {
        //log.Printf("Error retrieving process name for PID %d: %v", pid, err)
        return ""
    }
	return name
    //log.Printf("Process PID: %d, Name: %s", pid, name)
}