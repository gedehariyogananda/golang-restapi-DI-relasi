package user

import (
	"test/golang/domain"

	"github.com/gofiber/fiber/v2"
)

type api struct {
	userService domain.UserService
}

func NewApi(app *fiber.App, userService domain.UserService) {
	api := api{userService}

	apiRoute := app.Group("/api")

	userRoute := apiRoute.Group("/user")

	// api user
	userRoute.Get("/", api.findAll)
	userRoute.Get("/:id", api.findById)
	userRoute.Post("/", api.create)
	userRoute.Put("/:id", api.update)
	userRoute.Get("/:id/post", api.findAllPost)

}

func (a *api) findAll(c *fiber.Ctx) error {
	dataset := a.userService.FindAll()

	return c.JSON(dataset)
}

func (a *api) findById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	dataset := a.userService.FindByID(int64(id))

	return c.JSON(dataset)
}

func (a *api) create(c *fiber.Ctx) error {
	var user domain.User

	c.BodyParser(&user)

	dataset := a.userService.Store(user)

	return c.JSON(dataset)
}

func (a *api) update(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	var user domain.User

	c.BodyParser(&user)

	user.Id = int64(id)

	dataset := a.userService.Update(user)

	return c.JSON(dataset)
}

func (a *api) findAllPost(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	dataset := a.userService.FindAllPost(int64(id))

	return c.JSON(dataset)
}
