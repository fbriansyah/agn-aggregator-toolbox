package fiber

import (
	"context"
	"time"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application"
	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
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
