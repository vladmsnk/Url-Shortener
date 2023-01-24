package repo

const (
	InsertCreatedURL = `
	INSERT INTO urls (long_url, short_url, created_at) VALUES $1, $2. $3;
	`

	SelectLongURLByShortURL = `
	SELECT long_url FROM urls WHERE short_url = $1;
	`

	SelectShortURLByLongURL = `
	SELECT short_url FROM urls WHERE long_url = $1;
	`
)
