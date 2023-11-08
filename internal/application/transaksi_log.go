package application

import (
	"context"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
)

func (s *Service) GetTransaksiLogs(ctx context.Context, transaksiLog domain.TransaksiLog) ([]domain.TransaksiLog, error) {
	where := "1=1"

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
