package main

import (
	"fmt"
	"log"
	"luxormining/server/db"
)



type Listener struct {
	clients map[string]int64
	DB *db.DB
}

type Reply struct {
	Data bool
}



func Init() *Listener{
	//database initialise
	Db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(systemSpec())
	fmt.Println("system initializing ..... Done")
	return &Listener{
		clients: make(map[string]int64),
		DB: Db,
		}
}


func (l *Listener) Authorise(user map[string]interface{}, reply *interface{}) error {
	//get user credentials
	//s := strings.Split(user, " ")
	log.Printf("%v:\n", user["hostname"])
	log.Println("1 device(s) found:")
	log.Printf("0 - %v :\n", user["CPU"])


	//verify user credentials
	//if s[1] == s[1] {
	//save to db
	//	fmt.Println(r)
	//return verification bool
	*reply = Reply{true}
	log.Printf("Authorising...")
	//}else {

	//}


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