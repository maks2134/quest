swag:
	@echo "Generating OpenAPI schema..."
	@mkdir -p docs
	@go-swagger3 --module-path . \
		--main-file-path ./cmd/quest/main.go \
		--output ./docs/schema.json \
		--schema-without-pkg
	@echo "Schema generated successfully in ./docs/schema.json"

