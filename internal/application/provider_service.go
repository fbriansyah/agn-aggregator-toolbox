package application

import (
	"context"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/adapter/mariadb"
	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
	"github.com/fbriansyah/agn-aggregator-toolbox/util"
)

func (s *Service) GetProviderDetail(ctx context.Context, idproduk int32) (domain.ProdukProviderDomain, error) {
	provider, err := s.db.GetProviderByID(ctx, idproduk)
	if err != nil {
		return domain.ProdukProviderDomain{}, err
	}
	var productProvider domain.ProdukProviderDomain

	productProvider.FromMProviderProduk(provider)

	return productProvider, nil
}

// CreateProvider insert provider data to database
func (s *Service) CreateProvider(ctx context.Context, provider domain.ProdukProviderDomain) (domain.ProdukProviderDomain, error) {
	arg := mariadb.CreateProviderParams{
		KodeProduk:         util.ValidNullString(provider.KodeProduk),
		KodeProdukProvider: util.ValidNullString(provider.KodeProdukProvider),
		NamaProduk:         util.ValidNullString(provider.NamaProduk),
		Kategori:           util.ValidNullString(provider.Kategori),
		Provider:           util.ValidNullString(provider.Provider),
		Rctype:             util.ValidNullString(provider.Rctype),
		ProviderBank:       util.ValidNullString(provider.ProviderBank),
		FeeProduk:          util.ValidNullInt32(provider.FeeProduk),
		ClassRpc:           util.ValidNullString(provider.ClassRpc),
		ClassName:          util.ValidNullString(provider.ClassName),
		ClassTipe:          util.ValidNullString(provider.ClassTipe),
		ClassIp:            util.ValidNullString(provider.ClassIp),
		ClassPort:          util.ValidNullInt32(provider.ClassPort),
		ClassPath:          util.ValidNullString(provider.ClassPath),
		ClassSetting:       util.ValidNullString(provider.ClassSetting),
		Tipemodel:          util.ValidNullString(""),
		Screenname:         util.ValidNullString(provider.Screenname),
	}

	result, err := s.db.CreateProvider(ctx, arg)
	if err != nil {
		return domain.ProdukProviderDomain{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.ProdukProviderDomain{}, err
	}

	provider.Idproduk = int32(id)

	return provider, nil
}
