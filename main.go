package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		log.Println(c.Method(), c.Path())
		return c.Next()
	})
	app.Static("/", "./public")
	app.Get("/auto.pac", func(c *fiber.Ctx) error {
		c.Response().Header.Set("Content-Type", "application/x-ns-proxy-autoconfig")
		return c.SendString("auto.pac")
	})

	// API
	api := app.Group("/api")
	api.Use(cors.New())
	api.Get("/hostList", func(c *fiber.Ctx) error {
		return c.SendString("hostList")
	})
	api.Post("/updateHost", func(c *fiber.Ctx) error {
		return c.SendString("updateHost")
	})
	api.Post("/deleteHost", func(c *fiber.Ctx) error {
		return c.SendString("deleteHost")
	})
	api.Get("/ruleList", func(c *fiber.Ctx) error {
		return c.SendString("ruleList")
	})
	api.Post("/updateRule", func(c *fiber.Ctx) error {
		return c.SendString("updateRule")
	})
	api.Post("/deleteRule", func(c *fiber.Ctx) error {
		return c.SendString("deleteRule")
	})

	log.Fatal(app.Listen(fmt.Sprintf("127.0.0.1:%s", Port)))
}
