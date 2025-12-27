package main

import "os"

func main() {
	torrentPath := os.Args[1]
	output := os.Args[2]

	if torrentPath == "" || output == "" {
		panic("Bad")
	}

	meta := torrent.Parse(torrentPath)
	peers := tracker.GetPeers(meta)

	downloader := download.New(meta, peers)
	downloader.Start()
}
