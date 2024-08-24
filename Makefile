BINARY := app
BUILD_PATH := bin
SRC_DIR := ./cmd/server

.PHONY: all build run clean test watch

all: build

build: $(BUILD_PATH)/$(BINARY) 
$(BUILD_PATH)/$(BINARY):
	@echo "Building the server..."
	go build -o $(BUILD_PATH)/$(BINARY) -v $(SRC_DIR)

watch:
	@echo "Watching for changes..."
	air

run: build
	@echo "Running the server..."
	./$(BUILD_PATH)/$(BINARY)

test:
	@echo "Runnning tests..."
	go test ./...

clean:
	@echo "Cleaning..."
	rm -rf ./bin
