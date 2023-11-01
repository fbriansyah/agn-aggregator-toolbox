package fiber

import (
	"fmt"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/port"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type FiberAdapter struct {
	service port.ServicePort
	app     *fiber.App
}

func NewFiberAdapter(service port.ServicePort) *FiberAdapter {
	engine := html.New("./templates/fiber", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	return &FiberAdapter{
		service: service,
		app:     app,
	}
}

func (a *FiberAdapter) Run(address string) {
	a.router()
	fmt.Println(a.app.Listen(address))
}
