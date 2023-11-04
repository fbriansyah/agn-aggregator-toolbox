package fiber

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetPartnerProdukForm render pages/partner-produk/create-form.html
func (a *FiberAdapter) GetPartnerProdukForm(c *fiber.Ctx) error {
	id := c.Query("idproduk")
	return c.Render("pages/partner-produk/create-form", fiber.Map{
		"Idproduk": id,
	})
}

// GetListPartnerProduk render pages/partner-produk/list-partner
func (a *FiberAdapter) GetListPartnerProduk(c *fiber.Ctx) error {
	id := c.Query("idproduk")

	idproduk, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	partnerProduks, err := a.service.GetListPartnerProduk(ctx, int32(idproduk))
	if err != nil {
		return err
	}

	return c.Render("pages/partner-produk/list-partner", fiber.Map{
		"Idproduk": id,
		"Partners": partnerProduks,
	})
}

// GetPartnerProdukEditForm render pages/partner-produk/edit-form.html
func (a *FiberAdapter) GetPartnerProdukEditForm(c *fiber.Ctx) error {
	id := c.Query("id")
	idPartnerProduk, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	partner, err := a.service.GetPartnerProduk(ctx, int32(idPartnerProduk))
	if err != nil {
		return err
	}

	return c.Render("pages/partner-produk/edit-form", fiber.Map{
		"Idproduk": partner.Idproduk,
		"Partner":  partner,
	})
}
