package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Todo holds the schema definition for the Product entity.
type Stock struct {
	ent.Schema
}

// Fields of the Product.
func (Stock) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").
			MaxLen(26).
			NotEmpty().
			Unique().
			Immutable(),
		field.Int("product_id"),
		field.Time("expire_date").Optional().Nillable(),
		field.Time("inbound_date"),
		field.Time("inbound_qty"),
		field.Time("current_stock"),

		field.String("sku").MaxLen(190).NotEmpty().Unique(),
		field.String("name").MaxLen(255).NotEmpty(),
		field.Bool("expirable").Default(false),
		field.Bool("is_deleted").Default(false),
		field.Time("deleted_at").Optional().Nillable(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Product.
func (Stock) Edges() []ent.Edge {
	return nil
}
