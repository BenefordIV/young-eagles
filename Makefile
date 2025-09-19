YOUNG_EAGLES_BINARY = bin/app

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

build_app:
	@echo building the young-eagles application...
	go build -o $(YOUNG_EAGLES_BINARY) ./
	@echo young-eagles application built

build_up: build_app
	@echo stopping docker images if running...
	docker-compose down
	@echo building new images
	docker-compose up --build -d
	@echo docker images built and started

down:
	@echo stopping the docker images...
	docker-compose down
	@echo docker images stopped