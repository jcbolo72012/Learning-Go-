package seller

import(
	"github.com/gofiber/fiber"
)


type Seller struct{
	ID string `json:"ID"`
	Name string `json:"Name"`
	ASINS string `json:"ASINS"`
	rating float64 `json:"rating"`
}

var sellers[]Seller


func GetSellers(c *fiber.Ctx) {
	c.JSON(sellers)
}

func GetSeller(c *fiber.Ctx) {
	id := c.Params("id")
	for _, item := range sellers {
		if item.ID == id {
			c.JSON(item)
			return
		}
	}
}

func CreateSeller(c *fiber.Ctx) {

	seller := new(Seller)
	if err := c.BodyParser(seller); err != nil {
		c.Status(503).Send(err)
		return
	}
	sellers.append(sellers, seller)
	c.JSON(seller)
	
}

func DeleteSeller(c *fiber.Ctx) {
	id := c.Params("id")
	for index, item := range sellers {
		if item.ID == id {
			sellers = append(sellers[:index], sellers[index+1:]...)
			break
		}
	}
	c.JSON(sellers)
}


