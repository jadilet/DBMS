package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

const (
	dbName     = "onlineperevozka_development"
	dbUser     = "postgres"
	dbPassword = ""
)

func main() {
	var err error

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbInfo)

	PanicOnError(err)

	start := time.Now()
	rows, err := db.Query("select id, telephone from loads")

	PanicOnError(err)

	// Select Query
	var id int
	var telephone sql.NullString

	for rows.Next() {
		err = rows.Scan(&id, &telephone)

		PanicOnError(err)
		fmt.Println(id, telephone.String)
	}

	rows.Close()

	// Insert query
	var lastID int64
	err = db.QueryRow("INSERT INTO loads ( telephone, name, created_at, updated_at) VALUES($1, $2, $3, $4) RETURNING id", "0777285868", "Адилет Жолдошбеков", time.Now(), time.Now()).Scan(&lastID)

	PanicOnError(err)

	fmt.Println("Last inserted data id is ", lastID)

	// Update Query
	result, err := db.Exec("UPDATE loads set \"from\" = $1, \"to\" = $2, updated_at = $3 where id = 9", "Бишкек", "Жалалабат", time.Now())

	PanicOnError(err)

	affected, err := result.RowsAffected()
	PanicOnError(err)

	fmt.Println("Update Rows affected", affected)

	elapsed := time.Since(start)
	log.Printf("Job done in %s", elapsed)
}

// PanicOnError handles all errors
func PanicOnError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
