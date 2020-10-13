package stockadapter

import (
	"context"

	"emperror.dev/errors"

	"github.com/prasetyowira/sweep/internal/app/sweep/product"
	"github.com/prasetyowira/sweep/internal/app/sweep/product/productadapter/ent"
	product_ent "github.com/prasetyowira/sweep/internal/app/sweep/product/productadapter/ent/product"
)

type entStore struct {
	client *ent.Client
}

// NewEntStore returns a new sweep store backed by Ent ORM.
func NewEntStore(client *ent.Client) product.Store {
	return entStore{
		client: client,
	}
}

func (s entStore) Store(ctx context.Context, product product.Product) error {
	existing, err := s.client.Product.Query().Where(product_ent.UID(product.ID)).First(ctx)
	if ent.IsNotFound(err) {
		_, err := s.client.Product.Create().
			SetUID(product.ID).
			SetSku(product.SKU).
			SetName(product.Name).
			SetExpirable(product.Expirable).
			Save(ctx)
		if err != nil {
			return err
		}

		return nil
	}
	if err != nil {
		return err
	}

	_, err = s.client.Product.UpdateOneID(existing.ID).
		SetSku(product.SKU).
		SetName(product.Name).
		SetExpirable(product.Expirable).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s entStore) All(ctx context.Context) ([]product.Product, error) {
	messageModels, err := s.client.Product.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	messages := make([]product.Product, 0, len(messageModels))

	for _, messageModel := range messageModels {
		messages = append(messages, product.Product{
			ID:   messageModel.UID,
			SKU:   messageModel.Sku,
			Name:   messageModel.Name,
			Expirable: messageModel.Expirable,
		})
	}

	return messages, nil
}

func (s entStore) Get(ctx context.Context, id string) (product.Product, error) {
	messageModel, err := s.client.Product.Query().Where(product_ent.UID(id)).First(ctx)
	if ent.IsNotFound(err) {
		return product.Product{}, errors.WithStack(product.NotFoundError{ID: id})
	}

	return product.Product{
		ID:   messageModel.UID,
		SKU:   messageModel.Sku,
		Name:   messageModel.Name,
		Expirable: messageModel.Expirable,
	}, nil
}

func (s entStore) GetBySKU(ctx context.Context, sku string) (product.Product, error) {
	messageModel, err := s.client.Product.Query().Where(product_ent.Sku(sku)).First(ctx)
	if ent.IsNotFound(err) {
		return product.Product{}, errors.WithStack(product.NotFoundError{SKU: sku})
	}

	return product.Product{
		ID:   messageModel.UID,
		SKU:   messageModel.Sku,
		Name:   messageModel.Name,
		Expirable: messageModel.Expirable,
	}, nil
}
