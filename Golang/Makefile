APP_NAME = app
BUILD_DIR = $(PWD)/build

.PHONY: clean
clean:
	rm -rf ./build

.PHONY: security
security:
	gosec ./...

.PHONY: test
test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

.PHONY: build
build: clean test
	go build -o $(BUILD_DIR)/$(APP_NAME) main.go

.PHONY: swag
swag:
	swag init --parseDependency --parseInternal

.PHONY: run
run: swag build

.PHONY: run/live
run/live:
	air

.PHONY: run/prod
run/prod: swag build
		$(BUILD_DIR)/$(APP_NAME)
