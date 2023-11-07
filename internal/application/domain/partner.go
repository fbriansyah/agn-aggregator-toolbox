package domain

import "github.com/fbriansyah/agn-aggregator-toolbox/internal/adapter/mariadb"

type PartnerDomain struct {
	IdPartner   int32
	NamaPartner string
}

func (p *PartnerDomain) FromMPartner(partner mariadb.MPartnerid) {
	p.IdPartner = partner.Idpartner
	p.NamaPartner = partner.Nama.String
}
