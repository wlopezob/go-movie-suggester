package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/wlopezob/go-movie-suggester/api"
	"github.com/wlopezob/go-movie-suggester/pkg/models"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			msg := ""
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				msg = e.Error()
			}
			if msg == "" {
				msg = "cannot process the http call"
			}
			err = ctx.Status(code).JSON(models.InternalError{
				Message: msg,
			})
			return nil
		},
	})

	// midlerware
	app.Use(recover.New())
	app.Use(logger.New())
	key := "tokenKey"
	api.SetupMoviesRoutes(app, key)
	api.SetupUsersRoutes(app, key)

	_ = app.Listen(":3000")
}
