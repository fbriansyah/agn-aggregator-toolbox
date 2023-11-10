package fiber

import (
	"context"
	"time"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application"
	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
	"github.com/fbriansyah/agn-aggregator-toolbox/util"
	"github.com/gofiber/fiber/v2"
)

func (a *FiberAdapter) GetTransaksiLogs(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	blth := "202212" // time.Now().Format("200601")

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

	idLog, err := util.StrToInt32(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	log, err := a.service.GetTransaksiLogByID(ctx, idLog)

	return c.Render("", fiber.Map{
		"Log": log,
	})
}
