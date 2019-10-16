package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func connectDB() *sql.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "chitchat"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func query(db *sql.DB) {
	var id, name, emails string

	rows, err := db.Query(" select * from users where name=$1", "alice")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &emails)

		if err != nil {
			fmt.Println(err)
		}
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id, name, emails)
}

func main() {
	db := connectDB()
	query(db)
}
