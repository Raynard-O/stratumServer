package main

import (
	"fmt"
	"log"
	"luxormining/server/db"
	"sync"
	"time"
)



type Mining struct {
	m       sync.Mutex
	clients map[int64]*clients
	DB *db.DB
}

type Reply struct {
	Data bool
}



func Init() *Mining{
	//database initialise
	Db, err := db.Connect()


	Db.Db.AutoMigrate(&db.AuthorizationRequest{})
	Db.Db.AutoMigrate(&db.SubscriptionRequest{})

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(systemSpec())
	fmt.Println("system initializing ..... Done")
	return &Mining{
		clients: make(map[int64]*clients),
		DB: Db,
		}
}


func (m *Mining) Authorise(user map[string]interface{}, reply *interface{}) error {
	//get user credentials
	//s := strings.Split(user, " ")
	log.Printf("%v:\n", user["hostname"])
	log.Println("1 device(s) found:")
	log.Printf("0 - %v :\n", user["CPU"])
	//create db structure
	req := db.AuthorizationRequest{
		Name: user["hostname"].(string),
		CPU:        user["CPU"].(string),
		RequestedAt: time.Now().UTC().Local().String(),
	}
	//verify user credentials
	//if s[1] == s[1] {
	//save to db
	m.DB.Insert(req)
	//if err != nil {
	//	log.Println(err)
	//}
	//return verification bool
	*reply = Reply{true}
	log.Printf("Authorising...")
	//}else {
	//}
	return nil
}


func (m *Mining) Iam(iam map[string]int64, reply *interface{}) error {
	//get iam data
	fmt.Printf("IAM: iam  add for user:%v\n ", iam)
	//add to the array of user credentials
	//m.clients = iam
	*reply = Reply{true}
	return nil
}


func (m *Mining) Subscribe(Ext1 string, reply *interface{}) error {
	rMap := make(map[string]*db.SubscriptionRequest)
	sub := db.SubscriptionRequest{
		Extranonce1:     Ext1,
	}
	//get work from database if exist
	subTable := m.DB.Db.Table("subscriptions")
	subTable.Find(&sub)
	if sub.Extranonce1 == "" && sub.CompletedAt == ""{
		//if not allocate another work
		newSub := db.SubscriptionRequest{

			Extranonce1:     randomHex(20),
			Extranonce2Size: 4,
			UpdatedAt:       time.Now().UTC().Local().String(),
			CreatedAt:       time.Now().UTC().Local().String(),
			CompletedAt: "",
		}
		m.DB.Db.Create(&newSub)
		fmt.Printf("work continued: Extranonce1:  %v\n ", newSub.Extranonce1)
		rMap["work"] = &newSub
		*reply = rMap
	}else {
		//set updated_at to now
		m.DB.Db.Model(&sub).Update("updated_at",time.Now().UTC().Local().String() )

		fmt.Printf("work continued: Extranonce1:  %v\n ", sub.Extranonce1)
		rMap["work"] = &sub
		*reply = rMap
		//*reply = Reply{true}
	}

	return nil
}






func (m *Mining) Notify(Ext1 string, reply *interface{}) error {
	//create random jobs
	randomJobs := db.SubscriptionRequest{

		Extranonce1:     randomHex(20),
		Extranonce2Size: 4,
		UpdatedAt:       time.Now().UTC().Local().String(),
		CreatedAt:       time.Now().UTC().Local().String(),
		CompletedAt: "",
	}
	fmt.Println(randomJobs)
	//notify all connected miners of new job


	return nil
}

//authorise raynard ioj
