createmigration:
	migrate create -ext=sql -dir=internal/infa/sql/migrations -seq init

.PHONY: migrate

migrate:
	migrate -path=internal/infra/sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

migratedown:
	migrate -path=internal/infra/sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down

.PHONY: migrate migratedown createmigration

