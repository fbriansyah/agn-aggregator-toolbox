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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// all partners registered in database
	partners, err := a.service.GetPartners(ctx)
	if err != nil {
		return err
	}

	return c.Render("pages/partner-produk/create-form", fiber.Map{
		"Idproduk": id,
		"Partners": partners,
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

	// all partners registered in database
	listPartner, err := a.service.GetPartners(ctx)
	if err != nil {
		return err
	}

	partnersOption := make([]map[string]any, len(listPartner))

	for i, p := range listPartner {
		partnersOption[i] = map[string]any{
			"Label":      p.NamaPartner,
			"Value":      p.IdPartner,
			"IsSelected": (p.IdPartner == partner.Idpartner),
		}
	}

	return c.Render("pages/partner-produk/edit-form", fiber.Map{
		"Idproduk":       partner.Idproduk,
		"Partner":        partner,
		"PartnersOption": partnersOption,
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
	idPartnerProduk, err := util.StrToInt32(c.FormValue("idpartner_produk"))
	if err != nil {
		return err
	}

	ppDomain := domain.PartnerProdukDomain{
		IdpartnerProduk: idPartnerProduk,
		Idpartner:       idPartner,
		Idmerchant:      idMerchant,
		Idproduk:        idProduk,
		Prioritas:       prioritas,
		Status:          status,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	_, err = a.service.UpdatePartnerProduk(ctx, ppDomain)
	if err != nil {
		return err
	}

	return c.Redirect(fmt.Sprintf("/provider?id=%d", ppDomain.Idproduk))
}
