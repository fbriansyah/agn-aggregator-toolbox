package fiber

import (
	"context"
	"strconv"
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

func (a *FiberAdapter) ProdukProviderIndex(c *fiber.Ctx) error {
	idProvider := c.Query("id")
	id, err := strconv.ParseInt(idProvider, 10, 32)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	provider, err := a.service.GetProviderDetail(ctx, int32(id))
	if err != nil {
		return err
	}

	partners, err := a.service.GetListPartnerProduk(ctx, provider.Idproduk)
	if err != nil {
		return err
	}

	return c.Render("pages/provider/index", fiber.Map{
		"IDProvider": idProvider,
		"PageTitle":  "Provider",
		"Provider":   provider,
		"Idproduk":   provider.Idproduk,
		"Partners":   partners,
	})
}
