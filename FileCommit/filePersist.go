package filecommit
import (
	"os"
	"log"

)
// func merge(data [][]byte) [][]byte {

// }
func PersistFile(s string, data [][]byte){
	file,err := os.Create(s)
	if err != nil{
		log.Print(err)
		return
	}
	n,er:= file.Write(data[0])
	if er != nil{
		log.Print(er)
		return
	}
	log.Printf("%d bytes were written successfully",n)

}