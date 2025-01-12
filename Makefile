# Makefile for a Go project

# Build the application.
all: build test

install: tailwind-install templ-install
	@npm install
	@npm run build

templ-install:
	@if ! command -v templ > /dev/null; then \
		read -p "Go's 'templ' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/a-h/templ/cmd/templ@latest; \
			if [ ! -x "$$(command -v templ)" ]; then \
				echo "templ installation failed. Exiting..."; \
				exit 1; \
			fi; \
		else \
			echo "You chose not to install templ. Exiting..."; \
			exit 1; \
		fi; \
	fi

tailwind-install:
	@if [ ! -f tailwindcss ]; then curl -sL https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-x64 -o tailwindcss; fi
	@chmod +x tailwindcss

build: tailwind-install templ-install
	@echo "Building..."
	@templ generate
	@./tailwindcss -i web/styles/base.css -o web/assets/css/style.css
	@go build -o tmp/main cmd/api/main.go

# Run the application.
run:
	@go run cmd/api/main.go

# Test the application.
test:
	@echo "Testing..."
	@go test ./... -v

# Test the docker build.
docker:
	@echo "Testing docker build..."
	@docker build -t app .

# Clean the binary.
clean:
	@echo "Cleaning..."
	@rm -f tmp/main

# Live reload.
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

.PHONY: all build run test clean watch tailwind-install templ-install
