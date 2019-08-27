// Code generated (@generated) by entc, DO NOT EDIT.

package card

import (
	"time"

	"github.com/facebookincubator/ent/entc/integration/ent/schema"
)

const (
	// Label holds the string label denoting the card type in the database.
	Label = "card"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNumber holds the string denoting the number vertex property in the database.
	FieldNumber = "number"
	// FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at vertex property in the database.
	FieldUpdatedAt = "updated_at"

	// Table holds the table name of the card in the database.
	Table = "cards"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "cards"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"

	// OwnerInverseLabel holds the string label denoting the owner inverse edge type in the database.
	OwnerInverseLabel = "user_card"
)

// Columns holds all SQL columns are card fields.
var Columns = []string{
	FieldID,
	FieldNumber,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	fields = schema.Card{}.Fields()
	// descNumber is the schema descriptor for number field.
	descNumber = fields[0].Descriptor()
	// NumberValidator is a validator for the "number" field. It is called by the builders before save.
	NumberValidator = descNumber.Validators[0].(func(string) error)
	// descCreatedAt is the schema descriptor for created_at field.
	descCreatedAt = fields[1].Descriptor()
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt = descCreatedAt.Default.(func() time.Time)
	// descUpdatedAt is the schema descriptor for updated_at field.
	descUpdatedAt = fields[2].Descriptor()
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt = descUpdatedAt.Default.(func() time.Time)
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt = descUpdatedAt.UpdateDefault.(func() time.Time)
)
