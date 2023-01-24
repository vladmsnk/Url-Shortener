# Url-Shortener

Run locally
```
export PG_URL=postgres://user:pass@localhost:5432/postgres
migrate -path migrations -database 'postgres://user:pass@localhost:5432/postgres?sslmode=disable' up
cd cmd && go run app/main.go
```
And use BloomRPC to send RPC

