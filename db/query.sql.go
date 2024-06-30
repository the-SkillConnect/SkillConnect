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

const deleteAssignedProjectByID = `-- name: DeleteAssignedProjectByID :exec
DELETE FROM AssignedProject WHERE user_id = $1 AND project_id = $2
`

type DeleteAssignedProjectByIDParams struct {
	UserID    int32 `json:"user_id"`
	ProjectID int32 `json:"project_id"`
}

func (q *Queries) DeleteAssignedProjectByID(ctx context.Context, arg DeleteAssignedProjectByIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteAssignedProjectByID, arg.UserID, arg.ProjectID)
	return err
}

const deleteCommentByID = `-- name: DeleteCommentByID :exec
DELETE FROM ProjectComment WHERE id = $1
`

func (q *Queries) DeleteCommentByID(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCommentByID, id)
	return err
}

const deleteProjectByID = `-- name: DeleteProjectByID :exec
DELETE FROM Project WHERE id = $1
`

func (q *Queries) DeleteProjectByID(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteProjectByID, id)
	return err
}

const deleteUserByID = `-- name: DeleteUserByID :exec
DELETE FROM Users WHERE id = $1
`

func (q *Queries) DeleteUserByID(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUserByID, id)
	return err
}

const getAssignedProjectByID = `-- name: GetAssignedProjectByID :one
SELECT user_id, project_id, issued FROM AssignedProject WHERE user_id = $1 AND project_id = $2
`

type GetAssignedProjectByIDParams struct {
	UserID    int32 `json:"user_id"`
	ProjectID int32 `json:"project_id"`
}

func (q *Queries) GetAssignedProjectByID(ctx context.Context, arg GetAssignedProjectByIDParams) (Assignedproject, error) {
	row := q.db.QueryRowContext(ctx, getAssignedProjectByID, arg.UserID, arg.ProjectID)
	var i Assignedproject
	err := row.Scan(&i.UserID, &i.ProjectID, &i.Issued)
	return i, err
}

const getAssignedProjects = `-- name: GetAssignedProjects :many
SELECT user_id, project_id, issued FROM AssignedProject
`

func (q *Queries) GetAssignedProjects(ctx context.Context) ([]Assignedproject, error) {
	rows, err := q.db.QueryContext(ctx, getAssignedProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Assignedproject
	for rows.Next() {
		var i Assignedproject
		if err := rows.Scan(&i.UserID, &i.ProjectID, &i.Issued); err != nil {
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

const getCommentByID = `-- name: GetCommentByID :one
SELECT id, user_id, project_id, date, text FROM ProjectComment WHERE id = $1
`

func (q *Queries) GetCommentByID(ctx context.Context, id int32) (Projectcomment, error) {
	row := q.db.QueryRowContext(ctx, getCommentByID, id)
	var i Projectcomment
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProjectID,
		&i.Date,
		&i.Text,
	)
	return i, err
}

const getProjectByID = `-- name: GetProjectByID :one
SELECT id, title, description, total_amount, order_date, status, user_id, fee FROM Project WHERE id = $1
`

func (q *Queries) GetProjectByID(ctx context.Context, id int32) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProjectByID, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.TotalAmount,
		&i.OrderDate,
		&i.Status,
		&i.UserID,
		&i.Fee,
	)
	return i, err
}

const getProjects = `-- name: GetProjects :many
SELECT id, title, description, total_amount, order_date, status, user_id, fee FROM Project
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
			&i.Title,
			&i.Description,
			&i.TotalAmount,
			&i.OrderDate,
			&i.Status,
			&i.UserID,
			&i.Fee,
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
SELECT id, email, password, firstname, surname, mobile_phone FROM Users WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Firstname,
		&i.Surname,
		&i.MobilePhone,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, email, password, firstname, surname, mobile_phone FROM Users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Password,
			&i.Firstname,
			&i.Surname,
			&i.MobilePhone,
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

const insertAssignedProject = `-- name: InsertAssignedProject :one
INSERT INTO AssignedProject (user_id, project_id, issued)
VALUES ($1, $2, $3)
RETURNING user_id
`

type InsertAssignedProjectParams struct {
	UserID    int32        `json:"user_id"`
	ProjectID int32        `json:"project_id"`
	Issued    sql.NullBool `json:"issued"`
}

func (q *Queries) InsertAssignedProject(ctx context.Context, arg InsertAssignedProjectParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertAssignedProject, arg.UserID, arg.ProjectID, arg.Issued)
	var user_id int32
	err := row.Scan(&user_id)
	return user_id, err
}

const insertComment = `-- name: InsertComment :one
INSERT INTO ProjectComment (user_id, project_id, date, text)
VALUES ($1, $2, $3, $4)
RETURNING id
`

type InsertCommentParams struct {
	UserID    sql.NullInt32  `json:"user_id"`
	ProjectID sql.NullInt32  `json:"project_id"`
	Date      sql.NullTime   `json:"date"`
	Text      sql.NullString `json:"text"`
}

func (q *Queries) InsertComment(ctx context.Context, arg InsertCommentParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertComment,
		arg.UserID,
		arg.ProjectID,
		arg.Date,
		arg.Text,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const insertProject = `-- name: InsertProject :one
INSERT INTO Project (title, description, total_amount, status, user_id, fee, order_date)
VALUES ($1, $2, $3, $4, $5, $6,$7)
RETURNING id
`

type InsertProjectParams struct {
	Title       sql.NullString `json:"title"`
	Description sql.NullString `json:"description"`
	TotalAmount sql.NullString `json:"total_amount"`
	Status      sql.NullBool   `json:"status"`
	UserID      int32          `json:"user_id"`
	Fee         sql.NullString `json:"fee"`
	OrderDate   time.Time      `json:"order_date"`
}

func (q *Queries) InsertProject(ctx context.Context, arg InsertProjectParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertProject,
		arg.Title,
		arg.Description,
		arg.TotalAmount,
		arg.Status,
		arg.UserID,
		arg.Fee,
		arg.OrderDate,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO Users (email, password, firstname, surname, mobile_phone)
VALUES ($1, $2, $3, $4, $5)
RETURNING id
`

type InsertUserParams struct {
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	Firstname   sql.NullString `json:"firstname"`
	Surname     sql.NullString `json:"surname"`
	MobilePhone sql.NullString `json:"mobile_phone"`
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertUser,
		arg.Email,
		arg.Password,
		arg.Firstname,
		arg.Surname,
		arg.MobilePhone,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const updateAssignedProjectByID = `-- name: UpdateAssignedProjectByID :one
UPDATE AssignedProject
SET issued = $1
WHERE user_id = $2 AND project_id = $3
RETURNING user_id
`

type UpdateAssignedProjectByIDParams struct {
	Issued    sql.NullBool `json:"issued"`
	UserID    int32        `json:"user_id"`
	ProjectID int32        `json:"project_id"`
}

func (q *Queries) UpdateAssignedProjectByID(ctx context.Context, arg UpdateAssignedProjectByIDParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, updateAssignedProjectByID, arg.Issued, arg.UserID, arg.ProjectID)
	var user_id int32
	err := row.Scan(&user_id)
	return user_id, err
}

const updateCommentByID = `-- name: UpdateCommentByID :one
UPDATE ProjectComment
SET user_id = $1, project_id = $2, date = $3, text = $4
WHERE id = $5
RETURNING id
`

type UpdateCommentByIDParams struct {
	UserID    sql.NullInt32  `json:"user_id"`
	ProjectID sql.NullInt32  `json:"project_id"`
	Date      sql.NullTime   `json:"date"`
	Text      sql.NullString `json:"text"`
	ID        int32          `json:"id"`
}

func (q *Queries) UpdateCommentByID(ctx context.Context, arg UpdateCommentByIDParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, updateCommentByID,
		arg.UserID,
		arg.ProjectID,
		arg.Date,
		arg.Text,
		arg.ID,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const updateProjectByID = `-- name: UpdateProjectByID :one
UPDATE Project
SET title = $1, description = $2, total_amount = $3, status = $4, user_id = $5, fee = $6, order_date = $7
WHERE id = $8
RETURNING id
`

type UpdateProjectByIDParams struct {
	Title       sql.NullString `json:"title"`
	Description sql.NullString `json:"description"`
	TotalAmount sql.NullString `json:"total_amount"`
	Status      sql.NullBool   `json:"status"`
	UserID      int32          `json:"user_id"`
	Fee         sql.NullString `json:"fee"`
	OrderDate   time.Time      `json:"order_date"`
	ID          int32          `json:"id"`
}

func (q *Queries) UpdateProjectByID(ctx context.Context, arg UpdateProjectByIDParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, updateProjectByID,
		arg.Title,
		arg.Description,
		arg.TotalAmount,
		arg.Status,
		arg.UserID,
		arg.Fee,
		arg.OrderDate,
		arg.ID,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const updateUserByID = `-- name: UpdateUserByID :one
UPDATE Users
SET email = $1, password = $2, firstname = $3, surname = $4, mobile_phone = $5
WHERE id = $6
RETURNING id
`

type UpdateUserByIDParams struct {
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	Firstname   sql.NullString `json:"firstname"`
	Surname     sql.NullString `json:"surname"`
	MobilePhone sql.NullString `json:"mobile_phone"`
	ID          int32          `json:"id"`
}

func (q *Queries) UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, updateUserByID,
		arg.Email,
		arg.Password,
		arg.Firstname,
		arg.Surname,
		arg.MobilePhone,
		arg.ID,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}
