package graphql

import (
	"context"
	"fmt"
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

func (r *Resolver) ResolveInsertUser(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	fmt.Println("heree", input)
	insertParams := db.InsertUserParams{
		Email:         input["email"].(string),
		Password:      input["password"].(string),
		Firstname:     input["firstname"].(string),
		Surname:       input["surname"].(string),
		MobilePhone:   input["mobile_phone"].(string),
		WalletAddress: input["wallet_address"].(string),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	fmt.Println("herree", insertParams)
	id, err := r.DbInstance.InsertUser(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}
	fmt.Println("here")
	return r.DbInstance.GetUserByID(context.Background(), id)
}

func (r *Resolver) ResolveGetUserByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	return r.DbInstance.GetUserByID(context.Background(), int64(id))
}

func (r *Resolver) ResolveDeleteUser(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	err := r.DbInstance.DeleteUserByID(context.Background(), int64(id))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Resolver) ResolveUpdateUser(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	updateParams := db.UpdateUserByIDParams{
		ID:          int64(input["id"].(int)),
		Email:       input["email"].(string),
		Password:    input["password"].(string),
		Firstname:   input["firstname"].(string),
		Surname:     input["surname"].(string),
		MobilePhone: input["mobile_phone"].(string),
		UpdatedAt:   time.Now(),
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
