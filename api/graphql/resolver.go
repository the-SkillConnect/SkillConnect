package graphql

import (
	"context"
	"database/sql"

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

func (r *Resolver) ResolveUser(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	return r.DbInstance.GetUserByID(context.Background(), int32(id))
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
