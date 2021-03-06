package model

import (
	"database/sql"
	_ "fmt"
)

type Post struct {
	ID         string
	Title      string
	Body       string
	Author     string
	CreatedAt  string
	LastUpdate string
	LastLogin  string
}
var db *sql.DB
// CreateBlog method inserts form inputs to the posts table
func (p *Post) CreateBlog()  {
	_, err := db.Exec(`INSERT INTO posts(ID,Title,Body,Author,CreatedAt,LastUpdate) VALUES ($1,$2,$3,$4,$5,$6)
`, p.ID, p.Author, p.Title, p.Body, p.Author, p.CreatedAt, p.LastUpdate)
	if err != nil{
		panic(err)
	}
}
// GetAllContents from the post table
func (p *Post) GetAllContents()(blogPosts []Post, err error)   {
	rows, err := db.Query(`SELECT ID,Title,Body,Author,CreatedAt,LastUpdate FROM posts`)
	if err != nil{
		return
	}
	for rows.Next(){
		bp := Post{}
		err = rows.Scan(&p.ID, p.Author, p.Title, p.Body, p.Author, p.CreatedAt, p.LastUpdate)
		if err != nil{
			return
		}
		blogPosts = append(blogPosts,bp)
	}
	rows.Close()
	return

}