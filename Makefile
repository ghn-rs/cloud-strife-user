tidy:
	go mod tidy
	go fmt ./internal/...
	fieldalignment -fix ./internal/...
	go vet ./internal/...
	golangci-lint run --fix ./internal/...
	staticcheck ./internal/...

	go fmt ./main.go
	fieldalignment -fix ./main.go
	go vet ./main.go
	golangci-lint run --fix ./main.go
	staticcheck ./main.go

run:
	make clean
	make proto
	make tidy
	go run main.go

install_deps:
	# These needs sudo
	# apt install build-essential -y
    # curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.1.6
	go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/google/wire/cmd/wire@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/mysql

# Configuration
PROTO_SRC_DIR := proto/src
PROTO_GEN_DIR := proto/gen
MICROSERVICES := $(notdir $(wildcard $(PROTO_SRC_DIR)/*))

# Proto generation
proto-clean:
	@echo "Cleaning generated proto files..."
	rm -rf $(PROTO_GEN_DIR)/*

proto-gen:
	@echo "Generating proto files..."
	cd . && buf generate

proto: proto-clean proto-gen

.PHONY: clean
clean:
	@echo "Cleaning generated files..."
	rm -f proto/gen/*.pb.go
	rm -f bin/cloud-strife-user

.PHONY: all
all: clean proto build