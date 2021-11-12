package presenter

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/julhan07/fiber-service/middleware"
)

type presenter struct {
	route *fiber.App
}

func NewPresenterIndex(r *fiber.App) {

	p := &presenter{
		route: r,
	}
	p.routes()
}

func (p *presenter) routes() {

	p.route.Get("/", p.Welcome)
	p.route.Post("/user/register", middleware.AuthAge, p.GetUsers)
	p.route.Get("/user/list", middleware.AuthAge, p.GetUserDetail)
}

func (p *presenter) Welcome(c *fiber.Ctx) error {
	msg := fmt.Sprintf("👴 Hi Welcome %s My Age %s  ", c.Query("name"), c.Query("age"))
	return c.SendString(msg) // => 👴 julhan is 28 year
}
