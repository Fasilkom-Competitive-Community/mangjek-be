package sqlc

import "context"

const createDriver = `-- name: CreateDriver :one
INSERT INTO drivers ( id
                  , name
                  , email)
VALUES ($1, $2, $3)
RETURNING id
`

type CreateDriverParams struct {
	ID    string `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func (q *Queries) CreateDriver(ctx context.Context, arg CreateDriverParams) (string, error) {
	row := q.db.QueryRow(ctx, createDriver, arg.ID, arg.Name, arg.Email)
	var id string
	err := row.Scan(&id)
	return id, err
}

const deleteDriver = `-- name: DeleteeDriver :exec
DELETE
FROM drivers
WHERE id = $1
`

func (q *Queries) DeleteDriver(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteDriver, id)
	return err
}

const getDriver = `-- name: GetDriver :one
SELECT id
     , name
     , email
     , created_at
     , updated_at
FROM drivers
WHERE id = $1
`

func (q *Queries) GetDriver(ctx context.Context, id string) (Driver, error) {
	row := q.db.QueryRow(ctx, getDriver, id)
	var i Driver
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listDrivers = `-- name: ListDrivers :many
SELECT id
     , name
     , email
     , created_at
     , updated_at
FROM drivers
`

func (q *Queries) ListDrivers(ctx context.Context) ([]Driver, error) {
	rows, err := q.db.Query(ctx, listDrivers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Driver{}
	for rows.Next() {
		var i Driver
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDriver = `-- name: UpdateDriver :one
UPDATE drivers
SET name       = $2
  , email      = $3
  , updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id
`

type UpdateDriverParams struct {
	ID    string `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func (q *Queries) UpdateDriver(ctx context.Context, arg UpdateDriverParams) (string, error) {
	row := q.db.QueryRow(ctx, updateDriver, arg.ID, arg.Name, arg.Email)
	var id string
	err := row.Scan(&id)
	return id, err
}
