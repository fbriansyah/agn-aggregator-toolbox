package fiber

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func (a *FiberAdapter) ListProduk(c *fiber.Ctx) error {
	products, err := a.service.ListProduct(context.Background())
	if err != nil {
		return err
	}

	return c.Render("pages/product/index", fiber.Map{
		"PageTitle":    "Product",
		"AddButtonUrl": "/add-product",
		"Products":     products,
	}, "partials/base")
}
