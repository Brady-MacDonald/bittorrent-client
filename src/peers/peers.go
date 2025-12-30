package peers

import (
	"bytes"
	"net"

	"github.com/jackpal/bencode-go"
)

// Peer encodes connection information for a peer
type Peer struct {
	IP   net.IP
	Port uint16
}

// Unmarshal peers bencoded into struct
func Unmarshal(bencodePeers []byte) ([]Peer, error) {
	peers := []Peer{}

	bencodedReader := bytes.NewReader(bencodePeers)
	err := bencode.Unmarshal(bencodedReader, &peers)
	if err != nil {
		return peers, err
	}

	return peers, nil
}
