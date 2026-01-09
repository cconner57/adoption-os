# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run-be: run the Go backend (api)
.PHONY: run-be
run-be:
	@echo 'Starting Backend...'
	make -C backend run-be

## run-fe: run the Vue frontend
.PHONY: run-fe
run-fe:
	@echo 'Starting Frontend...'
	cd frontend && npm run dev -- --host

## db/psql: connect to the database
.PHONY: db/psql
db/psql:
	make -C backend db/psql

## seed: populate the database from the animals.csv file
.PHONY: seed
seed:
	@echo 'Seeding the database...'
	make -C backend seed