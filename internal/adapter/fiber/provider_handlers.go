package fiber

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetProviderCreateForm render pages/profider/create-form
func (a *FiberAdapter) GetProviderCreateForm(c *fiber.Ctx) error {
	kodeProduk := c.Query("kode_product")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	product, err := a.service.GetProdukByCode(ctx, kodeProduk)
	if err != nil {
		return err
	}

	return c.Render("pages/provider/create-form", fiber.Map{
		"KodeProduk": kodeProduk,
		"Product":    product,
	})
}

func (a *FiberAdapter) GetListProvider(c *fiber.Ctx) error {
	kodeProduk := c.Query("kode_product", "")
	if kodeProduk == "" {
		return fiber.ErrBadRequest
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	providers, err := a.service.GetProductProvider(ctx, kodeProduk)
	if err != nil {
		return err
	}

	return c.Render("pages/provider/list-provider", fiber.Map{
		"Providers": providers,
		"Product": fiber.Map{
			"KodeProduk": kodeProduk,
		},
	})
}
