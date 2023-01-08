package server

import (
	"github.com/atadzan/url-shortener/app/pkg/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupAndListen() {
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/r/:redirect", handler.Redirect)
	router.Get("/goly", handler.GetAllGolies)
	router.Get("/goly/:id", handler.GetGoly)
	router.Post("/goly", handler.CreateGoly)
	router.Patch("/goly", handler.UpdateGoly)
	router.Delete("/goly/:id", handler.DeleteGoly)

	router.Listen(":8001")
}
