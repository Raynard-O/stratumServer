package main

import (
	"math/rand"
	"net"
	"net/rpc/jsonrpc"
	"time"
)

type clients  struct {
	id int64
	conn *net.Conn
}
// handleConnection implements the jsonrpc.ServeConn method
func handleConnection(conn net.Conn) {
	jsonrpc.ServeConn(conn)
}


// CreateClient creates a map of all clients connected to the server with unique IDs
func (m *Mining) CreateClient(c *net.Conn, reply *interface{}) error {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	//lock resource
	m.m.Lock()
	defer m.m.Unlock()
	// get new ID
	id := r1.Int63()
	// add client to clients
	cli := new(clients)
	cli.conn = c
	cli.id = id
	//add client to clients
	m.clients[id] = cli
	return nil
}


func (m *Mining) GetAllClients(e string,  reply *map[int64]*clients) error {
	reply = &m.clients
	return nil
}

func Write(conn net.Conn, msg string) error {
	_, err := conn.Write([]byte(msg))
	return err
}