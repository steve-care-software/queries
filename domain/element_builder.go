package domain

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type elementBuilder struct {
	hashAdapter hash.Adapter
	element     hash.Hash
	pIndex      *uint
	query       Query
}

func createElementBuilder(
	hashAdapter hash.Adapter,
) ElementBuilder {
	out := elementBuilder{
		hashAdapter: hashAdapter,
		element:     nil,
		pIndex:      nil,
		query:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder(app.hashAdapter)
}

// WithIndex adds an index to the builder
func (app *elementBuilder) WithIndex(index uint) ElementBuilder {
	app.pIndex = &index
	return app
}

// WithQuery adds a query to the builder
func (app *elementBuilder) WithQuery(query Query) ElementBuilder {
	app.query = query
	return app
}

// WithElement adds an element to the builder
func (app *elementBuilder) WithElement(element hash.Hash) ElementBuilder {
	app.element = element
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build an Element instance")
	}

	if app.query == nil {
		return nil, errors.New("the query is mandatory in order to build an Element instance")
	}

	if app.element == nil {
		return nil, errors.New("the element's hash is mandatory in order to build an Element instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(fmt.Sprintf("%d", *app.pIndex)),
		app.query.Hash().Bytes(),
		app.element.Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createElement(*pHash, app.element, *app.pIndex), nil
}
