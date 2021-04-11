package main

import (
	"fmt"
	"strings"
)

type Listener int

type Reply struct {
	Data bool

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
//authorise raynard ioj