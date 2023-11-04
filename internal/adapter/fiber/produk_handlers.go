package fiber

import (
	"context"
	"strconv"
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

// SaveProduk recieve post form and create produk from the post data. After success save,
// it will redirects to /detail-product?kode_product={{.KodeProduk}}
func (a *FiberAdapter) SaveProduk(c *fiber.Ctx) error {
	productDomain := domain.ProductDomain{
		KodeProduk:         c.FormValue("kode-produk"),
		NamaProduk:         c.FormValue("nama-produk"),
		NamaProdukDisplay:  c.FormValue("nama-produk-display"),
		NamaStruk:          c.FormValue("nama-struk"),
		NamaStrukSingkatan: c.FormValue("nama-struk-singkatan"),
		ProdukAlias:        c.FormValue("produk-alias"),
		ProdukGroup:        c.FormValue("produk-group"),
		LabelIdpel:         c.FormValue("label-id"),
		ProdukDate:         time.Now(),
		Status:             1,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	_, err := a.service.CreateProduct(ctx, productDomain)
	if err != nil {
		return err
	}

	return c.Redirect("/detail-product?kode_product=" + productDomain.KodeProduk)

}

// UpdateProduk recieve post form and update the product based on form data
func (a *FiberAdapter) UpdateProduk(c *fiber.Ctx) error {

	s := c.FormValue("status")
	status, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return err
	}

	productDomain := domain.ProductDomain{
		KodeProduk:         c.FormValue("kode-produk"),
		NamaProduk:         c.FormValue("nama-produk"),
		NamaProdukDisplay:  c.FormValue("nama-produk-display"),
		NamaStruk:          c.FormValue("nama-struk"),
		NamaStrukSingkatan: c.FormValue("nama-struk-singkatan"),
		ProdukAlias:        c.FormValue("produk-alias"),
		ProdukGroup:        c.FormValue("produk-group"),
		LabelIdpel:         c.FormValue("label-id"),
		ProdukDate:         time.Now(),
		Status:             int32(status),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	produk, err := a.service.UpdateProduct(ctx, productDomain)
	if err != nil {
		return err
	}

	return c.Redirect("/detail-product?kode_product=" + produk.KodeProduk)
}
