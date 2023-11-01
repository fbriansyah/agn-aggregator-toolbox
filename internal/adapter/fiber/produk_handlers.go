package fiber

import (
	"context"
	"time"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
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

func (a *FiberAdapter) GetProducts(c *fiber.Ctx) error {
	kodeProduk := c.Query("kode-produk")
	namaProduk := c.Query("nama-produk")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	products, err := a.service.GetProducts(ctx, domain.ProductDomain{
		KodeProduk: kodeProduk,
		NamaProduk: namaProduk,
	})
	if err != nil {
		return err
	}
	return c.Render("pages/product/list-product", fiber.Map{
		"Products": products,
	})
}
