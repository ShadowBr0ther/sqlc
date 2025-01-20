// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/ShadowBr0ther/sqlc/internal/sql/ast"
	"github.com/ShadowBr0ther/sqlc/internal/sql/catalog"
)

var funcsLo = []*catalog.Function{
	{
		Name:       "lo_manage",
		Args:       []*catalog.Argument{},
		ReturnType: &ast.TypeName{Name: "trigger"},
	},
	{
		Name: "lo_oid",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "lo"},
			},
		},
		ReturnType: &ast.TypeName{Name: "oid"},
	},
}

func Lo() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsLo
	return s
}
