# Reading Large file over HTTP in Go

## Non-Streamming example
at first we have an example of `readLoop()` function or handler which read from connection and print out the content of the connection.
```go
func (fs *FileServer) readLoop(conn net.Conn) {
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		file := buf[:n]
		fmt.Println(file)
		fmt.Printf("received %d bytes over the network\n", n)
	}
}
```
in this case we send and bytes over the network and the server with listen and print out the content along with it size
- for example size is : 1000 bytes
```sh 
written 1000 bytes over the network
1 2 3 ....
received 1000 bytes over the network
```
- for example if the size is over 2048 (which is more and we initialize the buf) : 4000 bytes
```sh
written 4000 bytes over the network
1 2 435 21 ...

received 2048 bytes over the network
83 73 12 .....
received 1952 bytes over the network
```
4000 - 2048 = 1952 

## Applied streamming but with problem !  
```go
func (fs *FileServer) readLoop(conn net.Conn) {
	buf := new(bytes.Buffer)
	for {
		n, err := io.Copy(buf, conn)
		if err != nil {
			log.Fatal(err)
		}
		panic("Should Panic Here !!")
		fmt.Println(buf.Bytes())
		fmt.Printf("received %d bytes over the network\n", n)
	}
}
```
Eventhough we implemented the stream mechanism here via `io.Copy()` but the problem here is the `io.Copy()` function with stop until it reaches EOF and EOF is not sent over http therefore it just continue to read copy the stream until the connection is closed which is not\
Another solution of to copy stream ignore the EOF is to specify the size using `io.CopyN()` with an `n` size for reading n sizes stream if after reading n sizes stream the copy will stop 

