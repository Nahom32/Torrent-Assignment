package main

import (
	"fmt"
	"os"
	"torrentassignment/generateTorrent"
	"torrentassignment/leecher"
	"torrentassignment/seeder"

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
}
