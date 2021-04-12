package main

import (
	"bufio"
	"fmt"
	"log"
	"luxormining/server/db"
	"net"
	"net/rpc"
	"os"
	_ "time"
)



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

	//initializing server connections
	mining := Init()
	defer mining.DB.Db.Close()
	var rq []db.AuthorizationRequest
	mining.DB.FindAll(&rq)
	for _,v := range rq{
		log.Println(v)
	}
	//log.Println(&rq)
	rpc.Register(mining)

	for {
		conn, err := l.Accept()
		//save every connected miner
		var R *interface{}
		mining.CreateClient(&conn, R)
		go func() {
			for {
			//read from stdin
			reader := bufio.NewReader(os.Stdin)
			line, _, err := reader.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("line", line)
			err = Write(conn, string(line))
				if err != nil {
					log.Fatal(err)
				}
		}
		}()


		if err != nil {
			continue
		}
		go handleConnection(conn)
	}

}






