package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			NotEmpty().
			Unique(),
		field.String("name").
			NotEmpty(),
		field.String("password").
			Sensitive().
			NotEmpty(),
		field.String("mail_address").
			Unique().
			NotEmpty(),
		field.String("introduction"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
