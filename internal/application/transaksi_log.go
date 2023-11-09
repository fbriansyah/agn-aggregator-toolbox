package application

import (
	"context"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
	"github.com/fbriansyah/agn-aggregator-toolbox/util"
)

func (s *Service) GetTransaksiLogs(ctx context.Context, transaksiLog domain.TransaksiLog) ([]domain.TransaksiLog, error) {
	where := "1=1"

	var qw util.QueryWhere

	if transaksiLog.Blth != "" {
		qw.Equal("tl.blth", transaksiLog.Blth)
	}

	if transaksiLog.KodeProduk != "" {
		qw.Equal("mpp.KODE_PRODUK", transaksiLog.KodeProduk)
	}

	if transaksiLog.TglWaktu != "" {
		qw.Date("tl.tgl_waktu", transaksiLog.TglWaktu)
	}

	if !qw.IsEmpty() {
		where = qw.Build()
	}

	transakiLogs, err := s.db.GetTransaksiLog(ctx, where)
	if err != nil {
		return []domain.TransaksiLog{}, err
	}
	logs := make([]domain.TransaksiLog, len(transakiLogs))
	for i, l := range transakiLogs {
		var log domain.TransaksiLog
		log.FromModel(l)

		logs[i] = log
	}

	return logs, nil
}
