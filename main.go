package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {

	//Timestamp
	//timestamp := time.Now()
	//formattedtimestamp := timestamp.Format("2006-01-01 15:04:05") // YYYY-MM-DD HH:MM:SS format
	//year, month, day, time := timestamp.Date() // Extracting year, month, and day from the timestamp
	// Define the full file path

	//Create the dir for the log files
	dirPath := "C:\\logfiles"
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return
	}
	fmt.Println("Directory Logfiles created, or already exists.")

	filePath := filepath.Join(dirPath, "ipv4.txt")
	//file, _ := os.Create(filePath)
	// Open the file in append mode, or create it if it doesn’t exist
	//file.Close()
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	fmt.Println("IPV4 log file created!")
	//Get machines IPV4 information
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			fmt.Println("IP Addresses", ipNet.IP.String())
			file.WriteString(ipNet.IP.String() + "\n")

		}
	}
	fmt.Println("No IP address found")

	//Get interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting interfaces:", err)
		return
	}

	for _, iface := range interfaces {
		// Print the MAC (Hardware) Address
		fmt.Printf("Interface: %s, MAC Address: %s\n", iface.Name, iface.HardwareAddr.String())
		// Format the data properly and write it to the file
		_, err := file.WriteString(fmt.Sprintf("Interface: %s, MAC Address: %s\n", iface.Name, iface.HardwareAddr.String()))
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}

	iplist, _ := net.Interfaces()
	for index, ip := range iplist {
		fmt.Println(index)
		fmt.Println(ip.Addrs())
		addString, _ := ip.Addrs()

		for _, id := range addString {
			mystring := id.String()
			file.WriteString(mystring)
		}

		addrlist, _ := ip.Addrs()

		for _, golist := range addrlist {
			fmt.Println(golist)
		}

	}

	file.Close()

}
