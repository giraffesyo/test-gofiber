package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Rect struct {
	length, width int
}

func (r Rect) Area() int {
	return r.length * r.width
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/rect/:length/:width", func(c *fiber.Ctx) error {
		length, _ := strconv.ParseInt(c.Params("length"), 10, 32)
		width, _ := strconv.ParseInt(c.Params("width"), 10, 32)
		r := Rect{int(length), int(width)}

		area := fmt.Sprintf("%d", r.Area())

		return c.SendString(area)
	})

	app.Listen(":3000")
}
