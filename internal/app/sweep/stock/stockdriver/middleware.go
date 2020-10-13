package stockdriver

import (
	"context"

	"github.com/prasetyowira/sweep/internal/app/sweep/product"
)

// Middleware describes a service middleware.
type Middleware func(product.Service) product.Service

// LoggingMiddleware is a service level logging middleware for TodoList.
func LoggingMiddleware(logger product.Logger) Middleware {
	return func(next product.Service) product.Service {
		return loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   product.Service
	logger product.Logger
}

func (mw loggingMiddleware) CreateProduct(ctx context.Context, sku string, name string, expirable bool) (string, error) {
	logger := mw.logger.WithContext(ctx)

	logger.Info("creating product")

	id, err := mw.next.CreateProduct(ctx, sku, name, expirable)
	if err != nil {
		return id, err
	}

	logger.Info("created product", map[string]interface{}{
		"id": id,
	})

	return id, err
}

func (mw loggingMiddleware) ListProducts(ctx context.Context) ([]product.Product, error) {
	logger := mw.logger.WithContext(ctx)

	logger.Info("listing products")

	return mw.next.ListProducts(ctx)
}

func (mw loggingMiddleware) GetProduct(ctx context.Context, id string) (product.Product, error) {
	logger := mw.logger.WithContext(ctx)

	logger.Info("get one product", map[string]interface{}{
		"id": id,
	})

	return mw.next.GetProduct(ctx, id)
}

func (mw loggingMiddleware) GetProductBySKU(ctx context.Context, sku string) (product.Product, error) {
	logger := mw.logger.WithContext(ctx)

	logger.Info("get one product", map[string]interface{}{
		"sku": sku,
	})

	return mw.next.GetProductBySKU(ctx, sku)
}

