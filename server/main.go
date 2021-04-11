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
	//fmt.Print(".")
	//for {
	//	netData, err := bufio.NewReader(c).ReadString('\n')
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	temp := strings.TrimSpace(string(netData))
	//	if temp == "STOP" {
	//		break
	//	}
	//	fmt.Println(temp)
	//	counter := strconv.Itoa(count) + "\n"
	//	c.Write([]byte(string(counter)))
	//}
	//c.Close()

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


	listener := new(Listener)

	rpc.Register(listener)

	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}

}






