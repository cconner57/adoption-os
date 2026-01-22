# ==================================================================================== #
# VARIABLES
# ==================================================================================== #
BINARY_NAME=adoption-server
BACKEND_PATH=./backend/cmd/api
FRONTEND_DIR=./frontend
DISPLAYS_DIR=./displays
PICO_TARGET=pico2-w

# ✅ CRITICAL: This ensures your deployed server knows it is in Production
PROD_FLAGS=-env=production

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

## run-be: run the Go backend (dev mode)
.PHONY: run-be
run-be:
	@echo 'Starting Backend (Dev)...'
	make -C backend run-be

## run-be-prod: run the Go backend (production mode)
.PHONY: run-be-prod
run-be-prod:
	@echo 'Starting Backend (Production)...'
	make -C backend run-be-prod

## run-fe: run the Vue frontend (dev mode)
.PHONY: run-fe
run-fe:
	@echo 'Starting Frontend...'
	cd $(FRONTEND_DIR) && npm run dev -- --host

## dev: run both backend and frontend (requires parallel execution)
.PHONY: dev
dev:
	make -j2 run-be run-fe

## db/psql: connect to the database
.PHONY: db/psql
db/psql:
	make -C backend db/psql

## seed: populate the database
.PHONY: seed
seed:
	@echo 'Seeding database...'
	make -C backend seed

## clean: remove binary and temp files
.PHONY: clean
clean:
	@echo 'Cleaning...'
	rm -f $(BINARY_NAME) $(BINARY_NAME)-new
	rm -f server.log

# ==================================================================================== #
# HARDWARE / PICO
# ==================================================================================== #

## flash: flash firmware to Pico W (targets displays folder)
.PHONY: flash
flash:
	@echo 'Flashing firmware to $(PICO_TARGET)...'
	tinygo flash --target=$(PICO_TARGET) $(DISPLAYS_DIR)/main.go

# ==================================================================================== #
# DEPLOYMENT
# ==================================================================================== #

## build-fe: build the Vue frontend for production
.PHONY: build-fe
build-fe:
	@echo 'Building Frontend static files...'
	cd $(FRONTEND_DIR) && npm run build

## update-be: zero downtime backend update only
.PHONY: update-be
update-be:
	@echo "1. Pulling latest code..."
	git pull
	@echo "2. Building new binary..."
	go build -o $(BINARY_NAME)-new $(BACKEND_PATH)
	@echo "3. Swapping binaries..."
	mv $(BINARY_NAME)-new $(BINARY_NAME)
	@echo "4. Restarting server..."
	@-pkill $(BINARY_NAME) || true
	@nohup ./$(BINARY_NAME) $(PROD_FLAGS) > server.log 2>&1 &
	@echo "✅ Backend updated successfully!"

## deploy: full update (pull -> build FE -> update BE)
.PHONY: deploy
deploy:
	@echo "🚀 Starting Full Deployment..."
	git pull
	@make build-fe
	@echo "frontend built."
	@echo "Updating backend..."
	go build -o $(BINARY_NAME)-new $(BACKEND_PATH)
	mv $(BINARY_NAME)-new $(BINARY_NAME)
	@-pkill $(BINARY_NAME) || true
	@nohup ./$(BINARY_NAME) $(PROD_FLAGS) > server.log 2>&1 &
	@echo "✅ Full Deployment Complete! (Frontend + Backend)"