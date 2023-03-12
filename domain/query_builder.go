package domain

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type queryBuilder struct {
	hashAdapter hash.Adapter
	token       hash.Hash
	element     Element
}

func createQueryBuilder(
	hashAdapter hash.Adapter,
) QueryBuilder {
	out := queryBuilder{
		hashAdapter: hashAdapter,
		token:       nil,
		element:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *queryBuilder) Create() QueryBuilder {
	return createQueryBuilder(app.hashAdapter)
}

// WithToken adds a token to the builder
func (app *queryBuilder) WithToken(token hash.Hash) QueryBuilder {
	app.token = token
	return app
}

// WithElement adds an element to the builder
func (app *queryBuilder) WithElement(element Element) QueryBuilder {
	app.element = element
	return app
}

// Now builds a new Token instance
func (app *queryBuilder) Now() (Query, error) {
	if app.token == nil {
		return nil, errors.New("the token hash is mandatory in order to build a Token instance")
	}

	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Token instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.token.Bytes(),
		app.element.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createQuery(*pHash, app.token, app.element), nil
}
