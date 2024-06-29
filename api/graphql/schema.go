package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/the-SkillConnect/SkillConnect/db"
)


func NewSchema(dbInstance db.Querier) (graphql.Schema, error) {
	resolver := NewResolver(dbInstance)

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveUser,
				},
				"users": &graphql.Field{
					Type:    graphql.NewList(UserType),
					Resolve: resolver.ResolveGetUsers,
				},
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootMutation",
			Fields: graphql.Fields{
				"insertUser": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(InsertUserInputType)},
					},
					Resolve: resolver.ResolveInsertUser,
				},
				"deleteUser": &graphql.Field{
					Type: graphql.Boolean,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveDeleteUser,
				},
				"updateUser": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(UpdateUserInputType)},
					},
					Resolve: resolver.ResolveUpdateUser,
				},
			},
		}),
	})
}
