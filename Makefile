# Build
BINARY := app
BUILD_PATH := bin
CSS_PATH := ./views/static/css
TEMPLATE_PATH := ./views/1

SRC_DIR := ./cmd/server
PORT := 3000

# Docker
IMAGE_NAME := abspayd/music-guide
IMAGE_TAG := latest-dev
STAGE := dev
CONTAINER_NAME := music-guide-$(STAGE)

.PHONY: all
all: build

.PHONY: build
build: $(BUILD_PATH)/$(BINARY)
$(BUILD_PATH)/$(BINARY): css templ
	@echo "Building the server..."
	go build -o $(BUILD_PATH)/$(BINARY) -v $(SRC_DIR)

.PHONY: css
css: $(CSS_PATH)/%.css
$(CSS_PATH)/%.css:
	@echo "Generating CSS..."
	npm run tailwind:build

TEMPLATES := $(shell find . -name '*.templ')
.PHONY: templ
templ: $(TEMPLATES)
$(TEMPLATES):
	templ generate

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
	docker run -it --rm --env-file ./.env -p $(PORT):$(PORT) --name $(CONTAINER_NAME) $(IMAGE_NAME):$(IMAGE_TAG)

.PHONY: docker-watch
docker-watch:
	@echo "Running watchful docker container..."
	docker run -it --rm --env-file ./.env -p $(PORT):$(PORT) -v $(shell pwd):/usr/src/app --name $(CONTAINER_NAME) $(IMAGE_NAME):$(IMAGE_TAG)

.PHONY: docker-push
docker-push:
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
