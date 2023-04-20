package queries

const (
	CreateActivities = `
	INSERT INTO activities (
		code,
		name,
		description,
		owner,
		assignee)
	VALUES
	(?, ?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE deleted_at = NULL
	`

	DeleteActivities = `
	UPDATE activities
	INNER JOIN users as u ON u.id = a.owner
	SET a.is_deleted = true
	WHERE a.id = ?
	AND a.is_deleted = false
	AND u.username = ?
	`

	GetListActivities = `
	SELECT
		a.id,
		a.code,
		a.name,
		a.description,
		a.owner,
		a.assignee,
		a.created_at,
		a.updated_at
	FROM activities as a
	INNER JOIN users as u ON u.id = a.owner
	WHERE u.username = ?
	AND a.is_deleted = false
	LIMIT = ?
	OFFSET = ?
	ORDER BY a.created_at
	`

	GetOneActivitiesByID = `
	SELECT
		a.id,
		a.code,
		a.name,
		a.description,
		a.owner,
		a.assignee,
		a.created_at,
		a.updated_at
	FROM activities as a
	INNER JOIN users as u ON u.id = a.owner
	WHERE a.id = ?
	AND a.is_deleted = false
	AND u.username = ?
	LIMIT 1
	`

	UpdateActivities = `
	UPDATE activities
	INNER JOIN users as u ON u.id = a.owner
	SET
		a.code = ?,
		a.name = ?,
		a.description = ?,
		a.owner = ?,
		a.assignee = ?
	WHERE a.id = ?
	AND a.is_deleted = false
	AND u.username = ?
	`
)
