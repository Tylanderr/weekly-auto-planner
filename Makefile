run:
	@templ generate
	@go run main.go > output.html
	# @/bin/zen-browser output.html
	@/usr/lib/chromium/chromium output.html

setup:
	@docker compose up -d

teardown:
	@docker compose down
