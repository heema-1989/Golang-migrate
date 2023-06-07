-- name: CreateRoom :one
INSERT INTO "rooms" (room_name,user_id)
VALUES ($1,$2) RETURNING *;

-- name: GetRooms :many
SELECT rooms.id,rooms.room_name, users.full_name AS created_by FROM "rooms" INNER JOIN "users" ON rooms.user_id=users.id;

