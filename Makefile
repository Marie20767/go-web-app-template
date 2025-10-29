start:
	go run main.go

start-db:
	docker-compose up -d

stop-db:
	docker-compose down

lint: lint/install lint/run

lint/install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v2.5.0

lint/run:
	bin/golangci-lint run --config .golangci.yaml
