package main

import (
	"github.com/stretchr/testify/assert"
	"luxormining/server/db"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	m := Init()
	assert.NotEmpty(t, m)
}

func TestMining_Authorise(t *testing.T) {
	m := Init()
	assert.NotEmpty(t, m)

	req := db.AuthorizationRequest{
		Name: "macbook2User",
		CPU:        "Intel(R) Core(TM) i5-8257U CPU @ 1.40GHz",
		RequestedAt: time.Now().UTC().Local().String(),
	}
	ma := make(map[string]interface{})
	ma["user"] = req
	var reply interface{}
	err := m.Authorise(ma, &reply)

	assert.NoError(t, err)
	assert.NotEmpty(t, reply)
}

