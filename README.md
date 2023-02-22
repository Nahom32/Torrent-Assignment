# Torrent-Assignment
BitTorrent Client

BitTorrent is a protocol for downloading and distributing files across the Internet. In contrast with the traditional client/server relationship, in which downloaders connect to a central server (for example: watching a movie on Netflix, or loading the web page you’re reading now), participants in the BitTorrent network, called peers, download pieces of files from each other—this is what makes it a peer-to-peer protocol.
# Contributors
1) Fikremariam Fikadu UGR/9125/12 sec 1<br>
2) Paulos Dessie UGR/6912/12 sec 1<br>
3) Nahom Senay UGR/9334/12 sec 2<br>
4) Kaleab Tekalign UGR/3664/12 sec 2<br>
5) Surafel Getahun UGR/5965/12 sec 2<br>
# Directions
To generate .torrent file use the following command:<br>
go run main.go generateTorrent *filename.type* <FilePath><br>
To seed use the following command:<br>
go run main.go seed *filename.type* <FilePath><br>
To download or leech a file use the following command: <br>
go run main.go download *torrentfilename.type*<torrentFilePath> <br>
N.B. This project runs in localhost with no restart capabilites.
# Module Organization
The final software is composed of the following packages: ‘generateTorrent’, ‘fileCommit’, ‘leecher’, ‘seeder’ and there  is the main file that is used as an interface of communication between the software and the user. It also includes files that are used for the demonstration of the system. Like generated torrent files and mock files used to show the inner workings of the system.
## generateTorrent
The generateTorrent file is used to generate a ‘.torrent’ file used for parsing the metadata of the file to be exchanged. This file is encoded in a format called bencode. The bencode format supports different data types for encoding. This includes lists, int, string, and dictionaries or hash maps. To support the encoding and decoding scheme an imported module called bencode is used. This package can support both the encoding and decoding scheme that is involved in this project. The encoding scheme is used to write to generate the torrent file. This file is in bencoded format. So when it is required for further manipulation the decoding scheme will be used and it will be represented as a dictionary with a string value as a key. The value of the dictionary can be of dynamic type so the ‘interface{}’ keyword is used.
## GenerateTorrent Function
The generateTorrent package includes the GenerateTorrent function. The GenerateTorrent function is capitalised so that it can be referenced from different packages. It accepts the following parameters:
‘filePath’: it is of a string value it is used to parse the location of the file for the initial seeder.<br>
‘trackerURL’: it is used to input the url of the tracker server which connects to another peer.<br>
‘pieceLength’: it is used to input the length of each piece that is supposed to be exchanged.<br>
‘Pieces’: it is an array that includes the hashed values of each chunk that are going to be exchanged throughout the transaction process.<br>
The generate torrent file works as follows:<br>
A struct file is made containing the data format of the torrent File. This is abstracted using the ‘torrentFile’ struct. Then an instance of a torrentFile struct is created containing the parsed meta information of the parsed file through the file path. This will be bencoded and persisted to a file as a ‘.torrent’ file. If in the due process an error is found it returns the error if not it returns a null value.

## The Leecher
The leecher file contains the entity that downloads based on the torrent file available so it is expected to act as a client to the seeder entity. The leecher is expected to dial the server of the seeder. The leecher package contains the following files: the ‘getPiecesHash.go’ and ‘leech.go’.
### getPiecesHash.go
#### getPiecesHash function
The getPiecesHash file contains the getPieceHash function. The getPieceHash function is used to read a torrent path and convert it to a hash map so that the pieces can be read and be checked for integrity of a file. So after the integrity check the needed file can be downloaded or leeched.

### Leech.go
The leech.go file contains the client that dials to the server so that the seeder can achieve the downloading or leeching scheme. The leeching scheme is achieved by sending or requesting the indices of the chunks of the file to be seeded. So that the chunks can be buffered to a list and can be changed to a file.
Leech.go has the following functions: leech and ReadTorrentFile. The leech function is used to download from a seeded file by requesting an index and the ReadTorrentFile is used to read a torrent file and change it to a hashmap to be parsed.

## The Seeder
The seeder is used as a server and uploads or seeds file to leechers. The seeder works by handling a request that came from the leechers. It responds to the index requests of the leecher entity.  The seeder contains the ‘sendchunk’ and seed files. The seed file acts as a server to deliver the chunks to the leecher and the sendchunk file  emulsifies the data to be sent to the leecher.





