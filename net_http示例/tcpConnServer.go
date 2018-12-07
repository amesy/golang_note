package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handlers(conn net.Conn) {
	fmt.Fprintf(conn, "%s\n", time.Now().Format("2006-01-02 15:04:05"))
	conn.Close()
}

func main() {
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(conn)
		}
		go handlers(conn)
	}
}
