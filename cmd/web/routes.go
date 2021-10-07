package main

import (
	"github.com/gin-gonic/gin"
)


func (app *application) Router(r *gin.Engine) *gin.Engine {
	r.LoadHTMLGlob("./assets/html/*")
	r.Static("/static", "./assets/static")


	r.GET("/", app.home)
	r.GET("/login", app.getForm)
	r.POST("/register", app.registerUser)
	//r.POST("/assets", app.viewAllPost)
	return r
}