package fiber

import "github.com/gofiber/fiber/v2"

// GetPartnerProdukForm render pages/partner-produk/create-form.html
func (a *FiberAdapter) GetPartnerProdukForm(c *fiber.Ctx) error {
	id := c.Query("idproduk")
	return c.Render("pages/partner-produk/create-form", fiber.Map{
		"Idproduk": id,
	})
}

func (a *FiberAdapter) GetListPartnerProduk(c *fiber.Ctx) error {
	// id := c.Query("idproduk")

	return fiber.ErrNotFound
}

func (a *FiberAdapter) GetPartnerProdukEditForm(c *fiber.Ctx) error {
	// id := c.Query("id")

	return fiber.ErrNotFound
}
