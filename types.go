package main

import "math/rand"

type Account struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Number  int64  `json:"Number"`
	Balance int64  `json:"Balance"`
}

func NewAccount(name string) *Account {
	return &Account{
		ID:     rand.Intn(1000),
		Name:   name,
		Number: int64(rand.Intn(10000000)),
	}
}
