package graphql

import (
	"github.com/graphql-go/graphql"
)

type Input struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

var UserIdentityType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserIdentity",
	Fields: graphql.Fields{
		"id":             &graphql.Field{Type: graphql.Int},
		"email":          &graphql.Field{Type: graphql.String},
		"password":       &graphql.Field{Type: graphql.String},
		"firstname":      &graphql.Field{Type: graphql.String},
		"surname":        &graphql.Field{Type: graphql.String},
		"mobile_phone":   &graphql.Field{Type: graphql.String},
		"wallet_address": &graphql.Field{Type: graphql.String},
		"created_at":     &graphql.Field{Type: graphql.DateTime},
		"updated_at":     &graphql.Field{Type: graphql.DateTime},
	},
})

var InsertUserInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InsertUserIdentityInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"email":          &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"password":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"firstname":      &graphql.InputObjectFieldConfig{Type: graphql.String},
		"surname":        &graphql.InputObjectFieldConfig{Type: graphql.String},
		"mobile_phone":   &graphql.InputObjectFieldConfig{Type: graphql.String},
		"wallet_address": &graphql.InputObjectFieldConfig{Type: graphql.String},
	},
})

var UpdateUserInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateUserIdentityInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id":           &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"email":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"password":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"firstname":    &graphql.InputObjectFieldConfig{Type: graphql.String},
		"surname":      &graphql.InputObjectFieldConfig{Type: graphql.String},
		"mobile_phone": &graphql.InputObjectFieldConfig{Type: graphql.String},
	},
})
