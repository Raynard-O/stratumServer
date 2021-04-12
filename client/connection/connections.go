package connection

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"log"
	"luxormining/server/db"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"strings"
)


//InterfaceConnection is the interface for client connections
type InterfaceConnection interface {
	Conn()
	Close() error
	Listen()
}

//Connect holds the connection parameter
type Connect struct {
	client *rpc.Client
	conn net.Conn
	userData userAuthentifiedData
}

// userData is the client info
type userAuthentifiedData struct {
	iam int64
	name	string
}

//Reply is the  structure of reply data from the server methods
type Reply struct {
	Data bool
}

//New creates a new interface connection
func New()*Connect{
	return &Connect{}
}
// Conn handles the connection to the server
func (c *Connect) Conn() {
	// Dial the server at localhost 8080
	client , err := jsonrpc.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
	// saves client connection
	c.client = client
}

// SysInfo saves the basic system information
type SysInfo struct {
	Hostname string `bson:hostname`
	CPU      string `bson:cpu`
}

//systemSpec gets the machine's' details
func systemSpec() *SysInfo{
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()

	info := new(SysInfo)

	info.Hostname = hostStat.Hostname
	info.CPU = cpuStat[0].ModelName
	return info
}

//CallAuthorise is used to authorise access of client to sever by systemSpec credentials
func (c *Connect) CallAuthorise(){

	var serverReply Reply
	//get user credentials from Stdin
	u := systemSpec()
	//map system credentials for server identification
	m := make(map[string]interface{})
	m["hostname"]= u.Hostname
	m["CPU"] = u.CPU

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
	//// if credentials are true then allocate an iam credential for client
	//if serverReply.Data == true{
		//s1 := rand.NewSource(time.Now().UnixNano())
		//r1 := rand.New(s1)
		//c.userData.iam = r1.Int63()
		//m := make(map[string]int64)
		//m[c.userData.name] = c.userData.iam
		//err := c.client.Call("Mining.Iam", m, &serverReply)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//log.Printf("Reply: %v", serverReply)
}


// Subscribe is used to request a getwork from server or continue an existing work.
func (c *Connect) Subscribe(Extranonce1 string)  {
	// server repiles subscriptions with a map description of the job
	var serverReply map[string]*db.SubscriptionRequest

	// sending request to server for subscription
	fmt.Println(">>>> sending request>>")
	//send a getwork notification to server  along with an Extranonce1 id if exists
	err := c.client.Call("Mining.Subscribe", Extranonce1, &serverReply)

	//server returns a work and keeps track of subscriptions
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>>> getting work response>>>")
	log.Printf("Reply: %v", serverReply["work"])
}

// Close connection to server
func (c *Connect) Close() error {
	return c.client.Close()
}

func (c *Connect) Listen(){

	for {

		//read from stdin
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Panic(err)
		}

		serverRequest := strings.Split(string(line), " ")
		//check method being called
		switch serverRequest[0] {
		case "authorise":
			c.CallAuthorise()
		case "subscribe":
			if len(serverRequest) > 1{
				c.Subscribe(serverRequest[1])
			}else {
				c.Subscribe("")
			}
		default:
			fmt.Println("enter valid request (subscribe, authorise...")
		}
	}

}


