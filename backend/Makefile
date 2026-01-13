## run: runs the api server
run-be::
	go run ./cmd/api

## run-be-prod: runs the api server in production mode
run-be-prod::
	go run ./cmd/api -env=production

## db/psql: connects to the database using psql
db/psql:
	docker exec -it pet_shelter_db psql -U pet_admin -d shelter_db

## tidy: formats code and cleans dependencies
tidy:
	go fmt ./...
	go mod tidy

seed:
	go run ./cmd/api -seed