// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	DeleteAssignedProjectByID(ctx context.Context, arg DeleteAssignedProjectByIDParams) error
	DeleteCommentByID(ctx context.Context, id int32) error
	DeleteProjectByID(ctx context.Context, id int32) error
	DeleteRoleByID(ctx context.Context, id int32) error
	DeleteUserByID(ctx context.Context, id int32) error
	GetAssignedProjectByID(ctx context.Context, arg GetAssignedProjectByIDParams) (Assignedproject, error)
	GetAssignedProjects(ctx context.Context) ([]Assignedproject, error)
	GetCommentByID(ctx context.Context, id int32) (Projectcomment, error)
	GetProjectByID(ctx context.Context, id int32) (Project, error)
	GetProjects(ctx context.Context) ([]Project, error)
	GetRoleByID(ctx context.Context, id int32) (Role, error)
	GetRoles(ctx context.Context) ([]Role, error)
	GetUserByID(ctx context.Context, id int32) (User, error)
	GetUsers(ctx context.Context) ([]User, error)
	InsertAssignedProject(ctx context.Context, arg InsertAssignedProjectParams) (int32, error)
	InsertComment(ctx context.Context, arg InsertCommentParams) (int32, error)
	InsertProject(ctx context.Context, arg InsertProjectParams) (int32, error)
	InsertRole(ctx context.Context, type_ string) (int32, error)
	InsertUser(ctx context.Context, arg InsertUserParams) (int32, error)
	UpdateAssignedProjectByID(ctx context.Context, arg UpdateAssignedProjectByIDParams) (int32, error)
	UpdateCommentByID(ctx context.Context, arg UpdateCommentByIDParams) (int32, error)
	UpdateProjectByID(ctx context.Context, arg UpdateProjectByIDParams) (int32, error)
	UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) (int32, error)
}

var _ Querier = (*Queries)(nil)
