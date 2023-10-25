package port

import "github.com/fbriansyah/agn-aggregator-toolbox/internal/adapter/mariadb"

type DatabasePort interface {
	mariadb.DatabaseAdapter
}
