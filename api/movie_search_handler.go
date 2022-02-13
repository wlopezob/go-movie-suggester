package api

import "github.com/gofiber/fiber/v2"

func (ws *WebServices) SearchMovieHandler(c *fiber.Ctx) error {
	res, err := ws.search.Search(MovieFilter{})

	if err != nil {
		return fiber.NewError(400, "cannot bring movies")
	}
	if len(res) == 0 {
		return c.JSON([]interface{}{})
	}
	return c.JSON(res)
}
