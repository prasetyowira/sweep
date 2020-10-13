package product

import (
	"context"

	"emperror.dev/errors"
)

// Message describing a sweep and its content
type Product struct {
	ID			string
	SKU			string
	Name		string
	Expirable	bool
}

// +kit:endpoint:errorStrategy=service

// Service manages a list of messages.
type Service interface {
	// CreateProduct adds a new product to the products list.
	CreateProduct(ctx context.Context, sku string, name string, expirable bool) (id string, err error)

	// ListProducts returns the list of products.
	ListProducts(ctx context.Context) (products []Product, err error)

	// GetProduct returns the product by id.
	GetProduct(ctx context.Context, id string) (product Product, err error)

	// GetProductBySKU returns the product by sku.
	GetProductBySKU(ctx context.Context, sku string) (product Product, err error)
}

type service struct {
	idgenerator IDGenerator
	store       Store
}

// IDGenerator generates a new ID.
type IDGenerator interface {
	// Generate generates a new ID.
	Generate() (string, error)
}

// Store provides sweep persistence.
type Store interface {
	// Store stores a product.
	Store(ctx context.Context, product Product) error

	// All returns all prducts.
	All(ctx context.Context) ([]Product, error)

	// Get returns a single product by its ID.
	Get(ctx context.Context, id string) (Product, error)

	// Get returns a single product by its SKU.
	GetBySKU(ctx context.Context, sky string) (Product, error)
}

// NotFoundError is returned if a sweep cannot be found.
type NotFoundError struct {
	ID string
	SKU string
}

// Error implements the error interface.
func (NotFoundError) Error() string {
	return "product not found"
}

// Details returns error details.
func (e NotFoundError) Details() []interface{} {
	return []interface{}{"product_id", e.ID}
}

// NotFound tells a client that this error is related to a resource being not found.
// Can be used to translate the error to eg. status code.
func (NotFoundError) NotFound() bool {
	return true
}

// ServiceError tells the transport layer whether this error should be translated into the transport format
// or an internal error should be returned instead.
func (NotFoundError) ServiceError() bool {
	return true
}


// NewService returns a new Service.
func NewService(idgenerator IDGenerator, store Store) Service {
	return &service{
		idgenerator: idgenerator,
		store:       store,
	}
}

type validationError struct {
	violations map[string][]string
}

func (validationError) Error() string {
	return "invalid product"
}

func (e validationError) Violations() map[string][]string {
	return e.violations
}

// Validation tells a client that this error is related to a resource being invalid.
// Can be used to translate the error to eg. status code.
func (validationError) Validation() bool {
	return true
}

// ServiceError tells the transport layer whether this error should be translated into the transport format
// or an internal error should be returned instead.
func (validationError) ServiceError() bool {
	return true
}

func (s service) CreateProduct(ctx context.Context, sku string, name string, expirable bool) (string, error) {
	id, err := s.idgenerator.Generate()
	if err != nil {
		return "", err
	}

	if sku == "" {
		return "", errors.WithStack(validationError{violations: map[string][]string{
			"sku": {
				"sku cannot be empty",
			},
		}})
	}

	if name == "" {
		return "", errors.WithStack(validationError{violations: map[string][]string{
			"name": {
				"name cannot be empty",
			},
		}})
	}

	product := Product{
		ID:   id,
		SKU: sku,
		Name: name,
		Expirable: expirable,
	}

	err = s.store.Store(ctx, product)
	if err != nil {
		return "", err
	}

	return id, err
}

func (s service) ListProducts(ctx context.Context) ([]Product, error) {
	return s.store.All(ctx)
}

func (s service) GetProduct(ctx context.Context, id string) (Product, error) {
	product, err := s.store.Get(ctx, id)
	if err != nil {
		return Product{}, errors.WithMessage(err, "product not found")
	}

	return product, nil
}

func (s service) GetProductBySKU(ctx context.Context, sku string) (Product, error) {
	product, err := s.store.Get(ctx, sku)
	if err != nil {
		return Product{}, errors.WithMessage(err, "product not found")
	}

	return product, nil
}
