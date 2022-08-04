package procnet

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)
import "github.com/stretchr/testify/require"

func TestSocksFromText(t *testing.T) {
	socks, err := SocksFromText(`  sl  local_address                         remote_address                        st tx_queue rx_queue tr tm->when retrnsmt   uid  timeout inode
   0: 00000000000000000000000000000000:1B58 00000000000000000000000000000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 854815644 1 0000000000000000 100 0 0 10 0
   1: 0000000000000000FFFF0000860EA8C0:1B58 0000000000000000FFFF0000E70FA8C0:87F6 01 00000000:00000000 00:00000000 00000000     0        0 854824433 1 0000000000000000 20 4 1 10 -1`)
	require.NoError(t, err)
	require.Len(t, socks, 2)
	require.Equal(t, ":::7000", socks[0].LocalAddr.String())
	require.Equal(t, ":::0", socks[0].RemoteAddr.String())
	require.Equal(t, Listen, socks[0].State)
	require.Equal(t, "192.168.14.134", socks[1].LocalAddr.IP.String())
	require.Equal(t, uint16(7000), socks[1].LocalAddr.Port)
	require.Equal(t, "192.168.15.231", socks[1].RemoteAddr.IP.String())
	require.Equal(t, uint16(34806), socks[1].RemoteAddr.Port)
	require.Equal(t, Established, socks[1].State)
}

func TestSocksFromFile(t *testing.T) {
	content := `  sl  local_address rem_address   st tx_queue rx_queue tr tm->when retrnsmt   uid  timeout inode
   0: 0100007F:A6B0 0100007F:1B58 06 00000000:00000000 03:000013B0 00000000     0        0 0 3 0000000000000000
   1: 0100007F:9E6E 0100007F:1B58 06 00000000:00000000 03:0000040F 00000000     0        0 0 3 0000000000000000
   2: 0100007F:A076 0100007F:1B58 06 00000000:00000000 03:000007FE 00000000     0        0 0 3 0000000000000000
   3: 0100007F:A070 0100007F:1B58 06 00000000:00000000 03:000007F9 00000000     0        0 0 3 0000000000000000
   4: 0100007F:A282 0100007F:1B58 06 00000000:00000000 03:00000BE0 00000000     0        0 0 3 0000000000000000
   5: 0100007F:9C5C 0100007F:1B58 06 00000000:00000000 03:00000027 00000000     0        0 0 3 0000000000000000
   6: 0100007F:9E74 0100007F:1B58 06 00000000:00000000 03:00000416 00000000     0        0 0 3 0000000000000000`
	file, err := ioutil.TempFile("/tmp/", "procnet-unittest")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	require.NoError(t, os.WriteFile(file.Name(), []byte(content), 0600))
	socks, err := SocksFromPath(file.Name())
	require.NoError(t, err)
	require.Len(t, socks, 7)
	for _, sock := range socks {
		require.Equal(t, "127.0.0.1", sock.LocalAddr.IP.String())
		require.True(t, sock.LocalAddr.Port > 0) // different for each socket
		require.Equal(t, "127.0.0.1", sock.RemoteAddr.IP.String())
		require.Equal(t, uint16(7000), sock.RemoteAddr.Port)
		require.Equal(t, TimeWait, sock.State)
	}
}
