run:
	@templ generate
	@go run main.go > output.html
	@/usr/lib/chromium/chromium output.html
	# @go run main.go

testemail:
	@templ generate
	@go test -v ./pkg/html/
	@/usr/lib/chromium/chromium output.html

setup:
	@docker compose up -d

teardown:
	@docker compose down

all: build

build:
	@echo "Building..."
	@templ generate
	
	@go build -o main main.go

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/air-verse/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi
