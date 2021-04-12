package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

func handleConnection(conn net.Conn) {
	jsonrpc.ServeConn(conn)

}




func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	PORT := ":" + arguments[1]

	l, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()


	listener := Init()
	rpc.Register(listener)

	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}

}






