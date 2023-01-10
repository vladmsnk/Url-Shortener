package repo

const (
	SelectRandomActivities = `
SELECT id, s_title, description, price, available_from, available_to 
FROM activities ORDER BY random() LIMIT 3;
`
	InsertSelection = `
INSERT INTO selection (id, user_id, title, created_at) VALUES
$1, $2. $3, $4 ON CONFLICT (id) DO NOTHING;`

	InsertActivitiesForSelection = `
INSERT INTO activities_for_selection (selection_id, activity_id)
$1, UNNEST($2);`
)
