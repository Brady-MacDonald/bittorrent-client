# BitTorrent Client

Parse `.torrent` File
- Read .torrent file
- Decode bencode

### Bencode

#### Required

`announce`: Contains the URL of the tracker
`info`: Describes the actual contents

Tracker:
- Server which helps faciliate the peer-to-peer network
- keeps track of which peers have which files

#### Optional

```json
{
    "announce": "https://tracker_url.com",
    "info": {dict},
    "announce-list": [], //(optional list of additional trackers)
    "comment": "{optional}",
    "creation date": <int> (optional),
    "created by": <string> (optional),
    "encoding": <string> (optional)
}
```
