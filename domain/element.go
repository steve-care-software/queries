package domain

import "github.com/steve-care-software/libs/cryptography/hash"

type element struct {
	hash    hash.Hash
	element hash.Hash
	index   uint
	query   Query
}

func createElement(
	hash hash.Hash,
	elHash hash.Hash,
	index uint,
) Element {
	return createElementInternally(hash, elHash, index, nil)
}

func createElementWithQuery(
	hash hash.Hash,
	elHash hash.Hash,
	index uint,
	query Query,
) Element {
	return createElementInternally(hash, elHash, index, query)
}

func createElementInternally(
	hash hash.Hash,
	elHash hash.Hash,
	index uint,
	query Query,
) Element {
	out := element{
		hash:    hash,
		element: elHash,
		index:   index,
		query:   query,
	}

	return &out
}

// Hash returns the hash
func (obj *element) Hash() hash.Hash {
	return obj.hash
}

// Element returns the element
func (obj *element) Element() hash.Hash {
	return obj.element
}

// Index returns the index
func (obj *element) Index() uint {
	return obj.index
}

// HasQuery returns true if there is a query, false otherwise
func (obj *element) HasQuery() bool {
	return obj.query != nil
}

// Query returns the query, if any
func (obj *element) Query() Query {
	return obj.query
}
