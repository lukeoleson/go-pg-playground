POSTGRESQL_URL='postgres://lukeoleson:pass@localhost:5432/postgres?sslmode=disable'
MIGRATIONS_URL='db/migrations'

.PHONY: db-up
db-up:
	echo 'create database pagila;' | psql -U lukeoleson -d postgres
	cat data/pagila/pagila-schema.sql | psql -U lukeoleson -d pagila
	cat data/pagila/pagila-data.sql | psql -U lukeoleson -d pagila

.PHONY: db-down
db-down:
	echo '\connect postgres;' | psql -U lukeoleson -d postgres
	echo 'drop database pagila;' | psql -U lukeoleson -d postgres