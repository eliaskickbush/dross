-- name: CreateChunk :one
INSERT INTO chunks(content, tags) VALUES (?,?) RETURNING *;
-- name: GetChunks :many
SELECT * FROM chunks LIMIT ? OFFSET ?;
