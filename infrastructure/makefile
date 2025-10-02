PRODUCT_BINARY=productApp

up:
	@echo "Starting docker images"
	docker compose up
	@echo "Docker image started"

up_build: build_product
	@echo "Stopping docker images (if running ...)"
	docker compose down
	@echo "Building (when required) and starting docker images ..."
	docker compose up --build -d
	@echo "Docker image build and started"

down:
	@echo "Stopping docker image"
	docker compose down -v
	@echo "Done!"

build_product:
	@echo "Building docker binary"
	cd ../product-service && env GOOS=linux CGO_ENABLED=0 go build -o ${PRODUCT_BINARY} ./cmd/api
	@echo "Done!"