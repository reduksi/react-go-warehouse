# Migrate DB Postgre

### Migrate up
`migrate -path db/migrate -database "postgres://username:password@localhost:5432/yourdbname?sslmode=disable" up`

### Migrate down
`migrate -path db/migrate -database "postgres://username:password@localhost:5432/yourdbname?sslmode=disable" down`