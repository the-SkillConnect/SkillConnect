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

var DbInstance db.Querier

// Setter for DbInstance
func SetDbInstance(instance db.Querier) {
	DbInstance = instance
}

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
	},
})


// Input type for updating a user
var UpdateUserInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateUserInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id":          &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"email":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"password":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"firstname":   &graphql.InputObjectFieldConfig{Type: graphql.String},
		"surname":     &graphql.InputObjectFieldConfig{Type: graphql.String},
		"mobilePhone": &graphql.InputObjectFieldConfig{Type: graphql.String},
	},
})