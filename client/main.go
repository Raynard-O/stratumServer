package main

import (

	"luxormining/client/connection"

)

func main() {

	connectionInterface :=  connection.InterfaceConnection(connection.New())
	connectionInterface.Conn()
	defer connectionInterface.Close()
	connectionInterface.Listen()
	//connectionInterface.Call()
}


