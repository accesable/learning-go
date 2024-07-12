package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type FileServer struct{}

func (fs *FileServer) start() {
	ln, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go fs.readLoop(conn)
	}
}

func (fs *FileServer) readLoop(conn net.Conn) {
	buf := new(bytes.Buffer)
	for {
		n, err := io.CopyN(buf, conn, 5)
		fmt.Println("Read Loop")
		if err != nil {
			fmt.Printf("%v meet EOF", buf.Bytes())
			fmt.Println(err)
			log.Fatal(err)
		}
		// panic("Should Panic Here !!")
		fmt.Println(buf.Bytes())
		fmt.Printf("received %d bytes over the network\n", n)
	}
}

func main() {
	go func() {
		time.Sleep(4 * time.Second)
		sendFile(15)
	}()
	server := &FileServer{}
	server.start()
}

// testing
func sendFile(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", ":3001")
	if err != nil {
		return err
	}
	// read all to buffer
	// n, err := conn.Write(file)
	// convert to streaming
	// n, err := io.Copy(conn, bytes.NewReader(file))
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size))
	if err != nil {
		return err
	}
	fmt.Printf("written %d bytes over the network\n", n)
	return nil
}
