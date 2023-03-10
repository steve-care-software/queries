package domain

import (
	"github.com/steve-care-software/libs/cryptography/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	hashAdapter := hash.NewAdapter()
	return createElementBuilder(hashAdapter)
}

// Builder represents a query builder
type Builder interface {
	Create() Builder
	WithToken(token hash.Hash) Builder
	WithElement(element Element) Builder
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
