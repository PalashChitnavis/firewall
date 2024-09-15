package portprocess

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// ProcessInfo holds the process ID and process name
type ProcessInfo struct {
    PID  string
    Name string
}

// Global variable to store the mapping of port to ProcessInfo (PID and process name)
var PortProcessMap = make(map[int]ProcessInfo)
var mu sync.Mutex // Mutex to ensure safe concurrent access

// Initialize the map update process to run every 2 seconds
func Init() {
    go periodicUpdate()
}

// periodicUpdate runs every 2 seconds to refresh the port-to-process map
func periodicUpdate() {
    for {
        UpdatePortProcessMap()
        PrintMap()
        time.Sleep(2 * time.Second)
    }
}

// UpdatePortProcessMap updates the global map with port-to-process mappings
func UpdatePortProcessMap() {
    mu.Lock()
    defer mu.Unlock()

    // Clear existing map before updating
    PortProcessMap = make(map[int]ProcessInfo)

    // Execute the command to list all processes
    out, err := exec.Command("tasklist", "/FO", "CSV", "/NH").Output()
    if err != nil {
        fmt.Println("Error executing command:", err)
        return
    }

    // Process the output
    lines := strings.Split(string(out), "\n")
    for _, line := range lines {
        if len(line) > 0 {
            parts := strings.Split(line, ",")
            if len(parts) >= 2 {
                processName := strings.Trim(parts[0], "\"")
                pid := strings.Trim(parts[1], "\"")

                // For simplicity, assume a method to get all ports for a process
                ports := getAllPortsForProcess(pid)
                for _, port := range ports {
                    // Update the global map with ProcessInfo struct
                    PortProcessMap[port] = ProcessInfo{
                        PID:  pid,
                        Name: processName,
                    }
                }
            }
        }
    }
}

// getAllPortsForProcess retrieves all ports associated with a given process ID
func getAllPortsForProcess(pid string) []int {
    var ports []int

    // Execute the command to list all ports associated with the process ID
    out, err := exec.Command("netstat", "-ano").Output()
    if err != nil {
        fmt.Println("Error executing command:", err)
        return ports
    }

    lines := strings.Split(string(out), "\n")
    for _, line := range lines {
        parts := strings.Fields(line)
        if len(parts) >= 5 {
            // Extract the local address (in the form IP:PORT)
            localAddress := parts[1]
            // Extract the PID (process ID)
            currentPID := parts[len(parts)-1]

            if currentPID == pid {
                // Split the local address to get the port
                addressParts := strings.Split(localAddress, ":")
                if len(addressParts) == 2 {
                    port := addressParts[1]

                    // Convert port string to int
                    var portNumber int
                    fmt.Sscanf(port, "%d", &portNumber)
                    ports = append(ports, portNumber)
                }
            }
        }
    }
    return ports
}

// GetProcess returns the ProcessInfo (containing PID and process name) for a given port
func GetProcess(port int) (ProcessInfo, bool) {
    mu.Lock()
    defer mu.Unlock()

    processInfo, exists := PortProcessMap[port]
    return processInfo, exists
}

// GetAllProcesses returns the global map of all port-to-process mappings
func GetAllProcesses() map[int]ProcessInfo {
    mu.Lock()
    defer mu.Unlock()

    // Return a copy of the map to avoid modifying the original
    mapCopy := make(map[int]ProcessInfo)
    for port, process := range PortProcessMap {
        mapCopy[port] = process
    }
    return mapCopy
}

// PrintMap prints the current port-to-process mappings
func PrintMap() {
    mu.Lock()
    defer mu.Unlock()

    fmt.Println("Current Port to Process Mappings:")
    for port, processInfo := range PortProcessMap {
        fmt.Printf("Port: %d, PID: %s, Process: %s\n", port, processInfo.PID, processInfo.Name)
    }
}
