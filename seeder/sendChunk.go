package seeder

import(
	"net"
	"os"
	"io"
	// "strconv"
	// "fmt"
	"encoding/gob"
	// "crypto/sha1"
)

func sendChunk(buf int, filepath string, req net.Conn){

	f, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer f.Close()

	// Get the file info for the file to be shared
	info, err := f.Stat()
	if err != nil {
		return
	}

	// Calculate the number of pieces
	numPieces := info.Size() / int64(1024)
	if info.Size()%int64(1024) != 0 {
		numPieces++
	}

	// Create the torrent file data structure
	

	// Read the file piece by piece and hash each piece
	pieceBuf := make([]byte, 1024)

	for i := int64(0); i < numPieces; i++ {
		n, err := io.ReadFull(f, pieceBuf)
		if err != nil && err != io.ErrUnexpectedEOF {
			return 
		}
		if n == 0 {
			break
		}
		if int(i) == buf{
		// fmt.Println(pieceBuf[:n])
		encoder := gob.NewEncoder(req)
		err = encoder.Encode(pieceBuf[:n])
		}
	}


	// ===============================================================
	// file, err = os.Open(filepath)

	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()
	// fileInfo, err := file.Stat()

	// fileSize := fileInfo.Size()


	// if err != nil{
	// 	panic(err)
	// }
	// chunkSize := int64(1024)

	// numChunks := fileSize / chunkSize

	// if fileSize%chunkSize != 0 {
	// 	numChunks++
	// }

	// // idx, err := strconv.Atoi(string(buf))
	// // fmt.Println("index got in", idx)
	// offset := int64(buf) * chunkSize
	// // fmt.Println("the offset:", fileSize)
	
	// if chunkSize + offset > fileSize{
	// 	chunkSize = fileSize - offset + int64(1)
	// }
	// // fmt.Println("fileSize", fileSize)
	// chunk := make([]byte, chunkSize)
	// // fmt.Println("offset", offset)
	// _, err = file.ReadAt(chunk, offset)
	// if err != nil {
	// 		fmt.Println("got error at reading")
	// 		return
	// }
	// // fmt.Println(chunk)
	// fmt.Println("===============================================================================")
	// encoder := gob.NewEncoder(req)
	// err = encoder.Encode(chunk)
	// if f, ok := req.(interface{ Flush() }); ok {
    // f.Flush()
	// }
	// req.Write(chunk)
}