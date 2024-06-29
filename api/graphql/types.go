package graphql

import (
	"github.com/graphql-go/graphql"
)

type Input struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

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

var ProjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Project",
	Fields: graphql.Fields{
		"id":           &graphql.Field{Type: graphql.Int},
		"title":        &graphql.Field{Type: graphql.String},
		"description":  &graphql.Field{Type: graphql.String},
		"total_amount": &graphql.Field{Type: graphql.Int},
		"order_date":   &graphql.Field{Type: graphql.DateTime},
		"status":       &graphql.Field{Type: graphql.Boolean},
		"user_id":      &graphql.Field{Type: graphql.Int},
		"fee":          &graphql.Field{Type: graphql.Int},
	},
})

var InsertProjectInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InsertProjectInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"title":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"description":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"total_amount": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"order_date":   &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.DateTime)},
		"status":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
		"user_id":      &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"fee":          &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
})