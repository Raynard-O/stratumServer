package main

import (
	"fmt"
	"strings"
)

type Listener struct {
	clients map[string]int64
}

type Reply struct {
	Data bool

}

func Init() *Listener{
	return &Listener{clients: make(map[string]int64)}
}


func (l *Listener) Authorise(user string, reply *interface{}) error {
	//get user credentials
	s := strings.Split(user, " ")
	fmt.Printf("Authorise: name:%v password: %v :\n", s[0], s[1])

	//verify user credentials
	if s[1] == s[1] {
		*reply = Reply{true}
	}else {
		*reply = Reply{false}
	}
	//return verification bool

	return nil
}


func (l *Listener) Iam(iam map[string]int64, reply *interface{}) error {
	//get iam data
	fmt.Printf("IAM: iam  add for user:%v\n ", iam)
	//add to the array of user credentials
	l.clients = iam
	*reply = Reply{true}
	return nil
}

//authorise raynard ioj