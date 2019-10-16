package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	_ "github.com/lib/pq"
)

func connet() *pg.DB {

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "chitchat"
	)

	db := pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: dbname,
	})

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		panic(err)
	}
	return db
}

type User struct {
	Id        int64
	Name      string
	Emails    []string
	tableName struct{} `sql:"users"`
}

type Book struct {
	Id        int64
	Title     string
	AuthorId  int64
	Author    *User
	tableName struct{} `sql:"books"`
}

func (s Book) String() string {
	return fmt.Sprintf("Book<%d %s %s>", s.Id, s.Title, s.Author)
}

func (u User) String() string {
	return fmt.Sprintf("User(%d %s %v)", u.Id, u.Name, u.Emails)
}

func CreateTabel(db *pg.DB) error {
	for _, model := range []interface{}{&User{}, &Book{}} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateUser(db *pg.DB, user *User) (orm.Result, error) {
	res, err := db.QueryOne(user, `
		INSERT INTO users (name, emails) VALUES (?name, ?emails)
		RETURNING id
	`, user)
	return res, err
}

func CreateBook(db *pg.DB, book *Book) error {
	_, err := db.QueryOne(book, `
		INSERT INTO books (title, author_id) VALUES (?title, ?author_id)
		RETURNING id
	`, book)
	return err
}

func main() {
	db := connet()
	err := CreateTabel(db)
	if err != nil {
		panic(err)
	}

	user1 := &User{
		Name:   "alice",
		Emails: []string{"alice@qq.com", "alice@baidu.com"},
	}
	res, err := CreateUser(db, user1)
	if err != nil {
		panic(err)
	}
	fmt.Println(res) // &{0xc000114f30 1 1}

	res, err = CreateUser(db, &User{
		Name:   "bob",
		Emails: []string{"bob@qq.com", "bob@baidu.com"},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(res) // &{0xc000114fc0 1 1}

	book1 := &Book{
		Title:    "Cool book",
		AuthorId: user1.Id,
	}
	err = CreateBook(db, book1)
}
