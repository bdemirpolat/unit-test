package repository

import (
	"database/sql"
	"github.com/bdemirpolat/unit-test/models"
)

type UserRepository interface {
	Create(user models.User) error
	Delete(tableName string) error
}

type UserRepo struct {
	DB *sql.DB
}

func (u UserRepo) Create(user models.User) error {
	stmt, err := u.DB.Prepare("INSERT INTO users (username) VALUES (?);")
	if err != nil {
		return err
	}
	_,err = stmt.Exec(user.Username)
	return err
}

func (u UserRepo) Delete(tableName string) error {
	return nil
}
