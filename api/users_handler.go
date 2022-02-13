package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wlopezob/go-movie-suggester/pkg/models"
)

func (w *WebServices) CreateUserHandler(c *fiber.Ctx) error {
	var cmd CreateUserCMD
	err := c.BodyParser(&cmd)

	res, err := w.users.SaveUser(cmd)

	if err != nil {
		return fiber.NewError(400, "cannot create user")
	}
	signToken(w.tokenKey, res.ID)
	return c.Status(200).JSON(res)
}

func (w *WebServices) WishListHandler(c *fiber.Ctx) error {
	bearear := c.Get("Authorization")
	userId := extractUserIdFromJWT(bearear, w.tokenKey)
	if userId == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	var cmd models.WishMovieCMD
	c.BodyParser(&cmd)
	err := w.users.AddWishMovie(userId, cmd.MovieId, cmd.Comment)
	if err != nil {
		return fiber.NewError(400, "cannot add to the wishlist")
	}
	return c.JSON(struct {
		R string `json:"result"`
	}{
		R: "movie add to the wishlist",
	})
}

func (w *WebServices) ServeVideo(c *fiber.Ctx) error {
	c.Set("Content-Type", "video/mp4")
	err := c.SendFile("test.mp4", false)
	if err != nil {
		return fiber.NewError(500, "Cannot view video")
	}
	return nil
}

func (w *WebServices) LoginHandler(c *fiber.Ctx) error {
	var login models.LoginCMD
	err := c.BodyParser(&login)
	if err != nil {
		return fiber.NewError(400, "cannot parse params")
	}
	id := w.users.Login(login)
	if id == "" {
		return fiber.NewError(400, "user not found")
	}
	return c.JSON(struct {
		Token string `json:"token"`
	}{
		Token: signToken(w.tokenKey, id),
	})
}
