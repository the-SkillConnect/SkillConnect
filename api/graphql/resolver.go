package graphql

import (
	"context"
	"database/sql"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/the-SkillConnect/SkillConnect/db"
)

type Resolver struct {
	DbInstance db.Querier
}

func NewResolver(dbInstance db.Querier) *Resolver {
	return &Resolver{
		DbInstance: dbInstance,
	}
}

func (r *Resolver) ResolveGetUserByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	return r.DbInstance.GetUserByID(context.Background(), int32(id))
}

func (r *Resolver) ResolveGetProjectByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	return r.DbInstance.GetProjectByID(context.Background(), int32(id))
}

func (r *Resolver) ResolveInsertUser(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	insertParams := db.InsertUserParams{
		Email:       input["email"].(string),
		Password:    input["password"].(string),
		Firstname:   sql.NullString{String: input["firstname"].(string), Valid: input["firstname"] != nil},
		Surname:     sql.NullString{String: input["surname"].(string), Valid: input["surname"] != nil},
		MobilePhone: sql.NullString{String: input["mobilePhone"].(string), Valid: input["mobilePhone"] != nil},
	}
	id, err := r.DbInstance.InsertUser(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}
	return r.DbInstance.GetUserByID(context.Background(), id)
}

func (r *Resolver) ResolveInsertProject(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	insertParams := db.InsertProjectParams{
		Title:       sql.NullString{String: input["title"].(string), Valid: input["title"] != nil},
		Description: sql.NullString{String: input["description"].(string), Valid: input["description"] != nil},
		TotalAmount: sql.NullInt32{Int32: input["total_amount"].(int32), Valid: input["total_amount"] != nil},
		OrderDate:   sql.NullTime{Time: input["order_date"].(time.Time), Valid: input["order_date"] != nil},
		Status:      sql.NullBool{Bool: input["status"].(bool), Valid: input["status"] != nil},
		UserID:      sql.NullInt32{Int32: input["user_id"].(int32), Valid: input["user_id"] != nil},
		Fee:         sql.NullInt32{Int32: input["fee"].(int32), Valid: input["fee"] != nil},
	}
	id, err := r.DbInstance.InsertProject(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}
	return r.DbInstance.GetProjectByID(context.Background(), id)
}

func (r *Resolver) ResolveDeleteUser(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	err := r.DbInstance.DeleteUserByID(context.Background(), int32(id))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Resolver) ResolveUpdateUser(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	updateParams := db.UpdateUserByIDParams{
		ID:          int32(input["id"].(int)),
		Email:       input["email"].(string),
		Password:    input["password"].(string),
		Firstname:   sql.NullString{String: input["firstname"].(string), Valid: input["firstname"] != nil},
		Surname:     sql.NullString{String: input["surname"].(string), Valid: input["surname"] != nil},
		MobilePhone: sql.NullString{String: input["mobilePhone"].(string), Valid: input["mobilePhone"] != nil},
	}
	_, err := r.DbInstance.UpdateUserByID(context.Background(), updateParams)
	if err != nil {
		return nil, err
	}
	return r.DbInstance.GetUserByID(context.Background(), updateParams.ID)
}

func (r *Resolver) ResolveGetUsers(params graphql.ResolveParams) (interface{}, error) {
	return r.DbInstance.GetUsers(context.Background())
}

func (r *Resolver) ResolveGetProjects(params graphql.ResolveParams) (interface{}, error) {
	return r.DbInstance.GetProjects(params.Context)
}
