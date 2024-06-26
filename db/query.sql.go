// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const deleteAssignProject = `-- name: DeleteAssignProject :exec
DELETE FROM assign_project WHERE user_id = $1 AND project_id = $2
`

type DeleteAssignProjectParams struct {
	UserID    int64 `json:"user_id"`
	ProjectID int64 `json:"project_id"`
}

func (q *Queries) DeleteAssignProject(ctx context.Context, arg DeleteAssignProjectParams) error {
	_, err := q.db.ExecContext(ctx, deleteAssignProject, arg.UserID, arg.ProjectID)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM category WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const deleteCommentByID = `-- name: DeleteCommentByID :exec
DELETE FROM comments WHERE id = $1
`

func (q *Queries) DeleteCommentByID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCommentByID, id)
	return err
}

const deleteProjectByID = `-- name: DeleteProjectByID :exec
DELETE FROM projects WHERE id = $1
`

func (q *Queries) DeleteProjectByID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProjectByID, id)
	return err
}

const deleteUserByID = `-- name: DeleteUserByID :exec
DELETE FROM user_identity WHERE id = $1
`

func (q *Queries) DeleteUserByID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUserByID, id)
	return err
}

const deleteUserRecommendation = `-- name: DeleteUserRecommendation :exec
DELETE FROM user_recommendation WHERE given_id = $1 AND received_id = $2
`

type DeleteUserRecommendationParams struct {
	GivenID    int64 `json:"given_id"`
	ReceivedID int64 `json:"received_id"`
}

func (q *Queries) DeleteUserRecommendation(ctx context.Context, arg DeleteUserRecommendationParams) error {
	_, err := q.db.ExecContext(ctx, deleteUserRecommendation, arg.GivenID, arg.ReceivedID)
	return err
}

const getAssignedProjectsByUserID = `-- name: GetAssignedProjectsByUserID :many
SELECT user_id, project_id, created_at, updated_at FROM assign_project WHERE user_id = $1
`

func (q *Queries) GetAssignedProjectsByUserID(ctx context.Context, userID int64) ([]AssignProject, error) {
	rows, err := q.db.QueryContext(ctx, getAssignedProjectsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AssignProject
	for rows.Next() {
		var i AssignProject
		if err := rows.Scan(
			&i.UserID,
			&i.ProjectID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAssignedUsersByProjectID = `-- name: GetAssignedUsersByProjectID :many
SELECT user_id, project_id, created_at, updated_at FROM assign_project WHERE project_id = $1
`

func (q *Queries) GetAssignedUsersByProjectID(ctx context.Context, projectID int64) ([]AssignProject, error) {
	rows, err := q.db.QueryContext(ctx, getAssignedUsersByProjectID, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AssignProject
	for rows.Next() {
		var i AssignProject
		if err := rows.Scan(
			&i.UserID,
			&i.ProjectID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategories = `-- name: GetCategories :many
SELECT id, title FROM category
`

func (q *Queries) GetCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Title); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategory = `-- name: GetCategory :one
SELECT id, title FROM category WHERE id = $1
`

func (q *Queries) GetCategory(ctx context.Context, id int32) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, id)
	var i Category
	err := row.Scan(&i.ID, &i.Title)
	return i, err
}

const getCommentByID = `-- name: GetCommentByID :one
SELECT id, user_id, project_id, date, text FROM comments WHERE id = $1
`

func (q *Queries) GetCommentByID(ctx context.Context, id int64) (Comment, error) {
	row := q.db.QueryRowContext(ctx, getCommentByID, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProjectID,
		&i.Date,
		&i.Text,
	)
	return i, err
}

const getComments = `-- name: GetComments :many
SELECT id, user_id, project_id, date, text FROM comments
`

func (q *Queries) GetComments(ctx context.Context) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, getComments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProjectID,
			&i.Date,
			&i.Text,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCommentsWithUserAndProject = `-- name: GetCommentsWithUserAndProject :many
SELECT 
    c.id AS comment_id,
    c.date,
    c.text,
    ui.id AS user_id,
    ui.email,
    ui.firstname,
    ui.surname,
    p.id AS project_id,
    p.title
FROM 
    comments c
JOIN 
    user_identity ui ON c.user_id = ui.id
JOIN 
    projects p ON c.project_id = p.id
`

type GetCommentsWithUserAndProjectRow struct {
	CommentID int64     `json:"comment_id"`
	Date      time.Time `json:"date"`
	Text      string    `json:"text"`
	UserID    int64     `json:"user_id"`
	Email     string    `json:"email"`
	Firstname string    `json:"firstname"`
	Surname   string    `json:"surname"`
	ProjectID int64     `json:"project_id"`
	Title     string    `json:"title"`
}

func (q *Queries) GetCommentsWithUserAndProject(ctx context.Context) ([]GetCommentsWithUserAndProjectRow, error) {
	rows, err := q.db.QueryContext(ctx, getCommentsWithUserAndProject)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCommentsWithUserAndProjectRow
	for rows.Next() {
		var i GetCommentsWithUserAndProjectRow
		if err := rows.Scan(
			&i.CommentID,
			&i.Date,
			&i.Text,
			&i.UserID,
			&i.Email,
			&i.Firstname,
			&i.Surname,
			&i.ProjectID,
			&i.Title,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectAssignments = `-- name: GetProjectAssignments :many
SELECT 
    ap.project_id,
    ap.user_id AS assigned_user_id,
    ap.created_at,
    ap.updated_at,
    ui1.email AS assigned_user_email,
    ui1.firstname AS assigned_user_firstname,
    ui1.surname AS assigned_user_surname,
    ui2.id AS project_owner_id,
    ui2.email AS project_owner_email,
    ui2.firstname AS project_owner_firstname,
    ui2.surname AS project_owner_surname
FROM 
    assign_project ap
JOIN 
    projects p ON ap.project_id = p.id
JOIN 
    user_identity ui1 ON ap.user_id = ui1.id
JOIN 
    user_identity ui2 ON p.user_id = ui2.id
`

type GetProjectAssignmentsRow struct {
	ProjectID             int64        `json:"project_id"`
	AssignedUserID        int64        `json:"assigned_user_id"`
	CreatedAt             time.Time    `json:"created_at"`
	UpdatedAt             sql.NullTime `json:"updated_at"`
	AssignedUserEmail     string       `json:"assigned_user_email"`
	AssignedUserFirstname string       `json:"assigned_user_firstname"`
	AssignedUserSurname   string       `json:"assigned_user_surname"`
	ProjectOwnerID        int64        `json:"project_owner_id"`
	ProjectOwnerEmail     string       `json:"project_owner_email"`
	ProjectOwnerFirstname string       `json:"project_owner_firstname"`
	ProjectOwnerSurname   string       `json:"project_owner_surname"`
}

func (q *Queries) GetProjectAssignments(ctx context.Context) ([]GetProjectAssignmentsRow, error) {
	rows, err := q.db.QueryContext(ctx, getProjectAssignments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProjectAssignmentsRow
	for rows.Next() {
		var i GetProjectAssignmentsRow
		if err := rows.Scan(
			&i.ProjectID,
			&i.AssignedUserID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AssignedUserEmail,
			&i.AssignedUserFirstname,
			&i.AssignedUserSurname,
			&i.ProjectOwnerID,
			&i.ProjectOwnerEmail,
			&i.ProjectOwnerFirstname,
			&i.ProjectOwnerSurname,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectByID = `-- name: GetProjectByID :one
SELECT id, description, title, total_amount, done_status, user_id, fee, categories, created_at, updated_at FROM projects WHERE id = $1
`

func (q *Queries) GetProjectByID(ctx context.Context, id int64) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProjectByID, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Title,
		&i.TotalAmount,
		&i.DoneStatus,
		&i.UserID,
		&i.Fee,
		&i.Categories,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProjectDetails = `-- name: GetProjectDetails :one
SELECT 
    p.id AS project_id,
    p.description,
    p.title,
    p.total_amount,
    p.done_status,
    p.user_id,
    p.fee,
    p.categories,
    p.created_at,
    p.updated_at,
    c.title AS category_title
FROM 
    projects p
LEFT JOIN 
    category c ON p.categories = c.id
WHERE 
    p.id = $1
`

type GetProjectDetailsRow struct {
	ProjectID     int64          `json:"project_id"`
	Description   string         `json:"description"`
	Title         string         `json:"title"`
	TotalAmount   string         `json:"total_amount"`
	DoneStatus    sql.NullBool   `json:"done_status"`
	UserID        int64          `json:"user_id"`
	Fee           string         `json:"fee"`
	Categories    sql.NullInt64  `json:"categories"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	CategoryTitle sql.NullString `json:"category_title"`
}

func (q *Queries) GetProjectDetails(ctx context.Context, id int64) (GetProjectDetailsRow, error) {
	row := q.db.QueryRowContext(ctx, getProjectDetails, id)
	var i GetProjectDetailsRow
	err := row.Scan(
		&i.ProjectID,
		&i.Description,
		&i.Title,
		&i.TotalAmount,
		&i.DoneStatus,
		&i.UserID,
		&i.Fee,
		&i.Categories,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CategoryTitle,
	)
	return i, err
}

const getProjects = `-- name: GetProjects :many
SELECT id, description, title, total_amount, done_status, user_id, fee, categories, created_at, updated_at FROM projects
`

func (q *Queries) GetProjects(ctx context.Context) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, getProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.Description,
			&i.Title,
			&i.TotalAmount,
			&i.DoneStatus,
			&i.UserID,
			&i.Fee,
			&i.Categories,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, password, firstname, surname, mobile_phone, wallet_address, created_at, updated_at FROM user_identity WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (UserIdentity, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i UserIdentity
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Firstname,
		&i.Surname,
		&i.MobilePhone,
		&i.WalletAddress,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserProfileByUserID = `-- name: GetUserProfileByUserID :one
SELECT user_id, rating, description, done_projects, given_projects, recommendation_id, created_at, updated_at FROM user_profile WHERE user_id = $1
`

func (q *Queries) GetUserProfileByUserID(ctx context.Context, userID int64) (UserProfile, error) {
	row := q.db.QueryRowContext(ctx, getUserProfileByUserID, userID)
	var i UserProfile
	err := row.Scan(
		&i.UserID,
		&i.Rating,
		&i.Description,
		&i.DoneProjects,
		&i.GivenProjects,
		&i.RecommendationID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserProfileWithDetails = `-- name: GetUserProfileWithDetails :one
SELECT 
    ui.id AS user_id,
    ui.email,
    ui.firstname,
    ui.surname,
    ui.mobile_phone,
    up.rating,
    up.description AS profile_description,
    up.done_projects,
    up.given_projects,
    ur.given_id,
    ur.received_id,
    ur.description AS recommendation_description
FROM 
    user_identity ui
JOIN 
    user_profile up ON ui.id = up.user_id
LEFT JOIN 
    user_recommendation ur ON up.recommendation_id = ur.given_id
WHERE 
    ui.id = $1
`

type GetUserProfileWithDetailsRow struct {
	UserID                    int64          `json:"user_id"`
	Email                     string         `json:"email"`
	Firstname                 string         `json:"firstname"`
	Surname                   string         `json:"surname"`
	MobilePhone               string         `json:"mobile_phone"`
	Rating                    int64          `json:"rating"`
	ProfileDescription        sql.NullString `json:"profile_description"`
	DoneProjects              int64          `json:"done_projects"`
	GivenProjects             int64          `json:"given_projects"`
	GivenID                   sql.NullInt64  `json:"given_id"`
	ReceivedID                sql.NullInt64  `json:"received_id"`
	RecommendationDescription sql.NullString `json:"recommendation_description"`
}

func (q *Queries) GetUserProfileWithDetails(ctx context.Context, id int64) (GetUserProfileWithDetailsRow, error) {
	row := q.db.QueryRowContext(ctx, getUserProfileWithDetails, id)
	var i GetUserProfileWithDetailsRow
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.Firstname,
		&i.Surname,
		&i.MobilePhone,
		&i.Rating,
		&i.ProfileDescription,
		&i.DoneProjects,
		&i.GivenProjects,
		&i.GivenID,
		&i.ReceivedID,
		&i.RecommendationDescription,
	)
	return i, err
}

const getUserRecommendationByGivenID = `-- name: GetUserRecommendationByGivenID :one
SELECT given_id, received_id, description FROM user_recommendation WHERE given_id = $1
`

func (q *Queries) GetUserRecommendationByGivenID(ctx context.Context, givenID int64) (UserRecommendation, error) {
	row := q.db.QueryRowContext(ctx, getUserRecommendationByGivenID, givenID)
	var i UserRecommendation
	err := row.Scan(&i.GivenID, &i.ReceivedID, &i.Description)
	return i, err
}

const getUserRecommendationByReceivedID = `-- name: GetUserRecommendationByReceivedID :one
SELECT given_id, received_id, description FROM user_recommendation WHERE received_id = $1
`

func (q *Queries) GetUserRecommendationByReceivedID(ctx context.Context, receivedID int64) (UserRecommendation, error) {
	row := q.db.QueryRowContext(ctx, getUserRecommendationByReceivedID, receivedID)
	var i UserRecommendation
	err := row.Scan(&i.GivenID, &i.ReceivedID, &i.Description)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, email, password, firstname, surname, mobile_phone, wallet_address, created_at, updated_at FROM user_identity
`

func (q *Queries) GetUsers(ctx context.Context) ([]UserIdentity, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserIdentity
	for rows.Next() {
		var i UserIdentity
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Password,
			&i.Firstname,
			&i.Surname,
			&i.MobilePhone,
			&i.WalletAddress,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertAssignProject = `-- name: InsertAssignProject :one
INSERT INTO assign_project (user_id, project_id)
VALUES ($1, $2)
RETURNING user_id, project_id
`

type InsertAssignProjectParams struct {
	UserID    int64 `json:"user_id"`
	ProjectID int64 `json:"project_id"`
}

type InsertAssignProjectRow struct {
	UserID    int64 `json:"user_id"`
	ProjectID int64 `json:"project_id"`
}

func (q *Queries) InsertAssignProject(ctx context.Context, arg InsertAssignProjectParams) (InsertAssignProjectRow, error) {
	row := q.db.QueryRowContext(ctx, insertAssignProject, arg.UserID, arg.ProjectID)
	var i InsertAssignProjectRow
	err := row.Scan(&i.UserID, &i.ProjectID)
	return i, err
}

const insertCategory = `-- name: InsertCategory :one
INSERT INTO category (title)
VALUES ($1)
RETURNING id
`

func (q *Queries) InsertCategory(ctx context.Context, title string) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertCategory, title)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const insertComment = `-- name: InsertComment :one
INSERT INTO comments (user_id, project_id, date, text)
VALUES ($1, $2, $3, $4)
RETURNING id
`

type InsertCommentParams struct {
	UserID    int64     `json:"user_id"`
	ProjectID int64     `json:"project_id"`
	Date      time.Time `json:"date"`
	Text      string    `json:"text"`
}

func (q *Queries) InsertComment(ctx context.Context, arg InsertCommentParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, insertComment,
		arg.UserID,
		arg.ProjectID,
		arg.Date,
		arg.Text,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const insertProject = `-- name: InsertProject :one
INSERT INTO projects (description, title, total_amount, done_status, user_id, fee, categories)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id
`

type InsertProjectParams struct {
	Description string        `json:"description"`
	Title       string        `json:"title"`
	TotalAmount string        `json:"total_amount"`
	DoneStatus  sql.NullBool  `json:"done_status"`
	UserID      int64         `json:"user_id"`
	Fee         string        `json:"fee"`
	Categories  sql.NullInt64 `json:"categories"`
}

func (q *Queries) InsertProject(ctx context.Context, arg InsertProjectParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, insertProject,
		arg.Description,
		arg.Title,
		arg.TotalAmount,
		arg.DoneStatus,
		arg.UserID,
		arg.Fee,
		arg.Categories,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO user_identity (email, password, firstname, surname, mobile_phone, wallet_address)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id
`

type InsertUserParams struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	Firstname     string `json:"firstname"`
	Surname       string `json:"surname"`
	MobilePhone   string `json:"mobile_phone"`
	WalletAddress int64  `json:"wallet_address"`
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, insertUser,
		arg.Email,
		arg.Password,
		arg.Firstname,
		arg.Surname,
		arg.MobilePhone,
		arg.WalletAddress,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const insertUserProfile = `-- name: InsertUserProfile :one
INSERT INTO user_profile (user_id, rating, description, done_projects, given_projects, recommendation_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING user_id
`

type InsertUserProfileParams struct {
	UserID           int64          `json:"user_id"`
	Rating           int64          `json:"rating"`
	Description      sql.NullString `json:"description"`
	DoneProjects     int64          `json:"done_projects"`
	GivenProjects    int64          `json:"given_projects"`
	RecommendationID sql.NullInt64  `json:"recommendation_id"`
}

func (q *Queries) InsertUserProfile(ctx context.Context, arg InsertUserProfileParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, insertUserProfile,
		arg.UserID,
		arg.Rating,
		arg.Description,
		arg.DoneProjects,
		arg.GivenProjects,
		arg.RecommendationID,
	)
	var user_id int64
	err := row.Scan(&user_id)
	return user_id, err
}

const insertUserRecommendation = `-- name: InsertUserRecommendation :one
INSERT INTO user_recommendation (given_id, received_id, description)
VALUES ($1, $2, $3)
RETURNING given_id, received_id
`

type InsertUserRecommendationParams struct {
	GivenID     int64  `json:"given_id"`
	ReceivedID  int64  `json:"received_id"`
	Description string `json:"description"`
}

type InsertUserRecommendationRow struct {
	GivenID    int64 `json:"given_id"`
	ReceivedID int64 `json:"received_id"`
}

func (q *Queries) InsertUserRecommendation(ctx context.Context, arg InsertUserRecommendationParams) (InsertUserRecommendationRow, error) {
	row := q.db.QueryRowContext(ctx, insertUserRecommendation, arg.GivenID, arg.ReceivedID, arg.Description)
	var i InsertUserRecommendationRow
	err := row.Scan(&i.GivenID, &i.ReceivedID)
	return i, err
}

const updateCommentByID = `-- name: UpdateCommentByID :one
UPDATE comments
SET user_id = $1, project_id = $2, date = $3, text = $4
WHERE id = $5
RETURNING id
`

type UpdateCommentByIDParams struct {
	UserID    int64     `json:"user_id"`
	ProjectID int64     `json:"project_id"`
	Date      time.Time `json:"date"`
	Text      string    `json:"text"`
	ID        int64     `json:"id"`
}

func (q *Queries) UpdateCommentByID(ctx context.Context, arg UpdateCommentByIDParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, updateCommentByID,
		arg.UserID,
		arg.ProjectID,
		arg.Date,
		arg.Text,
		arg.ID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const updateProjectByID = `-- name: UpdateProjectByID :one
UPDATE projects
SET description = $1, title = $2, total_amount = $3, done_status = $4, user_id = $5, fee = $6, categories = $7
WHERE id = $8
RETURNING id
`

type UpdateProjectByIDParams struct {
	Description string        `json:"description"`
	Title       string        `json:"title"`
	TotalAmount string        `json:"total_amount"`
	DoneStatus  sql.NullBool  `json:"done_status"`
	UserID      int64         `json:"user_id"`
	Fee         string        `json:"fee"`
	Categories  sql.NullInt64 `json:"categories"`
	ID          int64         `json:"id"`
}

func (q *Queries) UpdateProjectByID(ctx context.Context, arg UpdateProjectByIDParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, updateProjectByID,
		arg.Description,
		arg.Title,
		arg.TotalAmount,
		arg.DoneStatus,
		arg.UserID,
		arg.Fee,
		arg.Categories,
		arg.ID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const updateUserByID = `-- name: UpdateUserByID :one
UPDATE user_identity
SET email = $1, password = $2, firstname = $3, surname = $4, mobile_phone = $5
WHERE id = $6
RETURNING id
`

type UpdateUserByIDParams struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Firstname   string `json:"firstname"`
	Surname     string `json:"surname"`
	MobilePhone string `json:"mobile_phone"`
	ID          int64  `json:"id"`
}

func (q *Queries) UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, updateUserByID,
		arg.Email,
		arg.Password,
		arg.Firstname,
		arg.Surname,
		arg.MobilePhone,
		arg.ID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const updateUserProfile = `-- name: UpdateUserProfile :one
UPDATE user_profile
SET rating = $1, description = $2, done_projects = $3, given_projects = $4, recommendation_id = $5
WHERE user_id = $6
RETURNING user_id
`

type UpdateUserProfileParams struct {
	Rating           int64          `json:"rating"`
	Description      sql.NullString `json:"description"`
	DoneProjects     int64          `json:"done_projects"`
	GivenProjects    int64          `json:"given_projects"`
	RecommendationID sql.NullInt64  `json:"recommendation_id"`
	UserID           int64          `json:"user_id"`
}

func (q *Queries) UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, updateUserProfile,
		arg.Rating,
		arg.Description,
		arg.DoneProjects,
		arg.GivenProjects,
		arg.RecommendationID,
		arg.UserID,
	)
	var user_id int64
	err := row.Scan(&user_id)
	return user_id, err
}
