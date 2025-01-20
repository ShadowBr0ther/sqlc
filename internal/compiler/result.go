package compiler

import (
	"github.com/ShadowBr0ther/sqlc/internal/sql/catalog"
)

type Result struct {
	Catalog *catalog.Catalog
	Queries []*Query
}
