// @/main.go
package main

import (
	"log"

	"ncip48/go-fiber-ecommerce/config"
	"ncip48/go-fiber-ecommerce/handlers"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Connect()

	app.Get("/users", handlers.GetUsers)
	app.Get("/users/:id", handlers.GetUser)
	app.Post("/users", handlers.AddUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.RemoveUser)
	app.Post("/users/check", handlers.CheckPassword)
	app.Post("/auth/login", handlers.LoginAction)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	// Restricted Routes
	app.Get("/profile", handlers.GetProfile)

	log.Fatal(app.Listen(":3030"))
}
