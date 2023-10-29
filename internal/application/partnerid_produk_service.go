package application

import (
	"context"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/adapter/mariadb"
	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
	"github.com/fbriansyah/agn-aggregator-toolbox/util"
)

func (s *Service) GetListPartnerProduk(ctx context.Context, idproduk int32) ([]domain.PartnerProdukDomain, error) {
	partnerProduks, err := s.db.ListPartnerIDProduk(ctx, util.ValidNullInt32(idproduk))
	if err != nil {
		return []domain.PartnerProdukDomain{}, err
	}
	listPartnerProduk := make([]domain.PartnerProdukDomain, len(partnerProduks))

	for i, pp := range partnerProduks {
		var partnerProduk domain.PartnerProdukDomain
		partnerProduk.FromMPartneridProduk(pp)

		listPartnerProduk[i] = partnerProduk
	}

	return listPartnerProduk, nil
}

func (s *Service) CreatePartnerProduk(ctx context.Context, data domain.PartnerProdukDomain) (domain.PartnerProdukDomain, error) {
	result, err := s.db.CreatePartnerIDProduk(ctx, mariadb.CreatePartnerIDProdukParams{
		Idproduk:   util.ValidNullInt32(data.Idproduk),
		Idmerchant: util.ValidNullInt32(data.Idmerchant),
		Idpartner:  util.ValidNullInt32(data.Idpartner),
		Prioritas:  util.ValidNullInt32(data.Prioritas),
	})
	if err != nil {
		return domain.PartnerProdukDomain{}, err
	}

	id, err := result.LastInsertId()
	if err == nil {
		data.IdpartnerProduk = int32(id)
	}

	return data, nil
}
