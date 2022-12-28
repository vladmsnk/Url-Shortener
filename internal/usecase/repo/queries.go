package repo

const (
	SelectRandomActivities = `
SELECT id, s_title, description, price, available_from, available_to 
FROM activities ORDER BY random() LIMIT 3;
`
)
