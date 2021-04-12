package connection

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"strings"
	"time"
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
	userData userAuthentifiedData
}
type userAuthentifiedData struct {
	iam int64
	name	string
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
}

// SysInfo saves the basic system information
type SysInfo struct {
	Hostname string `bson:hostname`
	CPU      string `bson:cpu`

}

func systemSpec() *SysInfo{
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()

	info := new(SysInfo)

	info.Hostname = hostStat.Hostname
	info.CPU = cpuStat[0].ModelName


	return info
}


func (c *Connect) CallAuthorise(serverRequest []string){

	var serverReply Reply
	//get user credentials from Stdin
	u := systemSpec()
	m := make(map[string]interface{})
	m["hostname"]= u.Hostname
	m["CPU"] = u.CPU
	//user := serverRequest[1] + " " + serverRequest[2]
	c.userData.name = u.Hostname
	// sending credentials to server for authentication
	fmt.Println(">>>> sending credentials>>")

	err := c.client.Call("Mining.Authorise", m, &serverReply)

	if err != nil {
		log.Fatal(err)
	}
	//get reply from server
	fmt.Println(">>>> receiving  authorizations>>")
	log.Printf("Reply: %v", serverReply)
	// if credentials are true then allocate an iam credential for client
	if serverReply.Data == true{
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		c.userData.iam = r1.Int63()
		m := make(map[string]int64)
		m[c.userData.name] = c.userData.iam
		err := c.client.Call("Mining.Iam", m, &serverReply)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Reply: %v", serverReply)
	}

}

func (c *Connect) Subscribe(Extranonce1 string)  {

	var serverReply Reply
	//m := make(map[string]string)
	// sending request to server for subscription
	fmt.Println(">>>> sending request>>")
	err := c.client.Call("Mining.Subscribe", Extranonce1, &serverReply)
	//send a getwork notification to server  along with an Extranonce1 id if exists
	//server returns a work and keeps track of subscriptions
	//(server first checks if Extranonce1 being received is done
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>>> getting work response>>>")
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
			c.Subscribe(serverRequest[1])
		default:
			fmt.Println("enter valid request (subscribe, authorise...")
		}

	}

}


