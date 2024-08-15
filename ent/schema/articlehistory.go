package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ArticleHistory holds the schema definition for the ArticleHistory entity.
type ArticleHistory struct {
	ent.Schema
}

// Fields of the ArticleHistory.
func (ArticleHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at"),
	}
}

// Edges of the ArticleHistory.
func (ArticleHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("article", Article.Type),
	}
}
