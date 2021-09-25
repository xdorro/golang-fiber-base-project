package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/xdorro/golang-fiber-base-project/pkg/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),

		field.String("email").
			NotEmpty().
			Annotations(
				entgql.OrderField("EMAIL"),
			),

		field.String("password").
			NotEmpty(),

		field.Int("status").
			Default(1).
			Annotations(
				entgql.OrderField("STATUS"),
			),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.BaseMixin{},
		mixin.TimeMixin{},
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("email", "status"),
	}
}