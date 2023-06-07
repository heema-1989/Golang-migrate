-- name: CreateUser :one
INSERT INTO "users" (full_name,user_name,email,password,verify_id)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetUserEmail :one
SELECT * FROM "users" WHERE email= $1;

-- name: GetUserCredentials :one
SELECT * FROM "users" WHERE email= $1 AND password= $2;