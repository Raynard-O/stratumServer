package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

// DB is database session container
type DB struct {
	//c *sqlx.DB
	Db *gorm.DB

}

// Permanent session
var cdb *DB

var dbLock sync.Mutex


// Connect will initialise a permanent SQL session
func Connect() (*DB, error) {

	if cdb == nil {
		fmt.Printf("Creating new DB connection")
		db, err := gorm.Open( "postgres", "host=localhost port=5432 user=luxor dbname=luxor sslmode=disable password=luxor")

		if err != nil {
			log.Printf("Could not connect to db: %v", err)
			return nil, err
		}
		cdb = &DB{Db: db}
	}

	return cdb, nil
}



//func (db *DB) Insert(req Request){
//	fmt.Println(&req)
//	db.Db.Create(&req)
//	var re []Request
//	db.Db.Find(&re)
//	fmt.Println(re)
//}


func (db *DB) Insert(req Request){
	fmt.Println(&req)
	db.Db.Create(&req)

}

func (db *DB) FindAll(req *[]Request){
	db.Db.Find(&req)
}

