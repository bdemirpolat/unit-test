package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func connectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "file:database.db")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	fmt.Println(sayHello("Mertaaa"))

}

func AddMigration() {
	db := connectDB()

	defer db.Close()

	createAlbumTable := `CREATE TABLE user (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"firstname" TEXT,
		"lastname" TEXT,
		"email" TEXT,
	  );`

	_, err := db.Exec(createAlbumTable)

	if err != nil {
		log.Printf("%q: %s\n", err, createAlbumTable)
		return
	}

	log.Println("Migration finished...")
}
