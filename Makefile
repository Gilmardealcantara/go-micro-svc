run:
	go run ./cmd/api/main.go

doc:
	swag fmt && swag init -g cmd/api/main.go --output docs/swagger

test_u:
	go test `go list ./... | grep -v  -E 'pkg/api/controllers|/cmd/api/|/docs|/tests'` -coverprofile=coverage.txt -covermode=atomic -tags testing
test_i:
	go test `go list ./... | grep /pkg/api/controllers` -coverprofile=coverage.txt -covermode=atomic -tags testing -v
test_ir:
	./scripts/integration_test_rancher.sh
coverage:
	go tool cover -func=coverage.txt | grep total | grep -Eo '[0-9]+\.[0-9]+'
coverage_html:
	go tool cover -html=coverage.txt


lint:
	golangci-lint run

db_conn:
	psql postgresql://go-service-template:password@0.0.0.0:5433/go-service-template
db_migrate_create:
	migrate create -ext sql -dir db/migrations -seq $(name)
db_migrate_up:
	migrate -database "postgresql://go-service-template:password@0.0.0.0:5433/go-service-template?sslmode=disable" -path db/migrations up
db_migrate_down:
	migrate -database "postgresql://go-service-template:password@0.0.0.0:5433/go-service-template?sslmode=disable" -path db/migrations down
db_migrate_force:
	migrate -database "postgresql://go-service-template:password@0.0.0.0:5433/go-service-template?sslmode=disable" -path db/migrations force $(version)
db_migrate_goto:
	migrate -database "postgresql://go-service-template:password@0.0.0.0:5433/go-service-template?sslmode=disable" -path db/migrations goto $(version)
