package ais

import (
	//"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/goburrow/serial"
)

type AISHandler struct {
	Port serial.Port
}

func (h *AISHandler) StartUDP(addr string) {
	udpAddr, _ := net.ResolveUDPAddr("udp", addr)
	conn, _ := net.ListenUDP("udp", udpAddr)
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, _, _ := conn.ReadFromUDP(buffer)
		msg := string(buffer[:n])
		h.parseAIS(msg)
	}
}

func (h *AISHandler) parseAIS(nmea string) {
	if strings.HasPrefix(nmea, "!AIVDM") {
		parts := strings.Split(nmea, ",")
		if len(parts) >= 7 {
			fmt.Printf("AIS Message: MMSI=%s, Type=%s\n", parts[5], parts[0])
		}
	}
}