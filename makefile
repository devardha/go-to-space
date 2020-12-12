dev:
	go run main.go

build: main.go
	go build -o bin/gotospace main.go

clean:
	rm -rf bin