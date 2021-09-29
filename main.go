package main

import (
	"go-authentication-boilerplate/database"
	"go-authentication-boilerplate/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// func init() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		fmt.Println("Error loading .env file")
// 	}
// }

func main() {

	database.ConnectToDB()
	app := CreateInstance()

	app.Use(cors.New())

	router.SetupRoutes(app)

	app.Listen(":3000")
}

func CreateInstance() *fiber.App {
	app := fiber.New()

	return app
}
