package seeder

import (
	"fmt"
	"net"
	"encoding/gob"
)

func Seed(filepath string){
	
	conn, err := net.Listen("tcp", ":8080")

	if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()

	for{
		fmt.Println("listening")
		req, err := conn.Accept()
		if err != nil {
        	fmt.Println(err)
        	return
    	}
		defer req.Close()
		if req != nil{
			HandleRequest(req, filepath)
		}
		fmt.Println("listened")
	}
}

func HandleRequest(req net.Conn, filepath string){
	var buf int
	decoder := gob.NewDecoder(req)
	err := decoder.Decode(&buf)
	// n, err := req.Read(buf)
	if err != nil{
		fmt.Println("some error occured")
		// req.Write([]byte("9"))
	}
	fmt.Println(buf)
	sendChunk(buf, filepath, req)
}