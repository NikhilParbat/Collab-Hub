DB_URL ?= ""

clean_win:
	if exist bin rmdir /s /q bin

clean_sh:
	rm -rf bin


build_win: clean_win
	@echo "Building for Windows..."
	@go build -o bin/app.exe main.go
	@echo "Build complete."


build_sh: clean_sh
	@echo "Building for Linux..."
	@go build -o bin/app main.go
	@echo "Build complete."

run_win: build_win
	./bin/app.exe


run_sh: build_sh
	chmod +x ./bin/app
	./bin/app

migrateup:
	migrate -path db/migrations -database $(DB_URL) -verbose up

migrateup1:
	migrate -path db/migrations -database $(DB_URL) -verbose up 1

migratedown:
	migrate -path db/migrations -database $(DB_URL) -verbose down

migratedown1:
	migrate -path db/migrations -database $(DB_URL) -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

startserver:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/AscendingHeavens/gs_auth_service/db/sqlc Store

.PHONY: \
	clean_win clean_sh \
	build_win build_sh \
	run_win run_sh \
	migrateup migrateup1 \
	migratedown migratedown1 \
	sqlc test startserver mock
