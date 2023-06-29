migrate-create:
	migrate create --ext sql --dir src/database/migrations  $(name)

migrate:
	migrate -database $(url) -path src/database/migrations up
