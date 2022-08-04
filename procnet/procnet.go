package procnet

import (
	"fmt"
	"net"
	"strings"
)

const (
	pathTCPTab  = "/proc/net/tcp"
	pathTCP6Tab = "/proc/net/tcp6"
	pathUDPTab  = "/proc/net/udp"
	pathUDP6Tab = "/proc/net/udp6"
)

// SockAddr represents an ip:port pair
type SockAddr struct {
	IP   net.IP
	Port uint16
}

func (s *SockAddr) String() string {
	return fmt.Sprintf("%v:%d", s.IP, s.Port)
}

// SockTabEntry type represents each line of the /proc/net/[tcp|udp]
type SockTabEntry struct {
	inode      string
	LocalAddr  *SockAddr
	RemoteAddr *SockAddr
	State      SkState
	UID        uint32
}

// SkState type represents socket connection state
type SkState uint8

func (s SkState) String() string {
	return skStates[s]
}

// TCPSocks returns a slice of active TCP sockets
func TCPSocks() ([]SockTabEntry, error) {
	return SocksFromPath(pathTCPTab)
}

// TCP6Socks returns a slice of active TCP IPv4 socket
func TCP6Socks() ([]SockTabEntry, error) {
	return SocksFromPath(pathTCP6Tab)
}

// UDPSocks returns a slice of active UDP sockets
func UDPSocks() ([]SockTabEntry, error) {
	return SocksFromPath(pathUDPTab)
}

// UDP6Socks returns a slice of active UDP IPv6 sockets
func UDP6Socks() ([]SockTabEntry, error) {
	return SocksFromPath(pathUDP6Tab)
}

// SocksFromPath returns a slice of active sockets described in the given file
func SocksFromPath(path string) ([]SockTabEntry, error) {
	return parseSocktabFromPath(path)
}

// SocksFromText returns a slice of active sockets described in the given text
func SocksFromText(text string) ([]SockTabEntry, error) {
	return parseSocktab(strings.NewReader(text))
}
