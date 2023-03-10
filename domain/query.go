package domain

import (
	"github.com/steve-care-software/libs/cryptography/hash"
)

type query struct {
	hash    hash.Hash
	token   hash.Hash
	element Element
}

func createQuery(
	hash hash.Hash,
	token hash.Hash,
	element Element,
) Query {
	out := query{
		hash:    hash,
		token:   token,
		element: element,
	}

	return &out
}

// Hash returns the hash
func (obj *query) Hash() hash.Hash {
	return obj.hash
}

// Token returns the token
func (obj *query) Token() hash.Hash {
	return obj.token
}

// Element returns the element
func (obj *query) Element() Element {
	return obj.element
}
