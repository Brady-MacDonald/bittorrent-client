package main

import (
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

	torrentReader, err := metadata.OpenTorrent(torrentFile)
	if err != nil {
		panic(err)
	}

	meta, err := metadata.ParseTorrent(torrentReader)
	if err != nil {
		panic(err)
	}

	tracker.GetPeers(meta)
	// downloader := download.New(meta, peers)
	// downloader.Start()
}
