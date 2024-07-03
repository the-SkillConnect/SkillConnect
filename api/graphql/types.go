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
		"userID":       &graphql.Field{Type: graphql.Int},
		"rating":       &graphql.Field{Type: graphql.Int},
		"description":  &graphql.Field{Type: graphql.String},
		"doneProject":  &graphql.Field{Type: graphql.Int},
		"givenProject": &graphql.Field{Type: graphql.Int},
		"createdAt":    &graphql.Field{Type: graphql.DateTime},
		"updatedAt":    &graphql.Field{Type: graphql.DateTime},
	},
})

var InsertUserProfileInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InsertUserProfileInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"user_id":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"rating":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"description":   &graphql.InputObjectFieldConfig{Type: graphql.String},
		"done_project":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"given_project": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
})

var UpdateUserProfileInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateUserProfileInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"user_id":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"rating":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"description":   &graphql.InputObjectFieldConfig{Type: graphql.String},
		"done_project":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"given_project": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
})

var UserRecommendationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserRecommendation",
	Fields: graphql.Fields{
		"given_id":    &graphql.Field{Type: graphql.Int},
		"received_id": &graphql.Field{Type: graphql.Int},
		"description": &graphql.Field{Type: graphql.String},
	},
})

var InsertUserRecommendationInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InsertUserRecommendationInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"given_id":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"received_id": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"description": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
	},
})

var ProjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Project",
	Fields: graphql.Fields{
		"id":           &graphql.Field{Type: graphql.Int},
		"description":  &graphql.Field{Type: graphql.String},
		"title":        &graphql.Field{Type: graphql.String},
		"total_amount": &graphql.Field{Type: graphql.String},
		"done_status":  &graphql.Field{Type: graphql.Boolean},
		"user_id":      &graphql.Field{Type: graphql.Int},
		"fee":          &graphql.Field{Type: graphql.String},
		"category_id":  &graphql.Field{Type: graphql.ID},
		"created_at":   &graphql.Field{Type: graphql.DateTime},
		"updated_at":   &graphql.Field{Type: graphql.DateTime},
	},
})

var InsertProjectInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InsertProjectInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"description":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"title":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"total_amount": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"done_status":  &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"user_id":      &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"fee":          &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"category_id":  &graphql.InputObjectFieldConfig{Type: graphql.ID},
	},
})

var UpdateProjectInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateProjectInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id":           &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"description":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"title":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"total_amount": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"done_status":  &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"user_id":      &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"fee":          &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"category_id":  &graphql.InputObjectFieldConfig{Type: graphql.ID},
	},
})

var CommentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.Int},
		"user_id":    &graphql.Field{Type: graphql.Int},
		"project_id": &graphql.Field{Type: graphql.Int},
		"date":       &graphql.Field{Type: graphql.DateTime},
		"text":       &graphql.Field{Type: graphql.String},
	},
})

var InsertCommentInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InsertCommentInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"user_id":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"project_id": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"text":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
	},
})

var UpdateCommentInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateCommentInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id":         &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"user_id":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"project_id": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"text":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
	},
})

var AssignProjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AssignProject",
	Fields: graphql.Fields{
		"user_id":    &graphql.Field{Type: graphql.Int},
		"project_id": &graphql.Field{Type: graphql.Int},
		"created_at": &graphql.Field{Type: graphql.DateTime},
		"updated_at": &graphql.Field{Type: graphql.DateTime},
	},
})

var InsertAssignProjectInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InsertAssignProjectInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"user_id":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"project_id": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
})

var DeleteAssignProjectInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "DeleteAssignProjectInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"user_id":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"project_id": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
})

var CategoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Category",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.Int},
		"title": &graphql.Field{Type: graphql.String},
	},
})

var InsertCategoryInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InsertCategoryInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"title": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
	},
})
