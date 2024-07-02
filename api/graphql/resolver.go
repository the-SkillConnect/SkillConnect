package graphql

import (
	"context"
	"database/sql"
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
	insertParams := db.InsertUserIdentityParams{
		Email:         input["email"].(string),
		Password:      input["password"].(string),
		FirstName:     input["first_name"].(string),
		Surname:       input["surname"].(string),
		MobilePhone:   input["mobile_phone"].(string),
		WalletAddress: input["wallet_address"].(string),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	fmt.Println("herree", insertParams)
	id, err := r.DbInstance.InsertUserIdentity(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}
	fmt.Println("here")
	return r.DbInstance.GetUserIdentityByID(context.Background(), id)
}

func (r *Resolver) ResolveGetUserByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	return r.DbInstance.GetUserIdentityByID(context.Background(), int64(id))
}

func (r *Resolver) ResolveDeleteUser(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	err := r.DbInstance.DeleteUserIdentityByID(context.Background(), int64(id))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Resolver) ResolveUpdateUser(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	updateParams := db.UpdateUserIdentityByIDParams{
		ID:          int64(input["id"].(int)),
		Email:       input["email"].(string),
		Password:    input["password"].(string),
		FirstName:   input["first_name"].(string),
		Surname:     input["surname"].(string),
		MobilePhone: input["mobile_phone"].(string),
		UpdatedAt:   time.Now(),
	}
	_, err := r.DbInstance.UpdateUserIdentityByID(context.Background(), updateParams)
	if err != nil {
		return nil, err
	}
	return r.DbInstance.GetUserIdentityByID(context.Background(), updateParams.ID)
}

func (r *Resolver) ResolveGetUsers(params graphql.ResolveParams) (interface{}, error) {
	return r.DbInstance.GetUsersIdentity(context.Background())
}

func (r *Resolver) ResolveInsertUserProfile(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	insertParams := db.InsertUserProfileParams{
		UserID:           int64(input["user_id"].(int)),
		Rating:           int64(input["rating"].(int)),
		Description:      sql.NullString{String: input["description"].(string), Valid: true},
		DoneProject:      int64(input["done_project"].(int)),
		GivenProject:     int64(input["given_project"].(int)),
		RecommendationID: sql.NullInt64{Int64: input["recommendation_id"].(int64), Valid: true},
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	id, err := r.DbInstance.InsertUserProfile(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}

	return r.DbInstance.GetUserProfileByUserID(context.Background(), id)
}

func (r *Resolver) ResolveGetUserProfile(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int64)
	return r.DbInstance.GetUserProfileByUserID(context.Background(), id)
}

func (r *Resolver) ResolveUpdateUserProfile(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	updateParams := db.UpdateUserProfileParams{
		UserID:           int64(input["user_id"].(int)),
		Rating:           input["rating"].(int64),
		Description:      sql.NullString{String: input["description"].(string), Valid: true},
		DoneProject:      input["done_project"].(int64),
		GivenProject:     input["given_project"].(int64),
		RecommendationID: sql.NullInt64{Int64: input["recommendation_id"].(int64), Valid: true},
		UpdatedAt:        time.Now(),
	}

	_, err := r.DbInstance.UpdateUserProfile(context.Background(), updateParams)
	if err != nil {
		return nil, err
	}

	return r.DbInstance.GetUserProfileByUserID(context.Background(), updateParams.UserID)
}

func (r *Resolver) ResolveDeleteUserProfile(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int64)
	err := r.DbInstance.DeleteUserProfileByID(context.Background(), id)
	if err != nil {
		return false, err
	}
	return true, nil
}
