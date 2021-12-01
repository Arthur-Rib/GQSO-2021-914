package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Calculadora API do Britinho")
	})
	fmt.Println("teste")
	app.Get("/soma/:op1/:op2", somaHandler)
	app.Listen(":3000")
}

func somaHandler(c *fiber.Ctx) error {
	op1Str := c.Params("op1")
	op2Str := c.Params("op2")
	result, err := soma(op1Str, op2Str)
	if err != nil {
		c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("Parâmetro inválido:\":%q\"", err))
	}
	return c.SendString(fmt.Sprintf("%.2f", result))
}
