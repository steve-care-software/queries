package references

import (
	"github.com/steve-care-software/libs/cryptography/hash"
	queries "github.com/steve-care-software/queries/domain"
)

// Builder represents a reference builder
type Builder interface {
	Create() Builder
	WithReference(reference queries.Queries) Builder
	WithQueries(queries []Query) Builder
	Now() (Reference, error)
}

// Reference represents a reference
type Reference interface {
	Reference() queries.Queries
	Queries() []Query
}

// QueryBuilder represents a query builder
type QueryBuilder interface {
	Create() QueryBuilder
	WithName(name string) QueryBuilder
	WithReference(reference queries.Query) QueryBuilder
	Now() (Query, error)
}

// Query represents a query reference
type Query interface {
	Name() string
	Reference() queries.Query
}

// TokensBuilder represents tokens builder
type TokensBuilder interface {
	Create() TokensBuilder
	WithList(list []Token) TokensBuilder
	Now() (Tokens, error)
}

// Tokens represents token references
type Tokens interface {
	List() []Token
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithName(name string) TokenBuilder
	WithReference(reference hash.Hash) TokenBuilder
	Now() (Token, error)
}

// Token represents a represents a token reference
type Token interface {
	Name() string
	Reference() hash.Hash
}
