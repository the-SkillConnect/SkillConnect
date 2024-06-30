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

// Project GraphQL type
var ProjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Project",
	Fields: graphql.Fields{
		"id":           &graphql.Field{Type: graphql.Int},
		"title":        &graphql.Field{Type: graphql.String},
		"description":  &graphql.Field{Type: graphql.String},
		"total_amount": &graphql.Field{Type: graphql.String},
		"orderDate":    &graphql.Field{Type: graphql.String},
		"status":       &graphql.Field{Type: graphql.Boolean},
		"user_id":      &graphql.Field{Type: graphql.Int},
		"fee":          &graphql.Field{Type: graphql.String},
	},
})

var InsertProjectInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InsertProjectInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"title":        &graphql.InputObjectFieldConfig{Type: graphql.String},
		"description":  &graphql.InputObjectFieldConfig{Type: graphql.String},
		"total_amount": &graphql.InputObjectFieldConfig{Type: graphql.String},
		"status":       &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"user_id":      &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"fee":          &graphql.InputObjectFieldConfig{Type: graphql.String},
	},
})
var UpdateProjectInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateProjectInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id":           &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"title":        &graphql.InputObjectFieldConfig{Type: graphql.String},
		"description":  &graphql.InputObjectFieldConfig{Type: graphql.String},
		"total_amount": &graphql.InputObjectFieldConfig{Type: graphql.String},
		"status":       &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"user_id":      &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"fee":          &graphql.InputObjectFieldConfig{Type: graphql.String},
	},
})

var AssignedProjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AssignedProject",
	Fields: graphql.Fields{
		"user_id":    &graphql.Field{Type: graphql.Int},
		"project_id": &graphql.Field{Type: graphql.Int},
		"issued":     &graphql.Field{Type: graphql.Boolean},
	},
})

var InsertAssignedProjectInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InsertAssignedProjectInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"user_id":    &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"project_id": &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"issued":     &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
	},
})

var UpdateAssignedProjectInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateAssignedProjectInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"user_id":    &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"project_id": &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"issued":     &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
	},
})
