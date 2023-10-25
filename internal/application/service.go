package application

import "github.com/fbriansyah/agn-aggregator-toolbox/internal/port"

type Service struct {
	db port.DatabasePort
}

func NewService(db port.DatabasePort) *Service {
	return &Service{
		db: db,
	}
}
