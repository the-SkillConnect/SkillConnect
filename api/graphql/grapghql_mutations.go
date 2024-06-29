package api

import (
	"context"
	"database/sql"

	"github.com/graphql-go/graphql"
	"github.com/the-SkillConnect/SkillConnect/db"
)

// Define the Root Mutation
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"insertUser": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(InsertUserInputType)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				input := p.Args["input"].(map[string]interface{})
				params := db.InsertUserParams{
					Email:       input["email"].(string),
					Password:    input["password"].(string),
					Firstname:   sql.NullString{String: input["firstname"].(string), Valid: input["firstname"] != nil},
					Surname:     sql.NullString{String: input["surname"].(string), Valid: input["surname"] != nil},
					MobilePhone: sql.NullString{String: input["mobilePhone"].(string), Valid: input["mobilePhone"] != nil},
				}
				id, err := DbInstance.InsertUser(context.Background(), params)
				if err != nil {
					return nil, err
				}
				return DbInstance.GetUserByID(context.Background(), id)
			},
		},
		"deleteUser": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(int)
				err := DbInstance.DeleteUserByID(context.Background(), int32(id))
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		"updateUser": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(UpdateUserInputType)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				input := p.Args["input"].(map[string]interface{})
				params := db.UpdateUserByIDParams{
					ID:          int32(input["id"].(int)),
					Email:       input["email"].(string),
					Password:    input["password"].(string),
					Firstname:   sql.NullString{String: input["firstname"].(string), Valid: input["firstname"] != nil},
					Surname:     sql.NullString{String: input["surname"].(string), Valid: input["surname"] != nil},
					MobilePhone: sql.NullString{String: input["mobilePhone"].(string), Valid: input["mobilePhone"] != nil},
				}
				_, err := DbInstance.UpdateUserByID(context.Background(), params)
				if err != nil {
					return nil, err
				}
				return DbInstance.GetUserByID(context.Background(), params.ID)
			},
		},
	},
})
