package api

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/the-SkillConnect/SkillConnect/db"
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
	},
})

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
					RoleID:      sql.NullInt32{Int32: int32(input["roleID"].(int)), Valid: input["roleID"] != nil},
				}
				fmt.Printf("params: %+v\n", params)
				id, err := DbInstance.InsertUser(context.Background(), params)
				if err != nil {
					return nil, err
				}
				return DbInstance.GetUserByID(context.Background(), id)
			},
		},
	},
})
