package post

import (
	"test/golang/domain"

	"github.com/gofiber/fiber/v2"
)

type api struct {
	postService domain.PostService
}

func NewApi(app *fiber.App, postService domain.PostService) {
	api := api{postService}

	apiRoute := app.Group("/api")

	postRoute := apiRoute.Group("/post")

	// api books
	postRoute.Post("/", api.createPost)
	postRoute.Get("/finduser/:id", api.findUser)

}

func (a *api) createPost(c *fiber.Ctx) error {
	var post domain.Post

	c.BodyParser(&post)

	dataset := a.postService.Store(post)

	return c.JSON(dataset)
}

func (a *api) findUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	dataset := a.postService.FindUser(int64(id))

	return c.JSON(dataset)

}
