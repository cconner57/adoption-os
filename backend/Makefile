## run: runs the api server
run-be::
	go run ./cmd/api

## db/psql: connects to the database using psql
db/psql:
	docker exec -it pet_shelter_db psql -U pet_admin -d shelter_db

## tidy: formats code and cleans dependencies
tidy:
	go fmt ./...
	go mod tidy

seed:
	go run cmd/seeder/main.go