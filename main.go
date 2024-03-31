package main

import (
	"fmt"
	"test/golang/config"
	"test/golang/module/post"
	"test/golang/module/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("tes project")

	// Connect to the database
	config.Connect()

	dbConnection := config.DB

	// book init modul
	userRepository := user.NewRepository(dbConnection)
	userService := user.NewService(userRepository)

	postRepository := post.NewRepository(dbConnection)
	postService := post.NewService(postRepository)

	// create fiber app
	app := fiber.New()

	// create book api
	user.NewApi(app, userService)
	post.NewApi(app, postService)

	// port
	app.Listen(":3000")

}
