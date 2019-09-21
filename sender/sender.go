package sender

import (
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	senderHost = "127.0.0.1"
	senderPort = "7373"
	senderType = "udp"
	service    = senderHost + ":" + senderPort
)

// ListenForRequests : Bind to senderPort on senderHost and wait for incoming requests
func ListenForRequests() {
	fmt.Println("Binding to port ", senderPort, " on ", senderHost, " using ", senderType)
	if strings.ToUpper(senderType) == "TCP" {
		TCPListener()
	} else if strings.ToUpper(senderType) == "UDP" {
		UDPListener()
	}
}

// TCPListener : Set up service and listen for TCP requests
func TCPListener() {
	tcpAddr, err := net.ResolveTCPAddr(senderType, service)
	if err != nil {
		log.Fatal(err)
	}
	ln, err := net.ListenTCP(senderType, tcpAddr)
	if err != nil {
		log.Fatal("Error binding", err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Error reading packet: ", err)
		}
		go handleTCPRequest(conn)
	}
}

// UDPListener : Set up service and listen for UDP requests
func UDPListener() {
	udpAddr, err := net.ResolveUDPAddr(senderType, service)
	if err != nil {
		log.Fatal(err)
	}
	ln, err := net.ListenUDP(senderType, udpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	for {
		go handleUDPRequest(ln)
	}
}

// handleTCPRequest : TCP handle
func handleTCPRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	data, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("Received data stream: ", string(buffer[:data]))
}

// handleUDPRequest : UDP handle
func handleUDPRequest(conn *net.UDPConn) {
	buffer := make([]byte, 1024)
	data, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Received data stream: ", string(buffer[:data]))
}
