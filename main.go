package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/bdemirpolat/unit-test/models"
	"github.com/bdemirpolat/unit-test/repository"
	"github.com/gofiber/fiber/v2"
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

func createTable(db *sql.DB) (sql.Result, error) {
	res, err := db.Exec("CREATE TABLE `users` (`id` INTEGER PRIMARY KEY AUTOINCREMENT,`username` VARCHAR(64) NULL);")
	return res, err
}

func main() {
	db := connectDB()
	create := flag.Bool("create-table", false, "create table")
	flag.Parse()
	if *create {
		_, err := createTable(db)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("table created")
		return
	}

	userRepo := &repository.UserRepo{DB: db}
	app := fiber.New()
	app.Post("/users", func(c *fiber.Ctx) error {
		user := models.User{}
		err := c.BodyParser(&user)
		if err != nil {
			return err
		}
		err = userRepo.Create(user)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(user)
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
