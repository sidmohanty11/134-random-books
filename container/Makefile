BROKER_BINARY=brokerApp

up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

up_build: build_broker
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

build_broker:
	@echo "Building broker binary..."
	cd ../broker && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} .
	@echo "Done!"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"