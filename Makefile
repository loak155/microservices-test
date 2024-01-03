.PHONY: gen-proto
gen-proto:
	protoc \
	-I. \
	--go_out=. \
	--go-grpc_out=. \
	proto/schema/*.proto

.PHONY: db-schema
db-schema:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml

