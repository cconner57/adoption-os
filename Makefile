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

## run-be-prod: run the Go backend in production mode
.PHONY: run-be-prod
run-be-prod:
	@echo 'Starting Backend (Production)...'
	make -C backend run-be-prod

## run-fe: run the Vue frontend
.PHONY: run-fe
run-fe:
	@echo 'Starting Frontend...'
	cd frontend && npm run dev -- --host

## build-fe: build the Vue frontend for production
.PHONY: build-fe
build-fe:
	@echo 'Building Frontend...'
	cd frontend && npm run build

## db/psql: connect to the database
.PHONY: db/psql
db/psql:
	make -C backend db/psql

## seed: populate the database from the animals.csv file
.PHONY: seed
seed:
	@echo 'Seeding the database...'
	make -C backend seed

## flash: flash the firmware to the Pico W
.PHONY: flash
flash:
	@echo 'Flashing firmware to Pico W...'
	tinygo flash --target=pico2-w .

## update: zero downtime server update
.PHONY: update-be
update-be:
	@echo "1. Pulling latest code..."
	git pull
	@echo "2. Building new binary (while old one keeps running)..."
	go build -o adoption-server-new ./backend/cmd/api
	@echo "3. Swapping binaries..."
	mv adoption-server-new adoption-server
	@echo "4. Restarting server..."
	@-pkill adoption-server || true
	@nohup ./adoption-server > server.log 2>&1 &
	@echo "✅ Success! New server is running."