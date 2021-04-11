package connection

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"strings"
)



type InterfaceConnection interface {
	Conn()
	Close() error
	//Call()
	Listen()
}

type Connect struct {
	client *rpc.Client
	conn net.Conn

}


type Reply struct {
	Data bool

}

type User struct {
	name string
	password string
}


func New()*Connect{
	return &Connect{}
}

func (c *Connect) Conn() {

	client , err := jsonrpc.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
	c.client = client
	//c.conn = conn
}

func (c *Connect) CallAuthorise(serverRequest []string){

	var serverReply Reply
	//get user credentials from Stdin
	user := serverRequest[1] + " " + serverRequest[2]
	// sending credentials to server for authentication
	fmt.Println(">>>> sending credentials>>")

	err := c.client.Call("Listener.Authorise", user, &serverReply)

	if err != nil {
		log.Fatal(err)
	}
	//get reply from server
	fmt.Println(">>>> receiving  credentials>>")
	log.Printf("Reply: %v", serverReply)
}


func (c *Connect) Close() error {
	return c.client.Close()
}

func (c *Connect) Listen(){

	for {
		//read from stdin
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}


		serverRequest := strings.Split(string(line), " ")
		//check method being called
		switch serverRequest[0] {
		case "authorise":
			c.CallAuthorise(serverRequest)
		case "subscribe":

		default:

		}

	}

}
