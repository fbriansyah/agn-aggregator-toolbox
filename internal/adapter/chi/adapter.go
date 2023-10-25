package chi

import (
	"log"
	"net/http"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/port"
)

type ChiAdapter struct {
	service port.ServicePort
}

type ChiAdapterConfig struct {
	ServerAddress string
}

func NewChiAdapter(service port.ServicePort) *ChiAdapter {

	return &ChiAdapter{
		service: service,
	}
}

func (adapter *ChiAdapter) Run(config ChiAdapterConfig) {
	srv := &http.Server{
		Addr:    config.ServerAddress,
		Handler: adapter.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
