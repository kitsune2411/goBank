package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	Name string `json:"Name"`
}

type Account struct {
	ID        int       `json:"ID"`
	Name      string    `json:"Name"`
	Number    int64     `json:"Number"`
	Balance   int64     `json:"Balance"`
	CreatedAt time.Time `json:"Created_at"`
}

func NewAccount(name string) *Account {
	return &Account{
		Name:      name,
		Number:    int64(rand.Intn(10000000)),
		CreatedAt: time.Now().UTC(),
	}
}
