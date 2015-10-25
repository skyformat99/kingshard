all: build

build: kingshard

kingshard:
	go get -u -v github.com/go-sql-driver/mysql

	go build -o bin/kingshard ./cmd/kingshard

clean:
	@rm -rf bin

test:
	go test ./go/... -race
