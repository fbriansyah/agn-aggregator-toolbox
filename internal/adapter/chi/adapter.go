package chi

import (
	"log"
	"net/http"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/port"
)

type M map[string]any

type ChiAdapter struct {
	service  port.ServicePort
	template port.TemplatePort
}

type ChiAdapterConfig struct {
	ServerAddress string
}

func NewChiAdapter(service port.ServicePort, template port.TemplatePort) *ChiAdapter {
	return &ChiAdapter{
		service:  service,
		template: template,
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
