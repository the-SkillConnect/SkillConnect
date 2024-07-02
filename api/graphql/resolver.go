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

func (r *Resolver) ResolveGetUserRecommendationByGivenID(params graphql.ResolveParams) (interface{}, error) {
	givenID := params.Args["given_id"].(int)
	return r.DbInstance.GetUserRecommendationByGivenID(context.Background(), int64(givenID))
}

func (r *Resolver) ResolveGetUserRecommendationByReceivedID(params graphql.ResolveParams) (interface{}, error) {
	receivedID := params.Args["received_id"].(int)
	return r.DbInstance.GetUserRecommendationByReceivedID(context.Background(), int64(receivedID))
}

func (r *Resolver) ResolveInsertUserRecommendation(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	insertParams := db.InsertUserRecommendationParams{
		GivenID:     int64(input["given_id"].(int)),
		ReceivedID:  int64(input["received_id"].(int)),
		Description: input["description"].(string),
	}

	id, err := r.DbInstance.InsertUserRecommendation(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}

	return r.DbInstance.GetUserRecommendationByGivenID(context.Background(), id.GivenID)
}

func (r *Resolver) ResolveInsertProject(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	insertParams := db.InsertProjectParams{
		Description: input["description"].(string),
		Title:       input["title"].(string),
		TotalAmount: input["total_amount"].(string),
		DoneStatus:  sql.NullBool{Bool: input["done_status"].(bool), Valid: true},
		UserID:      int64(input["user_id"].(int)),
		Fee:         input["fee"].(string),
		Categories:  sql.NullInt64{Int64: input["categories"].(int64), Valid: true},
	}

	id, err := r.DbInstance.InsertProject(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}

	return r.DbInstance.GetProjectByID(context.Background(), id)
}

func (r *Resolver) ResolveGetProjectByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	return r.DbInstance.GetProjectByID(context.Background(), int64(id))
}

func (r *Resolver) ResolveUpdateProject(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	updateParams := db.UpdateProjectByIDParams{
		ID:          int64(input["id"].(int)),
		Description: input["description"].(string),
		Title:       input["title"].(string),
		TotalAmount: input["total_amount"].(string),
		DoneStatus:  sql.NullBool{Bool: input["done_status"].(bool), Valid: true},
		UserID:      int64(input["user_id"].(int)),
		Fee:         input["fee"].(string),
		Categories:  sql.NullInt64{Int64: input["categories"].(int64), Valid: true},
	}

	_, err := r.DbInstance.UpdateProjectByID(context.Background(), updateParams)
	if err != nil {
		return nil, err
	}

	return r.DbInstance.GetProjectByID(context.Background(), updateParams.ID)
}

func (r *Resolver) ResolveDeleteProject(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	err := r.DbInstance.DeleteProjectByID(context.Background(), int64(id))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Resolver) ResolveInsertComment(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	insertParams := db.InsertCommentParams{
		UserID:    int64(input["user_id"].(int)),
		ProjectID: int64(input["project_id"].(int)),
		Date:      input["date"].(time.Time),
		Text:      input["text"].(string),
	}

	id, err := r.DbInstance.InsertComment(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}

	return r.DbInstance.GetCommentByID(context.Background(), id)
}

func (r *Resolver) ResolveGetCommentByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	return r.DbInstance.GetCommentByID(context.Background(), int64(id))
}

func (r *Resolver) ResolveUpdateComment(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	updateParams := db.UpdateCommentByIDParams{
		ID:        int64(input["id"].(int)),
		UserID:    int64(input["user_id"].(int)),
		ProjectID: int64(input["project_id"].(int)),
		Date:      input["date"].(time.Time),
		Text:      input["text"].(string),
	}

	_, err := r.DbInstance.UpdateCommentByID(context.Background(), updateParams)
	if err != nil {
		return nil, err
	}

	return r.DbInstance.GetCommentByID(context.Background(), updateParams.ID)
}

func (r *Resolver) ResolveDeleteComment(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	err := r.DbInstance.DeleteCommentByID(context.Background(), int64(id))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Resolver) ResolveInsertAssignProject(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	insertParams := db.InsertAssignProjectParams{
		UserID:    int64(input["user_id"].(int)),
		ProjectID: int64(input["project_id"].(int)),
	}

	assignProject, err := r.DbInstance.InsertAssignProject(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}

	return assignProject, nil
}

func (r *Resolver) ResolveDeleteAssignProject(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	deleteParams := db.DeleteAssignProjectParams{
		UserID:    int64(input["user_id"].(int)),
		ProjectID: int64(input["project_id"].(int)),
	}

	err := r.DbInstance.DeleteAssignProject(context.Background(), deleteParams)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Resolver) ResolveGetAssignedUsersByProjectID(params graphql.ResolveParams) (interface{}, error) {
	projectID := params.Args["project_id"].(int)
	return r.DbInstance.GetAssignedUsersByProjectID(context.Background(), int64(projectID))
}

func (r *Resolver) ResolveGetAssignedProjectsByUserID(params graphql.ResolveParams) (interface{}, error) {
	userID := params.Args["user_id"].(int)
	return r.DbInstance.GetAssignedprojectByUserID(context.Background(), int64(userID))
}

func (r *Resolver) ResolveInsertCategory(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	title := input["title"].(string)

	id, err := r.DbInstance.InsertCategory(context.Background(), title)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (r *Resolver) ResolveDeleteCategory(params graphql.ResolveParams) (interface{}, error) {
	id := int32(params.Args["id"].(int))
	err := r.DbInstance.DeleteCategory(context.Background(), id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Resolver) ResolveGetCategories(params graphql.ResolveParams) (interface{}, error) {
	return r.DbInstance.GetCategories(context.Background())
}

func (r *Resolver) ResolveGetCategory(params graphql.ResolveParams) (interface{}, error) {
	id := int32(params.Args["id"].(int))
	return r.DbInstance.GetCategory(context.Background(), id)
}
