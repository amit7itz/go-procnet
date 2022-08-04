### Go-ProcNet
A go module for parsing linux socket files, supporting:
- /proc/net/tcp
- /proc/net/tcp6
- /proc/net/udp
- /proc/net/udp6

Can also be used for socket files of a specific process under `/proc/<PID>/net/`

Based on the source code of https://github.com/cakturk/go-netstat

### Installation:

```
$ go get github.com/amit7itz/go-procnet
```

### Usage:
```go
package main

import (
	"fmt"
	"github.com/amit7itz/go-procnet/procnet"
)

func main() {
	// parsing the tcp sockets from /proc/net/tcp
	socks, err := procnet.TCPSocks()
	if err != nil {
		panic(err)
	}
	for _, sock := range socks {
		fmt.Printf("local ip: %s local port: %d remote IP: %s remote port: %d state: %s",
			sock.LocalAddr.IP, sock.LocalAddr.Port, sock.RemoteAddr.IP, sock.RemoteAddr.Port, sock.State)
	}

	// for parsing the sockets from a custom path
	socks, err = procnet.SocksFromPath("/proc/1234/net/udp")
	// ...

	// for parsing the sockets from the textual content of a socket file
	socks, err = procnet.SocksFromText(` sl  local_address                         remote_address                        st tx_queue rx_queue tr tm->when retrnsmt   uid  timeout inode
	0: 00000000000000000000000000000000:1B58 00000000000000000000000000000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 854815644 1 0000000000000000 100 0 0 10 0
	1: 0000000000000000FFFF0000860EA8C0:1B58 0000000000000000FFFF0000E70FA8C0:87F6 01 00000000:00000000 00:00000000 00000000     0        0 854824433 1 0000000000000000 20 4 1 10 -1`)
	// ...
}
```
