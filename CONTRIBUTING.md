#Migrations
use https://github.com/golang-migrate/migrate

On Windows download from cmd 
go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate

to create migrate:
migrate create -ext sql -dir migrations create_users

to make migration:

UP : migrate -path  migrations -database "postgres://user:password@host:port/dbname?sslmode=disable" up
DOWN : migrate -path  migrations -database "postgres://user:password@host:port/dbname?sslmode=disable" down