package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/techagentng/medium/cmd/interfaces"
	"github.com/techagentng/medium/cmd/model"
	"github.com/techagentng/medium/pkg/psql"
	"log"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	user    interfaces.User
}


var db, _ = psql.ConnectDb()
var app= &application{
	errorLog: log.New(os.Stderr,"ERROR\t",log.Ltime|log.Ldate|log.Lshortfile),
	infoLog: log.New(os.Stdout,"INFO\t",log.Ltime|log.Ldate),
	user: &model.UserModel{Db: db},

}

func main()  {
	r := gin.Default()
	app.Router(r)


	//staticFileServer := http.FileServer(http.Dir("../ui/static"))
	//r.StaticFS("/static/", http.Dir("."))
	//r.Handle("/static/",http.StripPrefix("/static",staticFileServer))

	//r.StaticFile("/favicon.ico", "./resources/favicon.ico")
	addr := flag.String("addr",":8000","Pass the network address")
	flag.Parse()
	err := r.Run(*addr)
	if err != nil {
		return
	}

}
