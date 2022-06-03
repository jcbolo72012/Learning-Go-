package main

import (
	"fmt"
	// "log"
	// "strconv"
	"github.com/gofiber/fiber/v2"
	"Learning-Go-/Go-API-Fiber/seller"
)


func setupRoutes(app *fiber.App) {

	app.Get("/api/v1/main", seller.GetSellers)
	app.Get("/api/v1/main/:id", seller.GetSeller)
	app.Post("/api/v1/main", seller.CreateSeller)
	app.Delete("/api/v1/main/:id", seller.DeleteSeller)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	fmt.Printf("Starting server at port 8000\n")
	app.Listen(":3000")
}