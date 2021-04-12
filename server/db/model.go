package db

import "github.com/jinzhu/gorm"

type AuthorizationRequest struct {
	gorm.Model
	//ID        string `json:"id" db:"id"`
	Name 	string `json:"name" db:"name"`
	CPU string `json:"cpu" db:"cpu"`
	RequestedAt string `json:"requested_at" db:"requested_at"`
}

type SubscriptionRequest struct {
	gorm.Model
	Extranonce1     string `json:"extranonce" db:"extranonce"`
	Extranonce2Size int `json:"extranonce2size" db:"extranonce2size"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
	CreatedAt string `json:"created_at" db:"created_at"`
	CompletedAt string `json:"completed_at" db:"completed_"`
}