APP_NAME = gon
VERSION ?= dev
BUILD_DIR = bin
GORELEASER ?= go tool goreleaser

.PHONY: all build install clean version

all: build

build:
	@echo "🔨 Building $(APP_NAME)..."
	go build -ldflags "-X github.com/mickamy/gon/cmd/version.Version=$(VERSION)" -o $(BUILD_DIR)/$(APP_NAME) .

install:
	@echo "📦 Installing $(APP_NAME)..."
	go install -ldflags "-X github.com/mickamy/gon/cmd/version.Version=$(VERSION)"

clean:
	@echo "🧹 Cleaning up..."
	rm -rf $(BUILD_DIR)

version:
	@echo "🔖 Version: $(VERSION)"

release:
	@echo "🚀 Running release..."
	$(GORELEASER) release --clean

snapshot:
	@echo "🔍 Running snapshot release (dry run)..."
	$(GORELEASER) release --snapshot --clean
