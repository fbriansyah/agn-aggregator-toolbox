package fiber

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
	"github.com/fbriansyah/agn-aggregator-toolbox/util"
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

// SavePartnerProduk insert data to partner produk
func (a *FiberAdapter) SavePartnerProduk(c *fiber.Ctx) error {
	idPartner, err := util.StrToInt32(c.FormValue("idpartner"))
	if err != nil {
		return err
	}
	idMerchant, err := util.StrToInt32(c.FormValue("idmerchant"))
	if err != nil {
		return err
	}
	idProduk, err := util.StrToInt32(c.FormValue("idproduk"))
	if err != nil {
		return err
	}
	prioritas, err := util.StrToInt32(c.FormValue("prioritas"))
	if err != nil {
		return err
	}
	status, err := util.StrToInt32(c.FormValue("status"))
	if err != nil {
		return err
	}

	ppDomain := domain.PartnerProdukDomain{
		Idpartner:  idPartner,
		Idmerchant: idMerchant,
		Idproduk:   idProduk,
		Prioritas:  prioritas,
		Status:     status,
	}

	fmt.Println(ppDomain)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	partnerProduk, err := a.service.CreatePartnerProduk(ctx, ppDomain)
	if err != nil {
		return err
	}

	return c.Redirect(fmt.Sprintf("/provider?id=%d", partnerProduk.Idproduk))
}

// UpdatePartnerProduk update partner produk data
func (a *FiberAdapter) UpdatePartnerProduk(c *fiber.Ctx) error {
	return fiber.NewError(fiber.ErrNotFound.Code, "UpdatePartnerProduk not yet implemented")
}
