-- name: CreateURL :one
INSERT INTO url (url_encoded,
                 url_original,
                 expiration_date)
VALUES ($1, $2, $3)RETURNING url_encoded, expiration_date;

-- name: GetURL :one
SELECT *
FROM url
WHERE url_encoded = $1 LIMIT 1;

-- name: ListURL :many
SELECT *
FROM url
ORDER BY $1 LIMIT $2
OFFSET $3;

-- name: UpdateURL :one
UPDATE url
SET expiration_date = $2
WHERE url_encoded = $1 RETURNING url_encoded, expiration_date;

-- name: DeleteURL :exec
DELETE
FROM url
WHERE url_encoded = $1;

-- name: DeleteDeprecatedURL :exec
DELETE FROM url
WHERE expiration_date < CURRENT_DATE;