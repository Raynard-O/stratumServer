package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DB is database session container
type DB struct {
	c *sqlx.DB
}

// Permanent session
var cdb *DB

var dbLock sync.Mutex


// Connect will initialise a permanent SQL session
func Connect() (*DB, error) {
	if cdb == nil {
		fmt.Printf("Creating new DB connection")
		newDsn := fmt.Sprintf("postgres://luxor:luxor@localhost:5432/luxor?sslmode=disable")
		//, os.Getenv("HOST"))
		conn, err := sqlx.Connect("postgres", newDsn)
		if err != nil {
			log.Printf("Could not connect to db: %v, %v", err, newDsn)
			return nil, err
		}
		cdb = &DB{c: conn}
	}
	err := cdb.Reconnect()
	if err != nil {
		fmt.Errorf("Could not reconnect to db")
		return nil, err
	}
	return cdb, nil
}

// Reconnect will try to reconnect to database if connection is lost
func (db *DB) Reconnect() error {
	err := db.c.Ping()
	if err != nil {
		fmt.Printf("Reconnecting to db")
		newDsn := fmt.Sprintf("host=localhost port=5432 user=luxor dbname=luxor password=luxor sslmode=disable")
		c, err := sqlx.Connect("mysql", newDsn)
		if err != nil {
			db.c = nil
			fmt.Errorf("Could not reconnect to db")
			return err
		}
		db.c = c
	}
	return nil
}



// GetAllTargets will return list of targets from database
func (db *DB) GetAllTargets(targetl string) (*sql.Rows, error) {
	err := db.Reconnect()
	if err != nil {
		return nil, err
	}
	// var targets []Target
	if targetl != "*" {
		row, err := db.c.Query("select id, message, created_on from targets where id=$1", targetl)
		if err != nil {
			fmt.Errorf("Can't list targets %v", err)
		}

		return row, err
	}
	row, err := db.c.Query("select id, message, created_on from targets")
	if err != nil {
		fmt.Errorf("Can't list targets %v", err)
	}
	return row, err

}

// SaveTarget saves an event into the database
func (db *DB) SaveTarget(v Request) (int64, error) {
	err := db.Reconnect()
	if err != nil {
		return 0, err
	}

	//sqlStatement := `INSERT INTO person (name, nickname) VALUES ($1, $2)`
	//_, err = db.Exec(sqlStatement, p.Name, p.Nickname)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	panic(err)
	//}
	//for _, v := range c.Data {
	//	_, err := db.c.Query("INSERT INTO targets (id, message, created_on) VALUES ($1, $2, $3) ", v.Id, v.Message, v.CreatedOn)
	//
	//	if err != nil {
	//		fmt.Errorf("Can't insert events %v", err)
	//		return 0, err
	//	}
	//}

	//for _, v := range c.Data {
		_, err = db.c.Query("INSERT INTO targets (id, request_at, created_on) VALUES ($1, $2, $3) ", v.ID, v.RequestedAt, v.CreatedOn)

		if err != nil {
			fmt.Errorf("Can't insert events %v", err)
			return 0, err
		}
	//}


	return 1, nil
}
