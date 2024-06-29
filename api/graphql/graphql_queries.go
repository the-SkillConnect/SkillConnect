package api

import (
	"context"

	"github.com/graphql-go/graphql"
)

// Define the Root Query
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(int)
				return DbInstance.GetUserByID(context.Background(), int32(id))
			},
		},
		"users": &graphql.Field{
			Type: graphql.NewList(UserType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return DbInstance.GetUsers(context.Background())
			},
		},
	},
})
