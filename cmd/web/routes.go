package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)


func (app *application) Router(r *gin.Engine) *gin.Engine {
	r.Use(static.Serve("/static", static.LocalFile("./assets/static", true)))
	r.LoadHTMLGlob("./assets/html/*")
	r.Static("/static", "./assets/static")


	r.GET("/", app.home)
	//r.GET("/register", app.RegisterPageHandler)
	r.POST("/register", app.registerUser)
	r.GET("/login", app.LoginPageHandler)
	r.POST("/login", app.LoginUser)
	//r.POST("/assets", app.viewAllPost)
	return r
}