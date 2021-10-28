package repository

import (
	"database/sql"

	"github.com/bdemirpolat/unit-test/models"
	"github.com/pkg/errors"
)

type UserRepository interface {
	Insert(user models.User) error
	Delete(userID int) error
	Update(userID int, user models.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (u userRepository) Insert(user models.User) error {
	statement, err := u.db.Prepare("INSERT INTO users (fisrtname, lastname, email) VALUES (?, ?, ?);")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(user.FirstName, user.LastName, user.Email)
	if err != nil {

		return errors.Wrap(err, "error insert user")
	}
	return err
}

func (u userRepository) Delete(userID int) error {
	statement, err := u.db.Prepare("delete from user where id =?;")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(userID)
	if err != nil {

		return errors.Wrap(err, "error delete user")
	}
	return nil
}

func (u userRepository) Update(userID int, user models.User) error {
	statement, err := u.db.Prepare("update user set firstname=?, lastname=?, email=? where id =?")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(user.FirstName, user.LastName, user.Email, userID)
	if err != nil {

		return errors.Wrap(err, "error update user")
	}
	return err
}
