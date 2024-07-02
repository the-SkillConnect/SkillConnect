-- name: DeleteUserIdentityByID :exec
DELETE FROM user_identity WHERE id = $1;

-- name: UpdateUserIdentityByID :one
UPDATE user_identity
SET email = $1, password = $2, firstname = $3, surname = $4, mobile_phone = $5, updated_at = $6
WHERE id = $7
RETURNING id;

-- name: GetUserIdentityByID :one
SELECT * FROM user_identity WHERE id = $1;

-- name: GetUsersIdentity :many
SELECT * FROM user_identity;

-- name: InsertUserIdentity :one
INSERT INTO user_identity (email, password, firstname, surname, mobile_phone, wallet_address, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;

-- name: DeleteProjectByID :exec
DELETE FROM projects WHERE id = $1;

-- name: UpdateProjectByID :one
UPDATE projects
SET description = $1, title = $2, total_amount = $3, done_status = $4, user_id = $5, fee = $6, categories = $7
WHERE id = $8
RETURNING id;

-- name: GetProjectByID :one
SELECT * FROM projects WHERE id = $1;

-- name: GetProjects :many
SELECT * FROM projects;

-- name: InsertProject :one
INSERT INTO projects (description, title, total_amount, done_status, user_id, fee, categories)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;

-- name: DeleteCommentByID :exec
DELETE FROM comments WHERE id = $1;

-- name: UpdateCommentByID :one
UPDATE comments
SET user_id = $1, project_id = $2, date = $3, text = $4
WHERE id = $5
RETURNING id;

-- name: GetCommentByID :one
SELECT * FROM comments WHERE id = $1;

-- name: GetComments :many
SELECT * FROM comments;

-- name: InsertComment :one
INSERT INTO comments (user_id, project_id, date, text)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: DeleteAssignProject :exec
DELETE FROM assign_project WHERE user_id = $1 AND project_id = $2;

-- name: InsertAssignProject :one
INSERT INTO assign_project (user_id, project_id)
VALUES ($1, $2)
RETURNING user_id, project_id;

-- name: GetAssignedProjectsByUserID :many
SELECT * FROM assign_project WHERE user_id = $1;

-- name: GetAssignedUsersByProjectID :many
SELECT * FROM assign_project WHERE project_id = $1;

-- name: UpdateUserProfile :one
UPDATE user_profile
SET rating = $1, description = $2, done_projects = $3, given_projects = $4, recommendation_id = $5, updated_at = $6
WHERE user_id = $7
RETURNING user_id;

-- name: GetUserProfileByUserID :one
SELECT * FROM user_profile WHERE user_id = $1;

-- name: InsertUserProfile :one
INSERT INTO user_profile (user_id, rating, description, done_projects, given_projects, recommendation_id,created_at,updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING user_id;

-- name: DeleteUserProfileByID :exec
DELETE FROM user_profile WHERE user_id = $1;

-- name: DeleteUserRecommendation :exec
DELETE FROM user_recommendation WHERE given_id = $1 AND received_id = $2;

-- name: InsertUserRecommendation :one
INSERT INTO user_recommendation (given_id, received_id, description)
VALUES ($1, $2, $3)
RETURNING given_id, received_id;

-- name: GetUserRecommendationByGivenID :one
SELECT * FROM user_recommendation WHERE given_id = $1;

-- name: GetUserRecommendationByReceivedID :one
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
    ui.id = $1;

-- name: GetProjectDetails :one
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
    p.id = $1;

-- name: GetProjectAssignments :many
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
    user_identity ui2 ON p.user_id = ui2.id;

-- name: GetCommentsWithUserAndProject :many
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
    projects p ON c.project_id = p.id;
