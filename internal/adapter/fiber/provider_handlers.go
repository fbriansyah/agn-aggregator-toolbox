package fiber

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
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

// GetListProvider render pages/provider/list-provider
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

// ProdukProviderIndex render pages/provider/index
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

	// all partners registered in database
	listPartner, err := a.service.GetPartners(ctx)

	return c.Render("pages/provider/index", fiber.Map{
		"IDProvider":  idProvider,
		"PageTitle":   "Provider",
		"Provider":    provider,
		"Idproduk":    provider.Idproduk,
		"Partners":    partners,
		"ListPartner": listPartner,
	})
}

func (a *FiberAdapter) GetProviderEditor(c *fiber.Ctx) error {
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
	return c.Render("pages/provider/edit-form", fiber.Map{
		"Provider": provider,
	})
}

// SaveProvider recieve provider data from create-form and save it to database.
// If success inserted data, it will redirects to /provider?id={{.Idproduk}}
func (a *FiberAdapter) SaveProvider(c *fiber.Ctx) error {
	fee := c.FormValue("fee_produk")
	port := c.FormValue("class_port")
	timeOut := c.FormValue("class_timeout")

	rp_fee, err := strconv.ParseInt(fee, 10, 32)
	if err != nil {
		return err
	}

	classPort, err := strconv.ParseInt(port, 10, 32)
	if err != nil {
		return err
	}

	classTimeout, err := strconv.ParseInt(timeOut, 10, 32)
	if err != nil {
		return err
	}

	providerDomain := domain.ProdukProviderDomain{
		KodeProduk:         c.FormValue("kode_produk"),
		KodeProdukProvider: c.FormValue("kode_produk_provider"),
		NamaProduk:         c.FormValue("nama_produk"),
		Kategori:           c.FormValue("kategori"),
		Provider:           c.FormValue("provider"),
		Rctype:             c.FormValue("rctype"),
		ProviderBank:       c.FormValue("provider_bank"),
		FeeProduk:          int32(rp_fee),
		ClassRpc:           c.FormValue("class_rpc"),
		ClassName:          c.FormValue("class_name"),
		ClassTipe:          c.FormValue("class_tipe"),
		ClassIp:            c.FormValue("class_ip"),
		ClassPort:          int32(classPort),
		ClassPath:          c.FormValue("class_path"),
		ClassSetting:       c.FormValue("class_setting"),
		ClassTimeout:       int32(classTimeout),
		Screenname:         c.FormValue("screenname"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	provider, err := a.service.CreateProvider(ctx, providerDomain)
	if err != nil {
		return err
	}

	return c.Redirect(fmt.Sprintf("/provider?id=%d", provider.Idproduk))
}

func (a FiberAdapter) UpdateProvider(c *fiber.Ctx) error {
	id := c.FormValue("idproduk")
	fee := c.FormValue("fee_produk")
	port := c.FormValue("class_port")
	timeOut := c.FormValue("class_timeout")

	idproduk, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return err
	}

	rp_fee, err := strconv.ParseInt(fee, 10, 32)
	if err != nil {
		return err
	}

	classPort, err := strconv.ParseInt(port, 10, 32)
	if err != nil {
		return err
	}

	classTimeout, err := strconv.ParseInt(timeOut, 10, 32)
	if err != nil {
		return err
	}

	providerDomain := domain.ProdukProviderDomain{
		KodeProduk:         c.FormValue("kode_produk"),
		KodeProdukProvider: c.FormValue("kode_produk_provider"),
		NamaProduk:         c.FormValue("nama_produk"),
		Kategori:           c.FormValue("kategori"),
		Provider:           c.FormValue("provider"),
		Rctype:             c.FormValue("rctype"),
		ProviderBank:       c.FormValue("provider_bank"),
		FeeProduk:          int32(rp_fee),
		ClassRpc:           c.FormValue("class_rpc"),
		ClassName:          c.FormValue("class_name"),
		ClassTipe:          c.FormValue("class_tipe"),
		ClassIp:            c.FormValue("class_ip"),
		ClassPort:          int32(classPort),
		ClassPath:          c.FormValue("class_path"),
		ClassSetting:       c.FormValue("class_setting"),
		ClassTimeout:       int32(classTimeout),
		Screenname:         c.FormValue("screenname"),
		Idproduk:           int32(idproduk),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	_, err = a.service.UpdateProvider(ctx, providerDomain)
	if err != nil {
		return err
	}

	return c.Redirect(fmt.Sprintf("/provider?id=%d", idproduk))
}
