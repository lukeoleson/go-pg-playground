# go-pg Playground

#v0.0.1
Local postgresql
```shell
brew services start postgresql
psql -d pagila
brew services stop postgresql
```

golang-migrate
```
# Set the env var for the db URL
export POSTGRESQL_URL='postgres://lukeoleson:pass@localhost:5432/postgres?sslmode=disable'
# up
migrate -database ${POSTGRESQL_URL} -path db/migrations up
#down
migrate -database ${POSTGRESQL_URL} -path db/migrations down
```

# Notes
## psql commands
```
\du             #list users
\l              #list databases
\dt             #list tables
\d TABLE_NAME   #describe the table
\c db_name      #connect to a db
```