package sip

import (
	"net"
)

type Server struct {
	Address string
}

func NewServer(addr string) *Server {
	return &Server{Address: addr}
}

func (s *Server) Start() error {
	addr, _ := net.ResolveUDPAddr("udp", s.Address)
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, addr, _ := conn.ReadFromUDP(buf)
		go s.handlePacket(buf[:n], addr)
	}
}

func (s *Server) handlePacket(data []byte, addr *net.UDPAddr) {
	// Базовая обработка SIP-пакетов
}