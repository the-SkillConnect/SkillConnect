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
		"first_name":     &graphql.Field{Type: graphql.String},
		"surname":        &graphql.Field{Type: graphql.String},
		"mobile_phone":   &graphql.Field{Type: graphql.String},
		"wallet_address": &graphql.Field{Type: graphql.String},
		"created_at":     &graphql.Field{Type: graphql.DateTime},
		"updated_at":     &graphql.Field{Type: graphql.DateTime},
	},
})

var InsertUserIdentityInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InsertUserIdentityInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"email":          &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"password":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"first_name":     &graphql.InputObjectFieldConfig{Type: graphql.String},
		"surname":        &graphql.InputObjectFieldConfig{Type: graphql.String},
		"mobile_phone":   &graphql.InputObjectFieldConfig{Type: graphql.String},
		"wallet_address": &graphql.InputObjectFieldConfig{Type: graphql.String},
	},
})

var UpdateUserIdentityInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateUserIdentityInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id":           &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"email":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"password":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"first_name":   &graphql.InputObjectFieldConfig{Type: graphql.String},
		"surname":      &graphql.InputObjectFieldConfig{Type: graphql.String},
		"mobile_phone": &graphql.InputObjectFieldConfig{Type: graphql.String},
	},
})

var UserProfileType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfile",
	Fields: graphql.Fields{
		"userID":           &graphql.Field{Type: graphql.Int},
		"rating":           &graphql.Field{Type: graphql.Int},
		"description":      &graphql.Field{Type: graphql.String},
		"doneProjects":     &graphql.Field{Type: graphql.Int},
		"givenProjects":    &graphql.Field{Type: graphql.Int},
		"recommendationID": &graphql.Field{Type: graphql.Int},
		"createdAt":        &graphql.Field{Type: graphql.DateTime},
		"updatedAt":        &graphql.Field{Type: graphql.DateTime},
	},
})

var InsertUserProfileInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InsertUserProfileInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"user_id":           &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"rating":            &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"description":       &graphql.InputObjectFieldConfig{Type: graphql.String},
		"done_projects":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"given_projects":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"recommendation_id": &graphql.InputObjectFieldConfig{Type: graphql.Int},
	},
})

var UpdateUserProfileInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateUserProfileInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"user_id":           &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"rating":            &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"description":       &graphql.InputObjectFieldConfig{Type: graphql.String},
		"done_projects":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"given_projects":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"recommendation_id": &graphql.InputObjectFieldConfig{Type: graphql.Int},
	},
})
