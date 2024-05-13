migrate_database:
	migrate -database postgres://postgres:password@localhost:5432/core_sample?sslmode=disable -path db/migrations up

migrate_database_down:
	migrate -database postgres://postgres:@localhost:5432/udestiny?sslmode=disable -path db/migrations down 1