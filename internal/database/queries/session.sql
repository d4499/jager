-- name: CreateSession :one
INSERT INTO sessions (
  id, user_id, expires_at
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetSession :one
SELECT
  sessions.id,
  sessions.user_id,
  sessions.expires_at
FROM
  sessions
WHERE
  sessions.id = @id;
