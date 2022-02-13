package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupMoviesRoutes(app *fiber.App, tokenKey string) {
	s := start(tokenKey)
	grp := app.Group("/movies")
	grp.Get("/", s.SearchMovieHandler)
}

func SetupUsersRoutes(app *fiber.App, tokenKey string) {
	s := start(tokenKey)
	grp := app.Group("/users")
	grp.Post("/", s.CreateUserHandler)
	grp.Post("/login", s.LoginHandler)
	app.Use(jwtMiddleware(tokenKey))

	grp.Post("/wishlist", s.WishListHandler)
}
