package service

import (
	"net"
	"github.com/younisshah/gelffy/util"
)

// StartUDPServer starts UDP server on the given port
// Return a UDP connection
func StartUDPServer(port string) *net.UDPConn {

	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	util.CheckError(err)

	conn, err := net.ListenUDP("udp", addr)
	util.CheckError(err)

	return conn
}
