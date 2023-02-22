package generateTorrent

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
	"github.com/zeebo/bencode"
	"bytes"
)

type torrentFile struct {
	Name        string
	Length      int64
	PieceLength int64
	Pieces      []byte
}

func GenerateTorrent(filePath string, trackerURL string, pieceLength int64) error {
	// Open the file to be shared
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Get the file info for the file to be shared
	info, err := f.Stat()
	if err != nil {
		return err
	}

	// Calculate the number of pieces
	numPieces := info.Size() / pieceLength
	if info.Size()%pieceLength != 0 {
		numPieces++
	}

	// Create the torrent file data structure
	tf := torrentFile{
		Name:        info.Name(),
		Length:      info.Size(),
		PieceLength: pieceLength,
		Pieces:      make([]byte, numPieces),
	}

	// Read the file piece by piece and hash each piece
	pieceBuf := make([]byte, pieceLength)
	hashes := make([][]byte, numPieces)
	h := sha1.New()

	for i := int64(0); i < numPieces; i++ {
		n, err := io.ReadFull(f, pieceBuf)
		if err != nil && err != io.ErrUnexpectedEOF {
			return err
		}
		if n == 0 {
			break
		}
		h.Reset()
		_, err = h.Write(pieceBuf[:n])
		if err != nil {
			return err
		}
		hash := h.Sum(nil)
		fmt.Println(hash)
		hashes[i] = make([]byte, len(hash))
		copy(hashes[i], hash)
	}

	tf.Pieces = bytes.Join(hashes, nil)
	fmt.Println(len(tf.Pieces))
	// Create the torrent dictionary
	torrent := map[string]interface{}{
		"announce":    trackerURL,
		"creation date": time.Now().Unix(),
		"info": map[string]interface{}{
			"length": tf.Length,
			"name":   tf.Name,
			"piece length": tf.PieceLength,
			"pieces": fmt.Sprintf("%x", tf.Pieces),
		},
		
	}

	// fmt.Println(torrent["info"])
	// Create the .torrent file
	torrentFilename := fmt.Sprintf("%s.torrent", filepath.Base(filePath))
	torrentFile, err := os.Create(torrentFilename)
	if err != nil {
		return err
	}
	defer torrentFile.Close()

	// Encode the torrent dictionary and write it to the .torrent file
	
	//err = encodeTorrent(torrent, torrentFile)
	err = bencode.NewEncoder(torrentFile).Encode(&torrent)
	if err != nil {
		return err
	}

	return nil
}

// func encodeTorrent(dict map[string]interface{}, w io.Writer) error {
// 	for k, v := range dict {
// 		fmt.Fprintf(w, "%d:%s", len(k), k)
// 		switch v := v.(type) {
// 		case string:
// 			fmt.Fprintf(w, "%d:%s", len(v), v)
// 		case int64:
// 			fmt.Fprintf(w, "i%de", v)
// 		case []byte:
// 			fmt.Fprintf(w, "%d:", len(v))
// 			w.Write(v)
// 		case []interface{}:
// 			fmt.Fprint(w, "l")
// 			for _, e := range v {
// 				encodeTorrent(map[string]interface{}{"": e}, w)
// 			}
// 			fmt.Fprint(w, "e")
// 		case map[string]interface{}:
// 			fmt.Fprint(w, "d")
// 			encodeTorrent(v, w)
// 			fmt.Fprint(w, "e")
// 		default:
// 			return nil
// 		}
// 	}
// 	return nil
// }
