package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

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
	fmt.Println(Sum(5, 5))
}

func Sum(a int, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

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
