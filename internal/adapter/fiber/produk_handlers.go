package fiber

import (
	"context"
	"time"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
	"github.com/gofiber/fiber/v2"
)

// ListProduk render pages/product/index
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

// GetProducts render pages/product/list-product
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

// CreateProdukFormPage render pages/product/create-form
func (a *FiberAdapter) CreateProdukFormPage(c *fiber.Ctx) error {
	return c.Render("pages/product/create-form", nil)
}

// DetailProductPage render pages/product/detail-product
func (a *FiberAdapter) DetailProductPage(c *fiber.Ctx) error {
	kodeProduk := c.Query("kode_product", "")
	if kodeProduk == "" {
		return fiber.ErrBadRequest
	}
	product, err := a.service.GetProdukByCode(context.Background(), kodeProduk)
	if err != nil {
		return err
	}

	providers, err := a.service.GetProductProvider(context.Background(), kodeProduk)
	if err != nil {
		return err
	}

	return c.Render("pages/product/detail-product", fiber.Map{
		"PageTitle": "Detail Product",
		"Product":   product,
		"Providers": providers,
	})
}
