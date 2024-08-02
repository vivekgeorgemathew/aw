build:
	CGO_ENABLED=0 go build -o aw
test:
	go test ./... -v
test-cover:
	go test ./... -coverprofile=c.out
	go tool cover -html="c.out"
run: build
	./aw
docker-build:
	docker build . -t arc-wolf-test:latest
docker-run: docker-build
	docker run -p 8080:8080 arc-wolf-test:latest
docker-clean:
	docker image rm -f arc-wolf-test:latest

.PHONY:
	 test run build docker-build docker-run docker-clean test-cover