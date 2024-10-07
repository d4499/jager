-- name: CreateJobApplication :one
INSERT INTO job_applications (
  id, title, company, user_id, applied_date
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetAllJobApplications :many
SELECT * FROM job_applications
WHERE user_id = $1;

-- name: DeleteJobApplication :exec
DELETE FROM job_applications WHERE id = $1;
