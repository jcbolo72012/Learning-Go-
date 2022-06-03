package main

import (
	"fmt"
	// "log"
	// "strconv"
	"github.com/gofiber/fiber/v2"
)

type Seller struct{
	ID string `json:"ID"`
	Name string `json:"Name"`
	ASINS string `json:"ASINS"`
	Rating float64 `json:"rating"`
}

// func GetSellers(c *fiber.Ctx) {
// 	c.JSON(sellers)
// }

// func GetSeller(c *fiber.Ctx) {
// 	id := c.Params("id")
// 	for _, item := range sellers {
// 		if item.ID == id {
// 			c.JSON(item)
// 			return
// 		}
// 	}
// }

// func CreateSeller(c *fiber.Ctx) {

// 	seller := new(Seller)
// 	if err := c.BodyParser(seller); err != nil {
// 		c.Status(503).Send(err)
// 		return
// 	}
// 	sellers.append(sellers, seller)
// 	c.JSON(seller)
	
// }

// func DeleteSeller(c *fiber.Ctx) {
// 	id := c.Params("id")
// 	for index, item := range sellers {
// 		if item.ID == id {
// 			sellers = append(sellers[:index], sellers[index+1:]...)
// 			break
// 		}
// 	}
// 	c.JSON(sellers)
// }


var sellers[]Seller

// func setupRoutes(app *fiber.App) {

// 	app.Get("/api/v1/main", GetSellers)
// 	app.Get("/api/v1/main/:id", GetSeller)
// 	app.Post("/api/v1/main", CreateSeller)
// 	app.Delete("/api/v1/main/:id", DeleteSeller)
// }

func main() {
	app := fiber.New()

	sellers = append(sellers, Seller{ID: "1", Name:"Charlie", ASINS: "B0009999", Rating: 5})

	app.Get("/api/v1/get_sellers", func(c *fiber.Ctx) error{
		return c.JSON(sellers)
	})

	app.Get("/api/v1/get_seller", func(c *fiber.Ctx) error{
		id := c.Params("id")
		for _, item := range sellers {
			if item.ID == id {
				return c.JSON(item)
			}
		}
		return nil
	})

	app.Post("/api/v1/create_seller", func(c *fiber.Ctx) error{
		seller := new(Seller)

		if err := c.BodyParser(seller); err != nil {
			return err
		}
		sellers = append(sellers, *seller)
		return c.JSON(sellers)
	})

	app.Delete("/api/v1/delete_seller", func(c *fiber.Ctx) error{
		id := c.Params("id")
		for index, item := range sellers {
			if item.ID == id {
				sellers = append(sellers[:index], sellers[index+1:]...)
				return c.JSON(sellers)
			}
		}
		return c.JSON(sellers)
	})
	// setupRoutes(app)
	fmt.Printf("Starting server at port 3000\n")
	app.Listen(":3000")
}