package tracker

import "github.com/Brady-MacDonald/bittorrent-client/src/metadata"

// Talk to the tracker
// Central location to connect with peers

type Peers struct{}

func GetPeers(metaData metadata.Torrent) Peers {
	return Peers{}
}
