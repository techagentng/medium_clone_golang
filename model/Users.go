package model

import (
	"database/sql"
	_ "fmt"
)

type Users struct {
	UserId string
	Username string
	Firstname string
	Lastname string
	Followers [] string
	Post []string
	Dob string
	Password string
	Email string
	CreatedAt string
	LastLogin string
}

type UsersInter interface {
	CreateUser(u *Users)

}

type Database struct {
	Db  *sql.DB
}


// CreateUser method inserts form inputs to the posts table
func (d  *Database)CreateUser(u *Users)  {
	_, err := d.Db.Exec(`INSERT INTO users(UserId,Firstname,Followers,Password,Email,Username,CreatedAt,LastLogin) VALUES ($1,$2,$3,$4,$5,$6,$7)
`,u.UserId,u.Firstname,u.Followers,u.Password,u.Email,u.Username,u.CreatedAt,u.LastLogin)
	if err != nil{
		panic(err)
	}
}


//// Login method to login a user
//func (u *Users) Login()  {
//
//}