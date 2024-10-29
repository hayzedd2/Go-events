package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hayzedd2/Go-events/db"
	"github.com/hayzedd2/Go-events/utils"
)

type User struct {
	ID       int64
	UserName string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
	UserId   string
}
type UserLogin struct {
	ID       int64
	UserName string
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	u.UserId = uuid.New().String()
	query := "INSERT INTO users(email, username, password, userId) values(?,?,?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err

	}
	defer stmt.Close()
	if len(u.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, u.UserName, hashedPassword, u.UserId)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userId
	return nil
}

func (u *UserLogin) ValidateCredentials() (*User, error) {
	query := "SELECT id,userName,password,userId FROM users WHERE email =?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	var user User
	err := row.Scan(&user.ID, &user.UserName, &retrievedPassword, &user.UserId)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	passwordIsValid := utils.ComparePassword(u.Password, retrievedPassword)
	if !passwordIsValid {
		return nil, errors.New("invalid credentials")
	}
	return &user, nil
}

func GetUserByUserId(userId string) (*User, error) {
	query := "SELECT id, email, userName, userId FROM users WHERE userId = ?"
	row := db.DB.QueryRow(query, userId)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.UserName, &user.UserId)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
