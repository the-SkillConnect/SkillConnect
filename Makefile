test:
	go test ./... -v

build:
	go build -o ./bin/skillConnect ./*.go

run: build
	./bin/skillConnect
	
sqlc: 
	sqlc generate