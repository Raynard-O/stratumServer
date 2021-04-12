package main

import (
	"fmt"
	"log"
	"luxormining/server/db"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	_ "time"
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





	//req := db.Request{
	//	Name: "user[].(string)",
	//	CPU:        "user[11122].(string)",
	//	RequestedAt: time.Now().UTC().Local().String(),
	//}

	listener := Init()
	defer listener.DB.Db.Close()
	var rq []db.Request
	listener.DB.FindAll(&rq)

	log.Println(&rq)
	rpc.Register(listener)

	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}

}






