package tracker

import (
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Brady-MacDonald/bittorrent-client/src/metadata"
	"github.com/Brady-MacDonald/bittorrent-client/src/peers"
	"github.com/jackpal/bencode-go"
)

// Talk to the tracker
// Central location to connect with peers

type bencodeTrackerResp struct {
	IP   net.IP
	Port uint16
}

// Make http Get request to Tracker to receive list of peers
func GetPeers(url string) (*[]peers.Peer, error) {
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	trackerResp := &bencodeTrackerResp{}
	err = bencode.Unmarshal(resp.Body, trackerResp)
	if err != nil {
		return nil, err
	}

	// peers.Unmarshal()
	// return peers.Unmarshal([]byte(trackerResp.Peers))
	return nil, nil
}

// Based on the torrents Announce field
// Build URL to make request to Tracker
func BuildTrackerURL(torrent *metadata.Torrent, peerID [20]byte) (string, error) {
	base, err := url.Parse(torrent.Announce)
	if err != nil {
		return "", err
	}

	params := url.Values{
		"info_hash":  []string{string(torrent.Info.PieceLength)},
		"peer_id":    []string{string(peerID[:])},
		"port":       []string{strconv.Itoa(int(6969))},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{strconv.Itoa(torrent.Info.Length)},
	}

	base.RawQuery = params.Encode()
	return base.String(), nil
}
