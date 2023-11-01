all:
	go run ./cmd/app/

proto:
	protoc -I ./models/ --go_out=./ ./models/person.proto
