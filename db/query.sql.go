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
DELETE FROM comment WHERE id = $1
`

func (q *Queries) DeleteCommentByID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCommentByID, id)
	return err
}

const deleteProjectByID = `-- name: DeleteProjectByID :exec
DELETE FROM project WHERE id = $1
`

func (q *Queries) DeleteProjectByID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProjectByID, id)
	return err
}

const deleteUserIdentityByID = `-- name: DeleteUserIdentityByID :exec
DELETE FROM user_identity WHERE id = $1
`

func (q *Queries) DeleteUserIdentityByID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUserIdentityByID, id)
	return err
}

const deleteUserProfileByID = `-- name: DeleteUserProfileByID :exec
DELETE FROM user_profile WHERE user_id = $1
`

func (q *Queries) DeleteUserProfileByID(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteUserProfileByID, userID)
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

const getAssignedProjectByUserID = `-- name: GetAssignedProjectByUserID :many
SELECT user_id, project_id, created_at, updated_at FROM assign_project WHERE user_id = $1
`

func (q *Queries) GetAssignedProjectByUserID(ctx context.Context, userID int64) ([]AssignProject, error) {
	rows, err := q.db.QueryContext(ctx, getAssignedProjectByUserID, userID)
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

const getCommentWithUserAndProject = `-- name: GetCommentWithUserAndProject :many
SELECT 
    c.id AS comment_id,
    c.date,
    c.text,
    ui.id AS user_id,
    ui.email,
    ui.first_name,
    ui.surname,
    p.id AS project_id,
    p.title
FROM 
    comment c
JOIN 
    user_identity ui ON c.user_id = ui.id
JOIN 
    project p ON c.project_id = p.id
WHERE 
    p.id = $1
`

type GetCommentWithUserAndProjectRow struct {
	CommentID int64     `json:"comment_id"`
	Date      time.Time `json:"date"`
	Text      string    `json:"text"`
	UserID    int64     `json:"user_id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	Surname   string    `json:"surname"`
	ProjectID int64     `json:"project_id"`
	Title     string    `json:"title"`
}

func (q *Queries) GetCommentWithUserAndProject(ctx context.Context, id int64) ([]GetCommentWithUserAndProjectRow, error) {
	rows, err := q.db.QueryContext(ctx, getCommentWithUserAndProject, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCommentWithUserAndProjectRow
	for rows.Next() {
		var i GetCommentWithUserAndProjectRow
		if err := rows.Scan(
			&i.CommentID,
			&i.Date,
			&i.Text,
			&i.UserID,
			&i.Email,
			&i.FirstName,
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

const getProject = `-- name: GetProject :many
SELECT id, description, title, total_amount, done_status, user_id, fee, category_id, created_at, updated_at FROM project
`

func (q *Queries) GetProject(ctx context.Context) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, getProject)
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
			&i.CategoryID,
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

const getProjectAssignments = `-- name: GetProjectAssignments :many
SELECT 
    ap.project_id,
    ap.user_id AS assigned_user_id,
    ap.created_at,
    ap.updated_at,
    ui1.email AS assigned_user_email,
    ui1.first_name AS assigned_user_first_name,
    ui1.surname AS assigned_user_surname,
    ui2.id AS project_owner_id,
    ui2.email AS project_owner_email,
    ui2.first_name AS project_owner_first_name,
    ui2.surname AS project_owner_surname
FROM 
    assign_project ap
JOIN 
    project p ON ap.project_id = p.id
JOIN 
    user_identity ui1 ON ap.user_id = ui1.id
JOIN 
    user_identity ui2 ON p.user_id = ui2.id
`

type GetProjectAssignmentsRow struct {
	ProjectID             int64     `json:"project_id"`
	AssignedUserID        int64     `json:"assigned_user_id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	AssignedUserEmail     string    `json:"assigned_user_email"`
	AssignedUserFirstName string    `json:"assigned_user_first_name"`
	AssignedUserSurname   string    `json:"assigned_user_surname"`
	ProjectOwnerID        int64     `json:"project_owner_id"`
	ProjectOwnerEmail     string    `json:"project_owner_email"`
	ProjectOwnerFirstName string    `json:"project_owner_first_name"`
	ProjectOwnerSurname   string    `json:"project_owner_surname"`
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
			&i.AssignedUserFirstName,
			&i.AssignedUserSurname,
			&i.ProjectOwnerID,
			&i.ProjectOwnerEmail,
			&i.ProjectOwnerFirstName,
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
SELECT id, description, title, total_amount, done_status, user_id, fee, category_id, created_at, updated_at FROM project WHERE id = $1
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
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProjectCommentByID = `-- name: GetProjectCommentByID :many
SELECT id, user_id, project_id, date, text FROM comment WHERE project_id = $1
`

func (q *Queries) GetProjectCommentByID(ctx context.Context, projectID int64) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, getProjectCommentByID, projectID)
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

const getUserIdentityByID = `-- name: GetUserIdentityByID :one
SELECT id, email, password, first_name, surname, mobile_phone, wallet_address, created_at, updated_at FROM user_identity WHERE id = $1
`

func (q *Queries) GetUserIdentityByID(ctx context.Context, id int64) (UserIdentity, error) {
	row := q.db.QueryRowContext(ctx, getUserIdentityByID, id)
	var i UserIdentity
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.FirstName,
		&i.Surname,
		&i.MobilePhone,
		&i.WalletAddress,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserProfileByUserID = `-- name: GetUserProfileByUserID :one
SELECT user_id, rating, description, done_project, given_project, created_at, updated_at FROM user_profile WHERE user_id = $1
`

func (q *Queries) GetUserProfileByUserID(ctx context.Context, userID int64) (UserProfile, error) {
	row := q.db.QueryRowContext(ctx, getUserProfileByUserID, userID)
	var i UserProfile
	err := row.Scan(
		&i.UserID,
		&i.Rating,
		&i.Description,
		&i.DoneProject,
		&i.GivenProject,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserProfileWithDetails = `-- name: GetUserProfileWithDetails :one
SELECT 
    ui.id AS user_id,
    ui.email,
    ui.first_name,
    ui.surname,
    ui.mobile_phone,
    up.rating,
    up.description AS profile_description,
    up.done_project,
    up.given_project
FROM 
    user_identity ui
JOIN 
    user_profile up ON ui.id = up.user_id
WHERE 
    ui.id = $1
`

type GetUserProfileWithDetailsRow struct {
	UserID             int64  `json:"user_id"`
	Email              string `json:"email"`
	FirstName          string `json:"first_name"`
	Surname            string `json:"surname"`
	MobilePhone        string `json:"mobile_phone"`
	Rating             int64  `json:"rating"`
	ProfileDescription string `json:"profile_description"`
	DoneProject        int64  `json:"done_project"`
	GivenProject       int64  `json:"given_project"`
}

func (q *Queries) GetUserProfileWithDetails(ctx context.Context, id int64) (GetUserProfileWithDetailsRow, error) {
	row := q.db.QueryRowContext(ctx, getUserProfileWithDetails, id)
	var i GetUserProfileWithDetailsRow
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.FirstName,
		&i.Surname,
		&i.MobilePhone,
		&i.Rating,
		&i.ProfileDescription,
		&i.DoneProject,
		&i.GivenProject,
	)
	return i, err
}

const getUserRecommendationByGivenID = `-- name: GetUserRecommendationByGivenID :many
SELECT given_id, received_id, description, created_at, updated_at FROM user_recommendation WHERE given_id = $1
`

func (q *Queries) GetUserRecommendationByGivenID(ctx context.Context, givenID int64) ([]UserRecommendation, error) {
	rows, err := q.db.QueryContext(ctx, getUserRecommendationByGivenID, givenID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserRecommendation
	for rows.Next() {
		var i UserRecommendation
		if err := rows.Scan(
			&i.GivenID,
			&i.ReceivedID,
			&i.Description,
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

const getUserRecommendationByReceivedID = `-- name: GetUserRecommendationByReceivedID :many
SELECT given_id, received_id, description, created_at, updated_at FROM user_recommendation WHERE received_id = $1
`

func (q *Queries) GetUserRecommendationByReceivedID(ctx context.Context, receivedID int64) ([]UserRecommendation, error) {
	rows, err := q.db.QueryContext(ctx, getUserRecommendationByReceivedID, receivedID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserRecommendation
	for rows.Next() {
		var i UserRecommendation
		if err := rows.Scan(
			&i.GivenID,
			&i.ReceivedID,
			&i.Description,
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

const getUsersIdentity = `-- name: GetUsersIdentity :many
SELECT id, email, password, first_name, surname, mobile_phone, wallet_address, created_at, updated_at FROM user_identity
`

func (q *Queries) GetUsersIdentity(ctx context.Context) ([]UserIdentity, error) {
	rows, err := q.db.QueryContext(ctx, getUsersIdentity)
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
			&i.FirstName,
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
INSERT INTO assign_project (user_id, project_id, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING user_id, project_id
`

type InsertAssignProjectParams struct {
	UserID    int64     `json:"user_id"`
	ProjectID int64     `json:"project_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InsertAssignProjectRow struct {
	UserID    int64 `json:"user_id"`
	ProjectID int64 `json:"project_id"`
}

func (q *Queries) InsertAssignProject(ctx context.Context, arg InsertAssignProjectParams) (InsertAssignProjectRow, error) {
	row := q.db.QueryRowContext(ctx, insertAssignProject,
		arg.UserID,
		arg.ProjectID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
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
INSERT INTO comment (user_id, project_id, date, text)
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
INSERT INTO project (description, title, total_amount, done_status, user_id, fee, category_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id
`

type InsertProjectParams struct {
	Description string       `json:"description"`
	Title       string       `json:"title"`
	TotalAmount string       `json:"total_amount"`
	DoneStatus  sql.NullBool `json:"done_status"`
	UserID      int64        `json:"user_id"`
	Fee         string       `json:"fee"`
	CategoryID  int64        `json:"category_id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

func (q *Queries) InsertProject(ctx context.Context, arg InsertProjectParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, insertProject,
		arg.Description,
		arg.Title,
		arg.TotalAmount,
		arg.DoneStatus,
		arg.UserID,
		arg.Fee,
		arg.CategoryID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const insertUserIdentity = `-- name: InsertUserIdentity :one
INSERT INTO user_identity (email, password, first_name, surname, mobile_phone, wallet_address, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id
`

type InsertUserIdentityParams struct {
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	FirstName     string    `json:"first_name"`
	Surname       string    `json:"surname"`
	MobilePhone   string    `json:"mobile_phone"`
	WalletAddress string    `json:"wallet_address"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (q *Queries) InsertUserIdentity(ctx context.Context, arg InsertUserIdentityParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, insertUserIdentity,
		arg.Email,
		arg.Password,
		arg.FirstName,
		arg.Surname,
		arg.MobilePhone,
		arg.WalletAddress,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const insertUserProfile = `-- name: InsertUserProfile :one
INSERT INTO user_profile (user_id, rating, description, done_project, given_project, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING user_id
`

type InsertUserProfileParams struct {
	UserID       int64     `json:"user_id"`
	Rating       int64     `json:"rating"`
	Description  string    `json:"description"`
	DoneProject  int64     `json:"done_project"`
	GivenProject int64     `json:"given_project"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (q *Queries) InsertUserProfile(ctx context.Context, arg InsertUserProfileParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, insertUserProfile,
		arg.UserID,
		arg.Rating,
		arg.Description,
		arg.DoneProject,
		arg.GivenProject,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var user_id int64
	err := row.Scan(&user_id)
	return user_id, err
}

const insertUserRecommendation = `-- name: InsertUserRecommendation :one
INSERT INTO user_recommendation (given_id, received_id, description, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING given_id, received_id
`

type InsertUserRecommendationParams struct {
	GivenID     int64     `json:"given_id"`
	ReceivedID  int64     `json:"received_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type InsertUserRecommendationRow struct {
	GivenID    int64 `json:"given_id"`
	ReceivedID int64 `json:"received_id"`
}

func (q *Queries) InsertUserRecommendation(ctx context.Context, arg InsertUserRecommendationParams) (InsertUserRecommendationRow, error) {
	row := q.db.QueryRowContext(ctx, insertUserRecommendation,
		arg.GivenID,
		arg.ReceivedID,
		arg.Description,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i InsertUserRecommendationRow
	err := row.Scan(&i.GivenID, &i.ReceivedID)
	return i, err
}

const updateProjectByID = `-- name: UpdateProjectByID :one
UPDATE project
SET description = $1, title = $2, total_amount = $3, done_status = $4, user_id = $5, fee = $6, category_id = $7, updated_at = $8
WHERE id = $9
RETURNING id
`

type UpdateProjectByIDParams struct {
	Description string       `json:"description"`
	Title       string       `json:"title"`
	TotalAmount string       `json:"total_amount"`
	DoneStatus  sql.NullBool `json:"done_status"`
	UserID      int64        `json:"user_id"`
	Fee         string       `json:"fee"`
	CategoryID  int64        `json:"category_id"`
	UpdatedAt   time.Time    `json:"updated_at"`
	ID          int64        `json:"id"`
}

func (q *Queries) UpdateProjectByID(ctx context.Context, arg UpdateProjectByIDParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, updateProjectByID,
		arg.Description,
		arg.Title,
		arg.TotalAmount,
		arg.DoneStatus,
		arg.UserID,
		arg.Fee,
		arg.CategoryID,
		arg.UpdatedAt,
		arg.ID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const updateUserIdentityByID = `-- name: UpdateUserIdentityByID :one
UPDATE user_identity
SET email = $1, password = $2, first_name = $3, surname = $4, mobile_phone = $5, updated_at = $6
WHERE id = $7
RETURNING id
`

type UpdateUserIdentityByIDParams struct {
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	FirstName   string    `json:"first_name"`
	Surname     string    `json:"surname"`
	MobilePhone string    `json:"mobile_phone"`
	UpdatedAt   time.Time `json:"updated_at"`
	ID          int64     `json:"id"`
}

func (q *Queries) UpdateUserIdentityByID(ctx context.Context, arg UpdateUserIdentityByIDParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, updateUserIdentityByID,
		arg.Email,
		arg.Password,
		arg.FirstName,
		arg.Surname,
		arg.MobilePhone,
		arg.UpdatedAt,
		arg.ID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const updateUserProfile = `-- name: UpdateUserProfile :one
UPDATE user_profile
SET rating = $1, description = $2, done_project = $3, given_project = $4, updated_at = $5
WHERE user_id = $6
RETURNING user_id
`

type UpdateUserProfileParams struct {
	Rating       int64     `json:"rating"`
	Description  string    `json:"description"`
	DoneProject  int64     `json:"done_project"`
	GivenProject int64     `json:"given_project"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserID       int64     `json:"user_id"`
}

func (q *Queries) UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, updateUserProfile,
		arg.Rating,
		arg.Description,
		arg.DoneProject,
		arg.GivenProject,
		arg.UpdatedAt,
		arg.UserID,
	)
	var user_id int64
	err := row.Scan(&user_id)
	return user_id, err
}
