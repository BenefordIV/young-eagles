build:
	go build -o bin/app main.go

run: build
	./bin/app

test:
	go test ./...

clean:
	rm -f bin/app

code-gen-db:
	sqlboiler psql
	findstr /rl "DeletedAt" internal/dbmodels | xargs sed /i = "" 's/DeletedAt/DeletedTS/g'
	findstr /rl "deleted_at" internal/dbmodels | xargs sed /i = "" 's/deleted_at/deleted_ts/g'

run:
	go run ./.