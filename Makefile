build:
	go build ./...
	
serve: build
	./libero
