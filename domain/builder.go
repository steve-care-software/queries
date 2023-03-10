package domain

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	token       hash.Hash
	element     Element
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		token:       nil,
		element:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithToken adds a token to the builder
func (app *builder) WithToken(token hash.Hash) Builder {
	app.token = token
	return app
}

// WithElement adds an element to the builder
func (app *builder) WithElement(element Element) Builder {
	app.element = element
	return app
}

// Now builds a new Token instance
func (app *builder) Now() (Query, error) {
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
