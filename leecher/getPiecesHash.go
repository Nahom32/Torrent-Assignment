package leecher

import (
	"fmt"
	"github.com/zeebo/bencode"
	"os"
	"reflect"
	"encoding/hex"
)

func getPiecesHash(torrentPath string)[]uint8{
	torrentFile, err := os.Open(torrentPath)

	if err != nil{
		fmt.Println("Error openning file: ", err)
		return nil
	}

	defer torrentFile.Close()

	var torrentMap map[string] interface{}

	err = bencode.NewDecoder(torrentFile).Decode(&torrentMap)

	torrentInfo := torrentMap["info"]

	// parse the info interface to map type

	torrentInfoMap, ok := torrentInfo.(map[string]interface{})
	if !ok {
		fmt.Println("Error reading the torrent file!")
		return nil
	}

	pieces := torrentInfoMap["pieces"]

	pieces, err = hex.DecodeString(pieces.(string))

	if err != nil {
		fmt.Println("Error occured during parsing the pieces:", err)
		return nil
	}

	hashPieces, ok := pieces.([]byte)
	// print("firstpiece", hashPieces[0])

	if !ok{
		fmt.Println("piece reading failed")
		return nil
	}
	fmt.Println(reflect.TypeOf(hashPieces))
	return hashPieces

}