package metadata

import (
	"bytes"
	"crypto/sha1"
	"io"
	"os"

	"github.com/jackpal/bencode-go"
)

// Parsed .torrent file contents
type Torrent struct {
	// Required
	Announce string   `bencode:"announce"` // URL of tracker to request peers from
	Info     InfoDict `bencode:"info"`

	//Optional
	Comment string `bencode:"comment"`
}

type InfoDict struct {
	Name        string `bencode:"name"`
	Length      int    `bencode:"length"`
	PieceLength int    `bencode:"piece length"` // Size of each piece
	Pieces      []byte `bencode:"pieces"`
}

// Parse the provided bencoded torrent content
func ParseTorrent(bencodedTorrent io.Reader) (*Torrent, error) {
	torrent := &Torrent{}
	err := bencode.Unmarshal(bencodedTorrent, torrent)
	if err != nil {
		return torrent, err
	}

	return torrent, nil
}

// Read torrent file and return io.Reader representation
func OpenTorrent(torrentFile string) (io.Reader, error) {
	torrentBytes, err := os.ReadFile(torrentFile)
	if err != nil {
		return nil, err
	}

	torrentReader := bytes.NewReader(torrentBytes)
	return torrentReader, nil
}

func GetTorrent(torrentFile string) (*Torrent, error) {
	torrentReader, err := OpenTorrent(torrentFile)
	if err != nil {
		return nil, err
	}

	torrent, err := ParseTorrent(torrentReader)
	if err != nil {
		return nil, err
	}

	return torrent, nil
}

// The info hash is the SHA1 of the raw bencoded info dict.
func InfoHash(infoBytes []byte) [20]byte {
	return sha1.Sum(infoBytes)
}
