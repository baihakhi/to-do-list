package queries

const (
	CreateUsers = `
	INSERT INTO users (code, name, user_group, email, password)
	VALUES($1, $2, $3, $4, $5)
	RETURNING code
	`

	DeleteUsers = `
	DELETE from users
	WHERE
		users.code = $1
	`

	GetListUsers = `
	SELECT
		u.id,
		u.code,
		u.name,
		u.user_group,
		u.email,
		u.status,
		u.created_at,
		u.updated_at
	FROM users as u
	WHERE u.name LIKE '%' || $3 || '%'
	ORDER BY u.created_at
	LIMIT $1
	OFFSET $2
	`

	GetOneUsersByCode = `
	SELECT
		u.id,
		u.code,
		u.name,
		u.user_group,
		u.email,
		u.status,
		u.created_at,
		u.updated_at
	FROM users as u
	WHERE u.code = $1
	LIMIT 1
	`

	UpdateUsers = `
	UPDATE activities
	INNER JOIN users as u ON u.id = a.owner
	SET
		a.code = $1,
		a.name = $2,
		a.description = $3,
		a.owner = $4,
		a.assignee = $5
	WHERE a.id = $6
	AND a.is_deleted = false
	AND u.username = $7
	`
)
