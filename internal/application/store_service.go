package application

import (
	"context"
	"time"
)

const (
	STORE_PRODUK = "PRODUK"
)

func (s *Service) GetStoreData(storeName, value string) ([]map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	if storeName == STORE_PRODUK {
		return s.getProductStore(ctx, value)
	}

	return make([]map[string]any, 0), nil
}

func (s *Service) getProductStore(ctx context.Context, value string) ([]map[string]any, error) {
	products, err := s.db.GetProducts(ctx, "1=1")
	if err != nil {
		return []map[string]any{}, err
	}
	options := make([]map[string]any, len(products))

	for i, p := range products {
		options[i] = map[string]any{
			"Label":      p.NamaProduk.String,
			"Value":      p.KodeProduk,
			"IsSelected": (p.KodeProduk == value),
		}
	}
	return options, nil
}
