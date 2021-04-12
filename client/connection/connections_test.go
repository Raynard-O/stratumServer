package connection

import (
	"github.com/stretchr/testify/assert"
	"net/rpc/jsonrpc"
	"testing"
)


//Test New connection initiation function
func TestNew(t *testing.T) {
	conn := &Connect{}
	n := New()
	assert.Equal(t, conn, n)
}


// Test the dial function and connection creation
func TestConn(t *testing.T) {
	// Dial the server at localhost 8080
		client , err := jsonrpc.Dial("tcp", "localhost:8080")
		assert.NoError(t, err)
		assert.NotEmpty(t, client)
}


// testing the SystemSpec function
func TestSystemSpec(t *testing.T) {
	spec := systemSpec()
	assert.IsType(t, &SysInfo{}, spec)
	assert.NotEmpty(t, spec)
}
//test the server authorise function
func TestConnect_CallAuthorise(t *testing.T) {
	//initializing connection interface
	n := New()
	//dail server
	client , err := jsonrpc.Dial("tcp", "localhost:8080")
	//assert no error
	assert.NoError(t, err)
	n.client = client
	n.CallAuthorise()
	assert.NotEmpty(t, n.userData)
}

func TestConnect_Subscribe(t *testing.T) {
	//initializing connection interface
	n := New()
	//dail server
	client , err := jsonrpc.Dial("tcp", "localhost:8080")
	//assert no error
	assert.NoError(t, err)
	n.client = client
	n.Subscribe("")
	assert.NotEmpty(t, n.userData)
}

// Test Listen
func TestConnect_Listen(t *testing.T) {
	n := New()
	//dail server
	client , err := jsonrpc.Dial("tcp", "localhost:8080")
	//assert no error
	assert.NoError(t, err)
	n.client = client
	assert.NotPanics(t,n.Listen )
}
