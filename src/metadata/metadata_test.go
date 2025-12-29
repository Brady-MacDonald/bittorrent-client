package metadata

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	bencodedTorrent, err := OpenTorrent("testdata/archlinux-2019.12.01-x86_64.iso.torrent")
	if err != nil {
		panic(err)
	}

	actualTorrent, err := ParseTorrent(bencodedTorrent)
	if err != nil {
		panic(err)
	}

	expectedTorrent := Torrent{}
	goldenPath := "testdata/archlinux-2019.12.01-x86_64.iso.torrent.golden.json"

	golden, err := os.ReadFile(goldenPath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(golden, &expectedTorrent)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, expectedTorrent, actualTorrent)
}
