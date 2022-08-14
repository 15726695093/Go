package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Worktime holds the schema definition for the Worktime entity.
type Worktime struct {
	ent.Schema
}

// Fields of the Worktime.
func (Worktime) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user"),
		field.Int64("day"),
		field.Int64("minute"),
		field.Time("created_at"),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the Worktime.
func (Worktime) Edges() []ent.Edge {
	return nil
}

func (Worktime) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "worktime"},
	}
}

func (Worktime) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("user", "day").
			Unique(),
	}
}
