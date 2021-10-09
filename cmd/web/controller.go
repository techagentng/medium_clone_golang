package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/techagentng/medium/cmd/model"
	"github.com/techagentng/medium/cmd/web/helpers"
	"net/http"
	"time"
)

func (app *application) home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.page.gohtml", nil)
}

func (app *application) RegisterPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "registerUser.page.gohtml", gin.H{
		"title": "Main website",
	})
}

// Handler for inserting post data into the psql
func (app *application) registerUser(c *gin.Context) {
	var user = &model.User{}
	// Init the struct to populate form data into psql
	user.Email = c.PostForm("email")
	user.FirstName = c.PostForm("firstname")
	user.LastName = c.PostForm("lastname")
	user.Password = c.PostForm("password")
	user.Confirm = c.PostForm("confirm")

	user.ID = uuid.New().String()
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		app.errorLog.Printf(err.Error())
		c.Redirect(http.StatusFound, "/register")
		return
	}

	user.Password = hashedPassword

	_, err = app.user.SaveUser(user)

	if err != nil {
		app.errorLog.Printf(err.Error())
		c.Redirect(http.StatusFound, "/signup")
		//c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	_, err = app.user.SaveUser(user)

	if err != nil {
		app.errorLog.Printf(err.Error())
		c.Redirect(http.StatusFound, "/register")
		//c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
}

func (app *application) viewAllPost(c *gin.Context)  {
	var blog = model.Post{}
	users, err := blog.GetAllContents()
	if err != nil{
		app.errorLog.Println("Can't read from the psql")
	}
	fmt.Println(users)
}
func (app *application) LoginPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.page.gohtml", nil)
}

func (app *application) LoginUser(c *gin.Context)  {
	var blog = model.Post{}
	users, err := blog.GetAllContents()
	if err != nil{
		app.errorLog.Println("Can't read from the psql")
	}
	fmt.Println(users)
}