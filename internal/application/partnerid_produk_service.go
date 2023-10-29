package application

import (
	"context"

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
