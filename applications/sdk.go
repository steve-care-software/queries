package applications

import (
	"github.com/steve-care-software/grammars/domain/trees"
	queries "github.com/steve-care-software/queries/domain"
	"github.com/steve-care-software/queries/domain/references"
)

// Application represents the query application
type Application interface {
	Compile(reference references.Reference) (queries.Queries, error)
	Execute(ast trees.Tree, queries queries.Queries) (map[string][]byte, error)
	Reference(tokenHashData map[string][]byte, reference references.Reference) (map[string][]byte, error)
}
