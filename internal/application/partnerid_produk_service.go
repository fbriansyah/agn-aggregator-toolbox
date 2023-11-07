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

func (s *Service) UpdatePartnerProduk(ctx context.Context, data domain.PartnerProdukDomain) (domain.PartnerProdukDomain, error) {
	_, err := s.db.UpdatePartnerIDProduk(ctx, mariadb.UpdatePartnerIDProdukParams{
		IdpartnerProduk: data.IdpartnerProduk,
		Idproduk:        util.ValidNullInt32(data.Idproduk),
		Idmerchant:      util.ValidNullInt32(data.Idmerchant),
		Idpartner:       util.ValidNullInt32(data.Idpartner),
		Prioritas:       util.ValidNullInt32(data.Prioritas),
		Status:          util.ValidNullInt32(data.Status),
	})
	if err != nil {
		return domain.PartnerProdukDomain{}, err
	}

	return data, nil
}

func (s *Service) GetPartnerProduk(ctx context.Context, idpartnerproduk int32) (domain.PartnerProdukDomain, error) {
	partner, err := s.db.GetPartnerIDProduk(ctx, idpartnerproduk)
	if err != nil {
		return domain.PartnerProdukDomain{}, err
	}

	var pp domain.PartnerProdukDomain
	pp.FromMPartneridProduk(partner)

	return pp, nil
}

func (s *Service) GetPartners(ctx context.Context) ([]domain.PartnerDomain, error) {
	partners, err := s.db.GetPartners(ctx)
	if err != nil {
		return []domain.PartnerDomain{}, err
	}

	listPartner := make([]domain.PartnerDomain, len(partners))

	for i, p := range partners {
		var partnerDomain domain.PartnerDomain
		partnerDomain.FromMPartner(p)
		listPartner[i] = partnerDomain
	}

	return listPartner, nil
}
