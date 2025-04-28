
default: build-all
# билд консоли
build-cli:
	go build ./cmd/cli

#форматирования и линтеры
f-f:
	go fmt ./...
f-v:
	go vet ./...
f-c:
	golangci-lint run
# билд основного приложения с консолью
build-app:
	go build ./cmd/app# билд основного приложения с консолью

build-app-b:
	go build ./cmd/app -ldflags -H=windowsgui
#билд фонового исполнения
build-app-d:
	go build -ldflags -H=windowsgui ./cmd/app

build-all:
	go build ./cmd/cli ./cmd/app

build-all-r:
	go build ./cmd/cli ./cmd/app -race

test-cmd:
	go test .\app\commands

pprof:
	go tool pprof

# docker
# билд мониторинга - графана + прометеус
d-mon:
	docker-compose -f=./docker/monitoring/docker-compose.yaml up -d

#запуск скрпита k6
d-k6:
	cat ./k6-script.js | docker run --rm -i grafana/k6 run -