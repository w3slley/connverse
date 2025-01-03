package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"sync"
)

const (
	HOST     = "localhost"
	PORT     = "8000"
	PROTOCOL = "tcp"
	CONN     = HOST + ":" + PORT
)

var wg sync.WaitGroup

func Read(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Disconnected")
			return
		}
		log.Print(message)
	}
}

func Write(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		_, err = writer.WriteString(message)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		if err = writer.Flush(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
}

func main() {
	wg.Add(1)

	conn, err := net.Dial(PROTOCOL, CONN)
	if err != nil {
		log.Println(err)
	}

	go Read(conn)
	go Write(conn)

	wg.Wait()
}
