package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Calculadora API do Britinho")
	})
	fmt.Println("teste")
	app.Get("/soma/:op1/:op2", soma)
	app.Listen(":3000")
}

func soma(c *fiber.Ctx) error {
	op1Str := c.Params("op1")
	op2Str := c.Params("op2")
	op1, err := strconv.ParseFloat(c.Params("op1"), 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("Parâmetro inválido:\":%s\"", op1Str))
	}
	op2, err := strconv.ParseFloat(c.Params("op2"), 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("Parâmetro inválido:\":%s\"", op2Str))
	}
	result := op1 + op2
	return c.SendString(fmt.Sprintf("%.2f", result))
}
