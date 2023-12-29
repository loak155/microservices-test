.PHONY: gen-proto
gen-proto:
	protoc \
	-I. \
	-I=${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate \
	--go_out=. \
	--go-grpc_out=. \
	--validate_out="lang=go:." \
	proto/schema/*.proto

.PHONY: db-schema
db-schema:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml

