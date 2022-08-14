package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Record holds the schema definition for the Record entity.
type Record struct {
	ent.Schema
}

// Fields of the Record.
func (Record) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user"),
		field.Int64("day"),
		field.Int64("type"),
		field.Time("created_at"),
		field.Time("updated_at").Default(time.Now),
	}
}

func (Record) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "reoord"},
	}
}

func (Record) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("user", "day", "type").
			Unique(),
	}
}

// Edges of the Record.
func (Record) Edges() []ent.Edge {
	return nil
}
