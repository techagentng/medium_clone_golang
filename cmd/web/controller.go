package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/techagentng/medium/model"
	"github.com/techagentng/medium/pkg/helpers"
	"net/http"
	"time"
)
// Handler for homepage
//var templateFiles = []string{
//	"./assets/html/home.page.gohtml",
//	"/base.layout.gohtml",
//	"./assets/html/footer.partial.gohtml",
//}
func (app *application) home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.page.gohtml", gin.H{
		"title": "Main website",
	})
// blog := pkg.Blog{}
	//c.JSON(200,gin.H{"message":"welcome"})
}
func (app *application) getForm(c *gin.Context) {
	c.HTML(http.StatusOK, "getForm.page.gohtml", gin.H{
		"title": "Main website",
	})
	// blog := pkg.Blog{}
	//c.JSON(200,gin.H{"message":"welcome"})
}

// Handler for inserting post data into the database
func (app *application) registerUser(c *gin.Context) {
	c.HTML(http.StatusOK, "getForm.page.gohtml", gin.H{
		"title": "Main website",
	})
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil{
		panic(err)
	}

// Init the struct to populate form data into database
	var user = &model.Users{
		Firstname: firstname,
		Lastname: lastname,
		Username: username,
		Email: email,
		Password: hashedPassword,
		CreatedAt: time.Now().Format("2006-January-02"),
		LastLogin: time.Now().Format("2006-January-02"),
	}
	app.user.CreateUser(user)
	c.Redirect(http.StatusSeeOther, "/getForm")
}

func (app *application) viewAllPost(c *gin.Context)  {
	var blog = model.Post{}
	users, err := blog.GetAllContents()
	if err != nil{
		app.errorLog.Println("Can't read from the database")
	}
	fmt.Println(users)
}