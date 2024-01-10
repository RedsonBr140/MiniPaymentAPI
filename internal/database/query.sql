-- name: ListUsers :many
SELECT * FROM users 
ORDER BY first_name;

-- name: CreateUser :one
INSERT INTO users (
    first_name, last_name, cpf, email, user_type
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;