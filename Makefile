server:
	cd cmd/api && go run main.go
migrate:
	cd db/migrate && go run migrate.go

.PHONY: server migrate
