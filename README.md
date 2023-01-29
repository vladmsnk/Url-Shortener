# Url-Shortener

Run locally
```
export PG_URL=postgres://user:pass@localhost:5432/postgres
migrate -path migrations -database 'postgres://user:pass@localhost:5432/postgres?sslmode=disable' up
cd cmd && go run app/main.go
```
Run using docker-compose
```
docker-compose up -d
```
