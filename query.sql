-- name: DeleteUserIdentityByID :exec
DELETE FROM user_identity WHERE id = $1;

-- name: UpdateUserIdentityByID :one
UPDATE user_identity
SET email = $1, password = $2, first_name = $3, surname = $4, mobile_phone = $5, updated_at = $6
WHERE id = $7
RETURNING id;

-- name: GetUserIdentityByID :one
SELECT * FROM user_identity WHERE id = $1;

-- name: GetUsersIdentity :many
SELECT * FROM user_identity;

-- name: InsertUserIdentity :one
INSERT INTO user_identity (email, password, first_name, surname, mobile_phone, wallet_address, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;

-- name: DeleteProjectByID :exec
DELETE FROM project WHERE id = $1;

-- name: UpdateProjectByID :one
UPDATE project
SET description = $1, title = $2, total_amount = $3, done_status = $4, user_id = $5, fee = $6, category_id = $7, updated_at = $8
WHERE id = $9
RETURNING id;

-- name: GetProjectByID :one
SELECT * FROM project WHERE id = $1;

-- name: GetProject :many
SELECT * FROM project;

-- name: InsertProject :one
INSERT INTO project (description, title, total_amount, done_status, user_id, fee, category_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id;

-- name: DeleteCommentByID :exec
DELETE FROM comment WHERE id = $1;

-- name: GetProjectCommentByID :many
SELECT * FROM comment WHERE project_id = $1;

-- name: InsertComment :one
INSERT INTO comment (user_id, project_id, date, text)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: DeleteAssignProject :exec
DELETE FROM assign_project WHERE user_id = $1 AND project_id = $2;

-- name: InsertAssignProject :one
INSERT INTO assign_project (user_id, project_id, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING user_id, project_id;

-- name: GetAssignedProjectByUserID :many
SELECT * FROM assign_project WHERE user_id = $1;

-- name: GetAssignedUsersByProjectID :many
SELECT * FROM assign_project WHERE project_id = $1;

-- name: UpdateUserProfile :one
UPDATE user_profile
SET rating = $1, description = $2, done_project = $3, given_project = $4, updated_at = $5
WHERE user_id = $6
RETURNING user_id;

-- name: GetUserProfileByUserID :one
SELECT * FROM user_profile WHERE user_id = $1;

-- name: InsertUserProfile :one
INSERT INTO user_profile (user_id, rating, description, done_project, given_project, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING user_id;

-- name: DeleteUserProfileByID :exec
DELETE FROM user_profile WHERE user_id = $1;

-- name: DeleteUserRecommendation :exec
DELETE FROM user_recommendation WHERE given_id = $1 AND received_id = $2;

-- name: InsertUserRecommendation :one
INSERT INTO user_recommendation (given_id, received_id, description, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING given_id, received_id;

-- name: GetUserRecommendationByGivenID :many
SELECT * FROM user_recommendation WHERE given_id = $1;

-- name: GetUserRecommendationByReceivedID :many
SELECT * FROM user_recommendation WHERE received_id = $1;

-- name: InsertCategory :one
INSERT INTO category (title)
VALUES ($1)
RETURNING id;

-- name: DeleteCategory :exec
DELETE FROM category WHERE id = $1;

-- name: GetCategory :one
SELECT * FROM category WHERE id = $1;

-- name: GetCategories :many
SELECT * FROM category;

-- name: GetUserProfileWithDetails :one
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
    ui.id = 2;


-- name: GetProjectAssignments :many
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
    user_identity ui2 ON p.user_id = ui2.id;

-- name: GetCommentWithUserAndProject :many
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
    p.id = $1;