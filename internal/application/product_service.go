package application

import (
	"context"
	"errors"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/adapter/mariadb"
	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application/domain"
	"github.com/fbriansyah/agn-aggregator-toolbox/util"
)

// ListProduct show all product without filter and limit
func (s *Service) ListProduct(ctx context.Context) ([]domain.ProductDomain, error) {
	products, err := s.db.ListProduk(ctx)
	if err != nil {
		return []domain.ProductDomain{}, err
	}

	listProduct := []domain.ProductDomain{}

	for _, product := range products {
		var p domain.ProductDomain
		p.FromMProduct(product)
		listProduct = append(listProduct, p)
	}

	return listProduct, nil
}

func (s *Service) GetProdukByCode(ctx context.Context, kodeProduk string) (domain.ProductDomain, error) {
	mProduct, err := s.db.GetProductByKode(ctx, kodeProduk)
	if err != nil {
		return domain.ProductDomain{}, err
	}

	produk := domain.ProductDomain{}
	produk.FromMProduct(mProduct)

	return produk, nil
}

func (s *Service) CreateProduct(ctx context.Context, produk domain.ProductDomain) (domain.ProductDomain, error) {
	arg := mariadb.CreateProductParams{
		KodeProduk:         produk.KodeProduk,
		NamaProduk:         util.ValidNullString(produk.NamaProduk),
		NamaProdukDisplay:  util.ValidNullString(produk.NamaProdukDisplay),
		NamaStruk:          util.ValidNullString(produk.NamaStruk),
		NamaStrukSingkatan: util.ValidNullString(produk.NamaStrukSingkatan),
		ProdukAlias:        util.ValidNullString(produk.ProdukAlias),
		ProdukGroup:        util.ValidNullString(produk.ProdukGroup),
		LabelIdpel:         util.ValidNullString(produk.LabelIdpel),
	}

	result, err := s.db.CreateProduct(ctx, arg)
	if err != nil {
		return domain.ProductDomain{}, nil
	}

	lastId, err := result.RowsAffected()
	if err != nil {
		return domain.ProductDomain{}, err
	}

	if lastId < int64(0) {
		return domain.ProductDomain{}, errors.New("error creating product")
	}

	return produk, nil
}

func (s *Service) GetProductProvider(ctx context.Context, kodeProduk string) ([]domain.ProdukProviderDomain, error) {
	providers, err := s.db.ListProviderProduk(ctx, util.ValidNullString(kodeProduk))
	if err != nil {
		return []domain.ProdukProviderDomain{}, err
	}
	listProvider := []domain.ProdukProviderDomain{}

	for _, provider := range providers {
		prov := domain.ProdukProviderDomain{}
		prov.FromMProviderProduk(provider)

		listProvider = append(listProvider, prov)
	}

	return listProvider, nil
}

func (s *Service) GetProducts(ctx context.Context, produk domain.ProductDomain) ([]domain.ProductDomain, error) {
	where := "1=1"
	var qw util.QueryWhere
	if produk.KodeProduk != "" {
		qw.Equal("KODE_PRODUK", produk.KodeProduk)
	}
	if produk.NamaProduk != "" {
		qw.Like("NAMA_PRODUK", produk.NamaProduk)
	}
	if !qw.IsEmpty() {
		where = qw.Build()
	}
	listProduct, err := s.db.GetProducts(ctx, where)
	if err != nil {
		return []domain.ProductDomain{}, nil
	}

	var products []domain.ProductDomain
	for _, prod := range listProduct {
		var product domain.ProductDomain
		product.FromMProduct(prod)
		products = append(products, product)
	}

	return products, nil
}
