-- name: CreateMagicLink :one
INSERT INTO magic_links (
  id, email, token, expires_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetMagicLinkByToken :one
SELECT * FROM magic_links
WHERE token = $1;

-- name: DeleteMagicLink :exec
DELETE FROM magic_links WHERE id = $1;
