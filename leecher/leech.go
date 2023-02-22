package leecher

import (
  "fmt"
  "log"
  "net"
  "os"
  "reflect"
  "github.com/zeebo/bencode"
  "encoding/gob"
  "crypto/sha1"
//   "bufio"
  //"torrentassignment/seeder"
  //"crypto/sha1"
)
func compare(chunk1 []uint8 , chunk2 []uint8) bool{
  return reflect.DeepEqual(chunk1,chunk2)
}
func Leech(torrentFile string){
	fmt.Println("LEECHING")
  	
    
	dataHold := getPiecesHash(torrentFile)
	// fmt.Println(length)

  	bitField := make([]int,len(dataHold))

  	remaining := len(dataHold)
	dataField := make([][]byte, len(dataHold)/20)
	// fmt.Println(bitField)
  	// i := 0
  	
  	for i := 0; i < remaining; i++ {
		fmt.Println(i * 20, len(dataHold))
		if i*20 == len(dataHold){
			break
		}
		if bitField[i] == 0{
			fmt.Println(i)
			conn, err := net.Dial("tcp", ":8080")
			encoder := gob.NewEncoder(conn)
			decoder := gob.NewDecoder(conn)
			if err != nil {
				fmt.Println(err)
				return
			}
			if f, ok := conn.(interface{ Flush() }); ok {
				print("flushed")
				f.Flush()
			}
			defer conn.Close()
			// fmt.Println("before encoding")
			err = encoder.Encode(i)
			// conn.Write([]byte(fmt.Sprint(i)))
			if err != nil{
				fmt.Println("errror encoding")
			}
			if f, ok := conn.(interface{ Flush() }); ok {
    			f.Flush()
			}
			// fmt.Println("after encoding")
			
			
			var message []byte
			err = decoder.Decode(&message)
			h := sha1.New()
			h.Write(message)
			hash := h.Sum(nil)
			// print("hash", hash)
			fmt.Println("hash values:", reflect.ValueOf(hash))
			fmt.Println("info piece hash:", reflect.ValueOf(dataHold[i* 20: (i+1) *20]))
			
			if compare(hash, []byte(dataHold[i* 20: (i+1) *20])){
			dataField[i] = message
			fmt.Println("merged")
			i ++
		}else{
			fmt.Println("File Corrupted or missing data")
		}
			
			// fmt.Println("recieced indx",message)
			if err != nil{
				fmt.Println("error")
				return
			}
		}
    }
	// fmt.Println(dataField)
	newMerged := Merge(dataField)
	PersistFile("newrequiements.txt",newMerged)
  }  
// func RequestChunk(id int,dataHold []uint8)uint8{
//   return dataHold[id]
// } 


func ReadTorrentFile(torrentFile string) map[string]interface{}{
    file, err := os.Open(torrentFile)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return nil
  }
  defer file.Close()

  // Decode the torrent file into a map
  var torrentMap map[string]interface{}
  err = bencode.NewDecoder(file).Decode(&torrentMap)
  if err != nil {
    fmt.Println("Error decoding file:", err)
    return nil
  }
  return torrentMap
}
func Merge(data [][]uint8) []uint8{
  merged := []uint8{}
  for _,j := range data{
    for _,k := range j{
      merged = append(merged, k)
    }
  }
  return merged
}



func PersistFile(s string, data []byte){
  file,err := os.Create(s)
  if err != nil{
    log.Print(err)
    return
  }
  n,er:= file.Write(data)
  if er != nil{
    log.Print(er)
    return
  }
  log.Printf("%d bytes were written successfully",n)

}