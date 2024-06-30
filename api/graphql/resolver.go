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

func (r *Resolver) ResolveGetUserByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	return r.DbInstance.GetUserByID(context.Background(), int32(id))
}

func (r *Resolver) ResolveGetProjectByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	return r.DbInstance.GetProjectByID(context.Background(), int32(id))
}
func (r *Resolver) ResolveGetProjectCommentByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	return r.DbInstance.GetProjectCommentByID(context.Background(), int32(id))
}

func (r *Resolver) ResolveGetProjectCommentsByProjectID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["project_id"].(int)
	return r.DbInstance.GetProjectCommentsByProjectID(context.Background(), int32(id))
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
	fmt.Printf("%+v\n", input)
	// orderDate := time.Now().Format("2006-01-02 15:04:05")
	insertParams := db.InsertProjectParams{
		Title:       sql.NullString{String: input["title"].(string), Valid: input["title"] != nil},
		Description: sql.NullString{String: input["description"].(string), Valid: input["description"] != nil},
		TotalAmount: sql.NullString{String: input["total_amount"].(string), Valid: input["total_amount"] != nil},
		Status:      sql.NullBool{Bool: input["status"].(bool), Valid: input["status"] != nil},
		OrderDate:   sql.NullTime{Time: time.Now(), Valid: true},
		UserID:      int32(input["user_id"].(int)),
		Fee:         sql.NullString{String: input["fee"].(string), Valid: input["fee"] != nil},
	}
	fmt.Printf("%+v\n", insertParams)
	id, err := r.DbInstance.InsertProject(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}
	return r.DbInstance.GetProjectByID(context.Background(), id)
}

func (r *Resolver) ResolveInsertProjectComment(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	insertParams := db.InsertProjectCommentParams{
		UserID:    int32(input["user_id"].(int)),
		ProjectID: int32(input["project_id"].(int)),
		Date:      sql.NullTime{Time: time.Now(), Valid: true},
		Text:      sql.NullString{String: input["text"].(string), Valid: input["text"] != nil},
	}
	fmt.Printf("%+v\n", insertParams)
	id, err := r.DbInstance.InsertProjectComment(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}
	return r.DbInstance.GetProjectCommentByID(context.Background(), id)
}

func (r *Resolver) ResolveDeleteUser(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	err := r.DbInstance.DeleteUserByID(context.Background(), int32(id))
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r *Resolver) ResolveDeleteProject(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	err := r.DbInstance.DeleteProjectByID(context.Background(), int32(id))
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r *Resolver) ResolveDeleteProjectComment(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	err := r.DbInstance.DeleteProjectCommentByID(context.Background(), int32(id))
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

func (r *Resolver) ResolveUpdateProject(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	updateParams := db.UpdateProjectByIDParams{
		ID:          int32(input["id"].(int)),
		Title:       sql.NullString{String: input["title"].(string), Valid: input["title"] != nil},
		Description: sql.NullString{String: input["description"].(string), Valid: input["description"] != nil},
		TotalAmount: sql.NullString{String: input["total_amount"].(string), Valid: input["total_amount"] != nil},
		Status:      sql.NullBool{Bool: input["status"].(bool), Valid: input["status"] != nil},
		OrderDate:   sql.NullTime{Time: time.Now(), Valid: true},
		UserID:      int32(input["user_id"].(int)),
		Fee:         sql.NullString{String: input["fee"].(string), Valid: input["fee"] != nil},
	}
	_, err := r.DbInstance.UpdateProjectByID(context.Background(), updateParams)
	if err != nil {
		return nil, err
	}
	return r.DbInstance.GetProjectByID(context.Background(), updateParams.ID)
}

// type UpdateProjectCommentByIDParams struct {
// 	UserID    sql.NullInt32  `json:"user_id"`
// 	ProjectID sql.NullInt32  `json:"project_id"`
// 	Date      sql.NullTime   `json:"date"`
// 	Text      sql.NullString `json:"text"`
// 	ID        int32          `json:"id"`
// }

func (r *Resolver) ResolveUpdateProjectComment(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	updateParams := db.UpdateProjectCommentByIDParams{
		ID:        int32(input["id"].(int)),
		UserID:    int32(input["user_id"].(int)),
		ProjectID: int32(input["project_id"].(int)),
		Date:      sql.NullTime{Time: time.Now(), Valid: true},
		Text:      sql.NullString{String: input["text"].(string), Valid: input["text"] != nil},
	}
	fmt.Printf("%+v\n", updateParams)
	_, err := r.DbInstance.UpdateProjectCommentByID(context.Background(), updateParams)
	if err != nil {
		return nil, err
	}
	return r.DbInstance.GetProjectCommentByID(context.Background(), updateParams.ID)
}

func (r *Resolver) ResolveGetUsers(params graphql.ResolveParams) (interface{}, error) {
	return r.DbInstance.GetUsers(context.Background())
}

func (r *Resolver) ResolveGetProjects(params graphql.ResolveParams) (interface{}, error) {
	return r.DbInstance.GetProjects(context.Background())
}

func (r *Resolver) ResolveGetAssignedProjectByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["project_id"].(int)
	return r.DbInstance.GetAssignedProjectByID(context.Background(), int32(id))
}
func (r *Resolver) ResolveGetAssignedProjects(params graphql.ResolveParams) (interface{}, error) {
	return r.DbInstance.GetAssignedProjects(context.Background())
}

func (r *Resolver) ResolveInsertAssignedProject(params graphql.ResolveParams) (interface{}, error) {

	input := params.Args["input"].(map[string]interface{})
	insertParams := db.InsertAssignedProjectParams{
		UserID:    int32(input["user_id"].(int)),
		ProjectID: int32(input["project_id"].(int)),
		Issued:    sql.NullBool{Bool: input["issued"].(bool), Valid: input["issued"] != nil},
	}
	id, err := r.DbInstance.InsertAssignedProject(context.Background(), insertParams)
	if err != nil {
		return nil, err
	}
	return r.DbInstance.GetAssignedProjectByID(context.Background(), id)
}

func (r *Resolver) ResolveUpdateAssignedProject(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	updateParams := db.UpdateAssignedProjectByIDParams{
		UserID:    int32(input["user_id"].(int)),
		ProjectID: int32(input["project_id"].(int)),
		Issued:    sql.NullBool{Bool: input["issued"].(bool), Valid: input["issued"] != nil},
	}
	_, err := r.DbInstance.UpdateAssignedProjectByID(context.Background(), updateParams)
	if err != nil {
		return nil, err
	}
	return r.DbInstance.GetAssignedProjectByID(context.Background(), updateParams.ProjectID)
}

func (r *Resolver) ResolveDeleteAssignedProject(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["project_id"].(int)
	err := r.DbInstance.DeleteAssignedProjectByID(context.Background(), int32(id))
	if err != nil {
		return false, err
	}
	return true, nil
}
