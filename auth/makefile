# === Format Go code
fmt:
	go fmt ./...

# === Run the Go server (Go code handles .env)
run:
	@echo "Running Go server..."
	go run ./cmd/main.go

# === Start MongoDB container
mongo-up:
	docker-compose up -d mongo

# === Stop MongoDB container
mongo-down:
	docker-compose down

# === Run tests
test:
	go test ./...

.PHONY: fmt run mongo-up mongo-down test
