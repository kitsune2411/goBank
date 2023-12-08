package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PgStore struct {
	db *sql.DB
}

func NewPgStore() (*PgStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PgStore{
		db: db,
	}, nil
}

func (s *PgStore) Init() error {
	return s.createAccountTable()
}

func (s *PgStore) createAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		name varchar(100),
		number serial,
		balance serial,
		created_at timestamp
	) `

	_, err := s.db.Exec(query)
	return err
}

func (s *PgStore) CreateAccount(acc *Account) error {
	query := `insert into account 
	(name, number, balance, created_at)
	values 
	($1, $2, $3, $4)`
	respon, err := s.db.Query(query, acc.Name, acc.Number, acc.Balance, acc.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", respon)
	return nil
}
func (s *PgStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PgStore) DeleteAccount(id int) error {
	return nil
}
func (s *PgStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
func (s *PgStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("select * from account")

	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {

		account := &Account{}
		err := rows.Scan(
			&account.ID,
			&account.Name,
			&account.Number,
			&account.Balance,
			&account.CreatedAt)
		if err != nil {
			return nil, err
		}
		fmt.Println(account)
		accounts = append(accounts, account)
	}
	return accounts, nil
}
