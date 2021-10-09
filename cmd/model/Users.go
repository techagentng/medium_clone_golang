package model

import (
	"database/sql"
	"fmt"
	_ "fmt"
)

type User struct {
	ID			string	`json:"id,omitempty"`
	FirstName	string	`json:"firstname,omitempty"`
	LastName	string	`json:"lastname,omitempty"`
	Email		string	`json:"email,omitempty"`
	Password	string	`json:"password,omitempty"`
	Confirm		string	`json:"confirm,omitempty"`
	CreatedAt	int64	`json:"created_at,omitempty"`
	UpdatedAt	int64	`json:"updated_at,omitempty"`
}

type UserModel struct {
	Db  *sql.DB
}

const (
	DBTABLE = "users"
)

func (model *UserModel) GetUserByEmail(email string) (bool, User) {
	var user User
	row := model.Db.QueryRow(fmt.Sprintf("SELECT id, email, password FROM %s WHERE email = $1",DBTABLE), email)
	if row.Err() != nil {
		return false, user
	}
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return false, user
	}
	return true, user
}

func (model *UserModel) SaveUser(user *User) (string, error) {
	fmt.Println(model.Db, "dftgyhuji")
	stmt, err := model.Db.Prepare(fmt.Sprintf("INSERT INTO %s(id, firstname, lastname, email, password, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7)", DBTABLE))
	if err != nil {
		return "", err
	}
	_, err = stmt.Exec(user.ID, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

func (model *UserModel) GetUserById(id string) (User, error){
	var user User
	row := model.Db.QueryRow(fmt.Sprintf("SELECT id, email, firstname, lastname, created_at FROM %s WHERE id = $1",DBTABLE), id)
	if row.Err() != nil {
		return user, row.Err()
	}
	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}