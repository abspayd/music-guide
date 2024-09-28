BINARY := app
BUILD_PATH := bin
SRC_DIR := ./cmd/server
IMAGE_NAME := abspayd/music-guide
STAGE := dev
IMAGE_TAG := latest-dev
CONTAINER_NAME := music-guide-$(STAGE)

.PHONY: all
all: build

.PHONY: build
build: $(BUILD_PATH)/$(BINARY) 
$(BUILD_PATH)/$(BINARY):
	@echo "Building the server..."
	go build -o $(BUILD_PATH)/$(BINARY) -v $(SRC_DIR)

.PHONY: watch
watch:
	@echo "Watching for changes..."
	air

.PHONY: run
run: build
	@echo "Running the server..."
	./$(BUILD_PATH)/$(BINARY)

.PHONY: test
test:
	@echo "Runnning tests..."
	go test ./...

.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -rf ./bin

.PHONY: docker-build
docker-build:
	@echo "Building docker image..."
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) --target $(STAGE) .

.PHONY: docker-run
docker-run:
	@echo "Running docker container..."
	docker run -it --rm -p 3000:3000 --name $(CONTAINER_NAME) $(IMAGE_NAME):$(IMAGE_TAG)

.PHONY: docker-watch
docker-watch:
	@echo "Running watchful docker container..."
	docker run -it --rm -p 3000:3000 -v $(shell pwd):/usr/src/app --name $(CONTAINER_NAME) $(IMAGE_NAME):$(IMAGE_TAG)

.PHONY: deploy
deploy:
	@echo "Starting deploy..."
	@echo "Building multi-platform docker image..."
	docker buildx create \
		--name builder \
		--driver docker-container \
		--use --bootstrap
	docker buildx build --platform linux/amd64,linux/arm64 -t $(IMAGE_NAME):latest --target prod --push . 
	@echo "Cleaning builder..."
	docker buildx stop builder
	docker buildx rm builder
