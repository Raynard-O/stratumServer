package db

import "github.com/jinzhu/gorm"

type Request struct {
	gorm.Model
	//ID        string `json:"id" db:"id"`
	Name 	string `json:"name" db:"name"`
	CPU string `json:"cpu" db:"cpu"`
	RequestedAt string `json:"requested_at" db:"requested_at"`

}
