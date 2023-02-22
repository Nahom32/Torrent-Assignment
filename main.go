package main

import (
	// "crypto/sha1"
	"fmt"
	// "io/ioutil"
	"os"
	//"net"
	"torrentassignment/generateTorrent"
	"torrentassignment/leecher"
	"torrentassignment/seeder"
	//"log"

)

func main() {
	args := os.Args

	fmt.Println(len(args))

	command := args[1]

	switch command{
	case "generateTorrent":
		if (len(args) == 3){
			fmt.Println("Generating torrent ")
			generateTorrent.GenerateTorrent(os.Args[2],"http://104.28.16.69/announce",1024)
			}else{
			fmt.Println("No")
		}
		
	case "seed":
		if (len(args) == 3){
			fmt.Println("Seeding torrent:", args[2])
			//log.Print(leecher.ReadTorrentFile(args[2]))
			seeder.Seed(args[2])
			}else{
			fmt.Println("No")
			

		}
	case "download":
		if (len(args) == 3){
			fmt.Println("Download the given torrent")
			leecher.Leech(args[2])
			}else{
			fmt.Println("No")
		}
		
	default:
		fmt.Println("Unknown command:", command)
	}

	// // Open the file for reading
	// file, err := os.Open("video.mkv")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// // Determine the size of the file
	// fileInfo, err := file.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fileSize := fileInfo.Size()

	// // Define the chunk size
	// chunkSize := int64(1024 * 1024) // 1 KB

	// // Compute the number of chunks
	// numChunks := fileSize / chunkSize
	// if fileSize%chunkSize != 0 {
	// 	numChunks++
	// }

	// // Read each chunk and compute the SHA1 hash
	// hash := sha1.New()
	// for i := int64(0); i < numChunks; i++ {
	// 	// Read the next chunk
	// 	offset := i * chunkSize
	// 	if i == numChunks-1 {
	// 		chunkSize = fileSize - offset
	// 	}
	// 	chunk := make([]byte, chunkSize)
	// 	_, err := file.ReadAt(chunk, offset)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	// Compute the SHA1 hash of the chunk
	// 	hash.Write(chunk)
	// 	chunkHash := hash.Sum(nil)
	// 	hash.Reset()

	// 	// Print the hash of the chunk
	// 	fmt.Printf("Chunk %d hash: %x\n", i+1, chunkHash)
	// }

	// // Print the final info hash
	// finalHash := hash.Sum(nil)
	// fmt.Printf("Final info hash: %x\n", finalHash)
}
