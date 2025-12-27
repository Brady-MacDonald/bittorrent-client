package main

import "crypto/sha1"

// Handle Torrent parsing

type MetaInfo struct {
	Announce string   `bencode:"announce"`
	Info     InfoDict `bencode:"info"`
}

type InfoDict struct {
	Name        string `bencode:"name"`
	PieceLength int    `bencode:"piece length"`
	Pieces      []byte `bencode:"pieces"`
	Length      int    `bencode:"length"`
}

// The info hash is the SHA1 of the raw bencoded info dict.
func InfoHash(infoBytes []byte) [20]byte {
	return sha1.Sum(infoBytes)
}
