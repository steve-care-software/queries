package domain

import (
	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// NewQueryBuilder creates a new query builder
func NewQueryBuilder() QueryBuilder {
	hashAdapter := hash.NewAdapter()
	return createQueryBuilder(hashAdapter)
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	hashAdapter := hash.NewAdapter()
	return createElementBuilder(hashAdapter)
}

// Builder represents queries builder
type Builder interface {
	Create() Builder
	WithGrammar(grammar grammars.Grammar) Builder
	WithQueries(queries []Query) Builder
	Now() (Queries, error)
}

// Queries represents queries
type Queries interface {
	Hash() hash.Hash
	Grammar() grammars.Grammar
	Queries() []Query
}

// QueryBuilder represents a query builder
type QueryBuilder interface {
	Create() QueryBuilder
	WithToken(token hash.Hash) QueryBuilder
	WithElement(element Element) QueryBuilder
	Now() (Query, error)
}

// Query represents a query
type Query interface {
	Hash() hash.Hash
	Token() hash.Hash
	Element() Element
}

// ElementBuilder represents an eleemnt builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithIndex(index uint) ElementBuilder
	WithQuery(query Query) ElementBuilder
	WithElement(element hash.Hash) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Hash() hash.Hash
	Element() hash.Hash
	Index() uint
	HasQuery() bool
	Query() Query
}
