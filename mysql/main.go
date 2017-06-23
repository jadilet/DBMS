package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func findById(id int) {
	var articleTitle sql.NullString

	row := db.QueryRow("select id, articleTitle from news where id = ?", id)

	err := row.Scan(&articleTitle)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("find method start")
	fmt.Println(articleTitle.Value())
}

func main() {
	var err error

	// creating structure database
	db, err := sql.Open("mysql", "root:2468951@tcp(localhost:3306)/AllInOne_development?charset=utf8")

	if err != nil {
		fmt.Println(err)
		return
	}

	// setting pool
	db.SetMaxOpenConns(10)

	// checking is connected to database
	err = db.Ping()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(db.Stats().OpenConnections)

	rows, err := db.Query("select articleTitle, id from news")

	for rows.Next() {
		var articleTitle sql.NullString
		var id int
		err = rows.Scan(&articleTitle, &id)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("title: ", id, articleTitle)
	}

	rows.Close()

	row := db.QueryRow("select id, articleTitle from news where id = 33")

	if err != nil {
		fmt.Println(err)
		return
	}

	var id int
	var articleTitle string

	err = row.Scan(&id, &articleTitle)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(id, articleTitle)

	/*
		result, err := db.Exec("insert into news(`articleTitle`) values(?)", "Host your Golang app on Digital Ocean with Dokku")

		if err != nil {
			fmt.Println(err)
			return
		}

		affected, err := result.RowsAffected()

		if err != nil {
			fmt.Println(err)
			return
		}

		lastID, err := result.LastInsertId()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(affected, lastID)
	*/

	/*
		result, err := db.Exec("update news set articleTitle=? where id = ?", "Google", 23)

		if err != nil {
			fmt.Println(err)
			return
		}

		affected, err := result.RowsAffected()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("update affected", affected)
	*/

}
