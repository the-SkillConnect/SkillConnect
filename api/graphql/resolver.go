package graphql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/the-SkillConnect/SkillConnect/db"
	"golang.org/x/crypto/bcrypt"
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

	userParams := CreateUserParams{
		FirstName:     input["first_name"].(string),
		LastName:      input["surname"].(string),
		Password:      input["password"].(string),
		Email:         input["email"].(string),
		MobilePhone:   input["mobile_phone"].(string),
		WalletAddress: input["wallet_address"].(string),
	}

	validationErrors := userParams.Validate()
	if len(validationErrors) > 0 {
		return nil, fmt.Errorf("validation errors: %v", validationErrors)
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(userParams.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	insertParams := db.InsertUserIdentityParams{
		Email:             userParams.Email,
		EncryptedPassword: string(encryptedPassword),
		FirstName:         userParams.FirstName,
		Surname:           userParams.LastName,
		MobilePhone:       userParams.MobilePhone,
		WalletAddress:     userParams.WalletAddress,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	id, err := r.DbInstance.InsertUserIdentity(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}

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

	userParams := CreateUserParams{
		FirstName:     input["first_name"].(string),
		LastName:      input["surname"].(string),
		Password:      input["password"].(string),
		Email:         input["email"].(string),
		MobilePhone:   input["mobile_phone"].(string),
		WalletAddress: input["wallet_address"].(string),
	}

	validationErrors := userParams.Validate()
	if len(validationErrors) > 0 {
		return nil, fmt.Errorf("validation errors: %v", validationErrors)
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(userParams.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	updateParams := db.UpdateUserIdentityByIDParams{
		ID:                int64(input["id"].(int)),
		Email:             userParams.Email,
		EncryptedPassword: string(encryptedPassword),
		FirstName:         userParams.FirstName,
		Surname:           userParams.LastName,
		MobilePhone:       userParams.MobilePhone,
		UpdatedAt:         time.Now(),
	}

	_, err = r.DbInstance.UpdateUserIdentityByID(context.Background(), updateParams)
	if err != nil {
		return nil, err
	}

	return r.DbInstance.GetUserIdentityByID(context.Background(), updateParams.ID)
}

func (r *Resolver) ResolveGetUsersIdentity(params graphql.ResolveParams) (interface{}, error) {
	return r.DbInstance.GetUsersIdentity(context.Background())
}

func (r *Resolver) ResolveInsertUserProfile(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	insertParams := db.InsertUserProfileParams{
		UserID:       int64(input["user_id"].(int)),
		Rating:       int64(input["rating"].(int)),
		Description:  input["description"].(string),
		DoneProject:  int64(input["done_project"].(int)),
		GivenProject: int64(input["given_project"].(int)),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	id, err := r.DbInstance.InsertUserProfile(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}

	return r.DbInstance.GetUserProfileByUserID(context.Background(), id)
}

func (r *Resolver) ResolveGetUserProfile(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	return r.DbInstance.GetUserProfileByUserID(context.Background(), int64(id))
}

func (r *Resolver) ResolveUpdateUserProfile(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	updateParams := db.UpdateUserProfileParams{
		UserID:       int64(input["user_id"].(int)),
		Rating:       input["rating"].(int64),
		Description:  input["description"].(string),
		DoneProject:  input["done_project"].(int64),
		GivenProject: input["given_project"].(int64),
		UpdatedAt:    time.Now(),
	}

	_, err := r.DbInstance.UpdateUserProfile(context.Background(), updateParams)
	if err != nil {
		return nil, err
	}

	return r.DbInstance.GetUserProfileByUserID(context.Background(), updateParams.UserID)
}

func (r *Resolver) ResolveDeleteUserProfile(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	err := r.DbInstance.DeleteUserProfileByID(context.Background(), int64(id))
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
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	id, err := r.DbInstance.InsertUserRecommendation(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}

	return r.DbInstance.GetUserRecommendationByGivenID(context.Background(), id.GivenID)
}

func (r *Resolver) ResolveDeleteUserRecommendation(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	deleteParams := db.DeleteUserRecommendationParams{
		GivenID:    int64(input["given_id"].(int)),
		ReceivedID: int64(input["received_id"].(int)),
	}

	err := r.DbInstance.DeleteUserRecommendation(context.Background(), deleteParams)
	if err != nil {
		return false, err
	}
	return true, nil
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
		CategoryID:  int64((input["category_id"].(int))),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	id, err := r.DbInstance.InsertProject(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}

	return r.DbInstance.GetProjectByID(context.Background(), int64(id))
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
		CategoryID:  int64(input["category_id"].(int)),
		UpdatedAt:   time.Now(),
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
		Date:      time.Now(),
		Text:      input["text"].(string),
	}

	id, err := r.DbInstance.InsertComment(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (r *Resolver) ResolveGetCommentByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["project_id"].(int)
	return r.DbInstance.GetProjectCommentByID(context.Background(), int64(id))
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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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
	return r.DbInstance.GetAssignedProjectByUserID(context.Background(), int64(userID))
}

func (r *Resolver) ResolveInsertCategory(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	title := input["title"].(string)

	id, err := r.DbInstance.InsertCategory(context.Background(), title)
	if err != nil {
		return nil, err
	}
	return r.DbInstance.GetCategory(context.Background(), id)
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

func (r *Resolver) ResolveGetUserProfileWithDetails(params graphql.ResolveParams) (interface{}, error) {
	id := int64(params.Args["id"].(int))
	return r.DbInstance.GetUserProfileWithDetails(context.Background(), id)
}

func (r *Resolver) ResolveGetProjectAssignments(params graphql.ResolveParams) (interface{}, error) {
	return r.DbInstance.GetProjectAssignments(context.Background())
}
