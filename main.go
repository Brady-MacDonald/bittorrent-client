package main

import (
	"crypto/rand"
	"os"

	"github.com/Brady-MacDonald/bittorrent-client/src/metadata"
	"github.com/Brady-MacDonald/bittorrent-client/src/tracker"
)

func main() {
	if len(os.Args) < 3 {
		panic("Must provide args")
	}

	torrentFile := os.Args[1]
	outputFile := os.Args[2]

	if torrentFile == "" || outputFile == "" {
		panic("Must provide input/output file")
	}

	torrent, err := metadata.GetTorrent(torrentFile)
	if err != nil {
		panic(err)
	}

	var peerID [20]byte
	_, err = rand.Read(peerID[:])
	if err != nil {
		panic(err)
	}

	trackerURL, err := tracker.BuildTrackerURL(torrent, peerID)
	if err != nil {
		panic(err)
	}

	_, err = tracker.GetPeers(trackerURL)
	if err != nil {
		panic(err)
	}
}
