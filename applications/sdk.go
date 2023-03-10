package applications

import (
	"github.com/steve-care-software/grammars/domain/trees"
	queries "github.com/steve-care-software/queries/domain"
)

// Application represents the query application
type Application interface {
	Execute(ast trees.Tree, query queries.Query) ([]byte, error)
}
