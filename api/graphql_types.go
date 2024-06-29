package api

import (
	"github.com/graphql-go/graphql"
	"github.com/the-SkillConnect/SkillConnect/db"
)

// Input struct for parsing incoming GraphQL requests
type Input struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

// Define your database instance here (replace with actual implementation)
var DbInstance db.Querier

// User GraphQL type
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.Int},
		"email":       &graphql.Field{Type: graphql.String},
		"password":    &graphql.Field{Type: graphql.String},
		"firstname":   &graphql.Field{Type: graphql.String},
		"surname":     &graphql.Field{Type: graphql.String},
		"mobilePhone": &graphql.Field{Type: graphql.String},
		"roleID":      &graphql.Field{Type: graphql.Int},
	},
})

// Input type for inserting a user
var InsertUserInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InsertUserInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"email":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"password":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"firstname":   &graphql.InputObjectFieldConfig{Type: graphql.String},
		"surname":     &graphql.InputObjectFieldConfig{Type: graphql.String},
		"mobilePhone": &graphql.InputObjectFieldConfig{Type: graphql.String},
		"roleID":      &graphql.InputObjectFieldConfig{Type: graphql.Int},
	},
})
