-- name: GetUserByEmail :one
SELECT *
FROM users
where email = $1
limit 1;

-- name: CreateUser :exec
INSERT INTO users (email, email_confirmed, password_hash, confirm_selector, confirm_verifier, recovery_selector, recovery_verifier, recovery_expiry)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8);

-- name: UpdateUser :exec
UPDATE users
SET password_hash     = $1,
    email_confirmed   = $2,
    confirm_selector  = $3,
    confirm_verifier  = $4,
    recovery_selector = $5,
    recovery_verifier = $6,
    recovery_expiry   = $7
WHERE email = $1;

-- name: CreateListForUser :one
INSERT INTO filter_lists (user_id)
VALUES ($1)
RETURNING token;

-- name: GetListForUser :one
SELECT token,
       downloaded,
       (SELECT COUNT(*) FROM filter_instances WHERE filter_instances.user_id = $1) AS instance_count
FROM filter_lists
WHERE filter_lists.user_id = $1
LIMIT 1;

-- name: CountListsForUser :one
SELECT COUNT(*)
FROM filter_lists
WHERE user_id = $1;

-- name: RotateListToken :exec
UPDATE filter_lists
SET token      = gen_random_uuid(),
    downloaded = false
WHERE user_id = $1
  AND token = $2;

-- name: HasUserDownloadedList :one
SELECT downloaded
FROM filter_lists
WHERE filter_lists.user_id = $1
LIMIT 1;

-- name: GetListForToken :one
SELECT id, downloaded
FROM filter_lists
WHERE token = $1
LIMIT 1;

-- name: MarkListDownloaded :exec
UPDATE filter_lists
SET downloaded = true
WHERE id = $1;

-- name: GetActiveFiltersForUser :many
SELECT DISTINCT filter_name
FROM filter_instances
WHERE user_id = $1;

-- name: CreateInstanceForUserAndFilter :exec
INSERT INTO filter_instances (filter_list_id, user_id, filter_name, params)
VALUES ((SELECT id FROM filter_lists WHERE user_id = $1), $1, $2, $3);

-- name: UpdateInstanceForUserAndFilter :exec
UPDATE filter_instances
SET params     = $3,
    updated_at = NOW()
WHERE (user_id = $1 AND filter_name = $2);

-- name: GetInstanceForUserAndFilter :one
SELECT params
FROM filter_instances
WHERE (user_id = $1 AND filter_name = $2);

-- name: CountInstanceForUserAndFilter :one
SELECT COUNT(*)
FROM filter_instances
WHERE (user_id = $1 AND filter_name = $2);

-- name: DeleteInstanceForUserAndFilter :exec
DELETE
FROM filter_instances
WHERE (user_id = $1 AND filter_name = $2);

-- name: GetInstancesForList :many
SELECT filter_name, params
FROM filter_instances
WHERE filter_list_id = $1
ORDER BY filter_name ASC;

-- name: GetStats :one
SELECT (SELECT COUNT(*) FROM filter_lists)                          as list_count,
       (SELECT COUNT(*) FROM filter_lists WHERE downloaded IS TRUE) as active_list_count,
       (SELECT COUNT(*) FROM filter_instances)                      as instance_count;
