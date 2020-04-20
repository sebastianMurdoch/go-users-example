package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastianMurdoch/go-users-example/config"
	"github.com/sebastianMurdoch/go-users-example/domain"
	"github.com/sebastianMurdoch/go-users-example/infrastructure"
	"os"
)

type App struct {
	Engine     *gin.Engine            `inject:"auto"`
	Repository domain.UsersRepository `inject:"auto"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	c := config.NewHerokuContainer()
	c.AddToDependencies(&infrastructure.UsersRepositoryImpl{})
	app := App{}
	c.InjectWithDependencies(&app)

	/* Status Ping */
	app.Engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	/* Get all users */
	app.Engine.GET("/users", func(c *gin.Context) {
		c.JSON(200, app.Repository.FindAll())
	})

	/* Post a new user*/
	app.Engine.POST("/users", func(c *gin.Context) {
		var user domain.User
		c.Bind(&user)
		app.Repository.Save(user)
		c.JSON(200, user)
	})

	app.Engine.Run(":"+port)
}