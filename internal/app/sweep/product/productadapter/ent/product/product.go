// Code generated by entc, DO NOT EDIT.

package product

import (
	"time"

	"github.com/prasetyowira/sweep/internal/app/sweep/product/productadapter/ent/schema"
)

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUID holds the string denoting the uid vertex property in the database.
	FieldUID = "uid"
	// FieldSku holds the string denoting the sku vertex property in the database.
	FieldSku = "sku"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"
	// FieldExpirable holds the string denoting the expirable vertex property in the database.
	FieldExpirable = "expirable"
	// FieldIsDeleted holds the string denoting the is_deleted vertex property in the database.
	FieldIsDeleted = "is_deleted"
	// FieldDeletedAt holds the string denoting the deleted_at vertex property in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at vertex property in the database.
	FieldUpdatedAt = "updated_at"

	// Table holds the table name of the product in the database.
	Table = "products"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldSku,
	FieldName,
	FieldExpirable,
	FieldIsDeleted,
	FieldDeletedAt,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	fields = schema.Product{}.Fields()

	// descUID is the schema descriptor for uid field.
	descUID = fields[0].Descriptor()
	// UIDValidator is a validator for the "uid" field. It is called by the builders before save.
	UIDValidator = func() func(string) error {
		validators := descUID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(uid string) error {
			for _, fn := range fns {
				if err := fn(uid); err != nil {
					return err
				}
			}
			return nil
		}
	}()

	// descSku is the schema descriptor for sku field.
	descSku = fields[1].Descriptor()
	// SkuValidator is a validator for the "sku" field. It is called by the builders before save.
	SkuValidator = func() func(string) error {
		validators := descSku.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(sku string) error {
			for _, fn := range fns {
				if err := fn(sku); err != nil {
					return err
				}
			}
			return nil
		}
	}()

	// descName is the schema descriptor for name field.
	descName = fields[2].Descriptor()
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator = func() func(string) error {
		validators := descName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()

	// descExpirable is the schema descriptor for expirable field.
	descExpirable = fields[3].Descriptor()
	// DefaultExpirable holds the default value on creation for the expirable field.
	DefaultExpirable = descExpirable.Default.(bool)

	// descIsDeleted is the schema descriptor for is_deleted field.
	descIsDeleted = fields[4].Descriptor()
	// DefaultIsDeleted holds the default value on creation for the is_deleted field.
	DefaultIsDeleted = descIsDeleted.Default.(bool)

	// descCreatedAt is the schema descriptor for created_at field.
	descCreatedAt = fields[6].Descriptor()
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt = descCreatedAt.Default.(func() time.Time)

	// descUpdatedAt is the schema descriptor for updated_at field.
	descUpdatedAt = fields[7].Descriptor()
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt = descUpdatedAt.Default.(func() time.Time)
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt = descUpdatedAt.UpdateDefault.(func() time.Time)
)
