package fiber

import (
	"context"
	"fmt"
	"time"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application"
	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
	"github.com/fbriansyah/agn-aggregator-toolbox/util"
	"github.com/gofiber/fiber/v2"
)

const (
	BACK_LINK_KEY = "back-link"
)

func (a *FiberAdapter) GetTransaksiLogs(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	blth := time.Now().Format("200601")

	logs, err := a.service.GetTransaksiLogs(ctx, domain.TransaksiLog{
		Blth: blth,
	})
	if err != nil {
		return err
	}

	productOptions, err := a.service.GetStoreData(application.STORE_PRODUK, "")
	if err != nil {
		return err
	}

	return c.Render("pages/transaksi-logs/transaksi-logs", fiber.Map{
		"Logs":          logs,
		"BLTH":          blth,
		"PageID":        "transaksi-logs",
		"PageTitle":     "Transaksi Log",
		"ProdukOptions": productOptions,
	})
}

func (a *FiberAdapter) SearchTransaksiLogs(c *fiber.Ctx) error {
	kodeProduk := c.Query("produk")
	tglWaktu := c.Query("tgl_waktu")
	blth := c.Query("blth")
	idpel := c.Query("idpel")

	if blth == "" {
		blth = time.Now().Format("200601")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	c.Cookie(&fiber.Cookie{
		Name:  BACK_LINK_KEY,
		Value: fmt.Sprintf("/transaksi-logs/search?blth=%s&tgl_waktu=%s&produk=%s&idpel=%s", blth, tglWaktu, kodeProduk, idpel),
	})

	logs, err := a.service.GetTransaksiLogs(ctx, domain.TransaksiLog{
		Blth:       blth,
		TglWaktu:   tglWaktu,
		KodeProduk: kodeProduk,
		Idpel:      idpel,
	})
	if err != nil {
		return err
	}

	return c.Render("pages/transaksi-logs/list-log", fiber.Map{
		"Logs": logs,
	})
}

func (a *FiberAdapter) DetailTransaksiLogs(c *fiber.Ctx) error {
	id := c.Query("id")
	blth := c.Query("blth")

	idLog, err := util.StrToInt32(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	log, err := a.service.GetTransaksiLogByID(ctx, idLog)
	if err != nil {
		return err
	}
	if blth == "" {
		blth = time.Now().Format("200601")
	}

	return c.Render("pages/transaksi-logs/log-details", fiber.Map{
		"Log":  log,
		"Blth": blth,
	})
}

func (a *FiberAdapter) TransaksiLogBack(c *fiber.Ctx) error {
	backLink := c.Cookies(BACK_LINK_KEY)

	return c.Redirect(backLink)
}
