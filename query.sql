-- name: DeleteUserByID :exec
DELETE FROM Users WHERE id = $1;
-- name: UpdateUserByID :one
UPDATE Users
SET email = $1, password = $2, firstname = $3, surname = $4, mobile_phone = $5, role_id = $6
WHERE id = $7
RETURNING id;
-- name: GetUserByID :one
SELECT * FROM Users WHERE id = $1;
-- name: GetUsers :many
SELECT * FROM Users;
-- name: InsertUser :one
INSERT INTO Users (email, password, firstname, surname, mobile_phone, role_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;


-- name: DeleteCommentByID :exec
DELETE FROM ProjectComment WHERE id = $1;
-- name: UpdateCommentByID :one
UPDATE ProjectComment
SET user_id = $1, project_id = $2, date = $3, text = $4
WHERE id = $5
RETURNING id;
-- name: GetCommentByID :one
SELECT * FROM ProjectComment WHERE id = $1;
-- name: InsertComment :one
INSERT INTO ProjectComment (user_id, project_id, date, text)
VALUES ($1, $2, $3, $4)
RETURNING id;


-- name: DeleteProjectByID :exec
DELETE FROM Project WHERE id = $1;
-- name: UpdateProjectByID :one
UPDATE Project
SET title = $1, description = $2, total_amount = $3, order_date = $4, status = $5, user_id = $6, fee = $7
WHERE id = $8
RETURNING id;
-- name: GetProjectByID :one
SELECT * FROM Project WHERE id = $1;
-- name: GetProjects :many
SELECT * FROM Project;
-- name: InsertProject :one
INSERT INTO Project (title, description, total_amount, order_date, status, user_id, fee)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;


-- name: DeleteAssignedProjectByID :exec
DELETE FROM AssignedProject WHERE user_id = $1 AND project_id = $2;
-- name: UpdateAssignedProjectByID :one
UPDATE AssignedProject
SET issued = $1
WHERE user_id = $2 AND project_id = $3
RETURNING user_id;
-- name: GetAssignedProjectByID :one
SELECT * FROM AssignedProject WHERE user_id = $1 AND project_id = $2;
-- name: GetAssignedProjects :many
SELECT * FROM AssignedProject;
-- name: InsertAssignedProject :one
INSERT INTO AssignedProject (user_id, project_id, issued)
VALUES ($1, $2, $3)
RETURNING user_id;


-- name: DeleteRoleByID :exec
DELETE FROM Role WHERE id = $1; 
-- name: GetRoleByID :one
SELECT * FROM Role WHERE id = $1;
-- name: GetRoles :many
SELECT * FROM Role;
-- name: InsertRole :one
INSERT INTO Role (type)
VALUES ($1)
RETURNING id;

