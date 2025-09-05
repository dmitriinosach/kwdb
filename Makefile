
default: build-all
# билд консоли


### Format\Lint ##
f-f:
	go fmt ./...
f-v:
	go vet ./...
f-c:
	golangci-lint run

### Build ###
# app
b-app:
	go build ./cmd/app
# билд консоли
b-cli:
	go build ./cmd/cli
# app detached
b-app-d:
	go build -ldflags -H=windowsgui ./cmd/app
# билд всех приложений
b-all:
	go build ./cmd/cli ./cmd/app
# билд app с race детектором
b-all-r:
	go build ./cmd/cli ./cmd/app -race

test-cmd:
	go test .\app\commands

pprof:
	go tool pprof

### Docker ##
# билд мониторинга - графана + прометеус
d-mon:
	docker-compose -f=./docker/monitoring/docker-compose.yaml up -d
# запуск скрпита k6
d-k6:
	cat ./k6-script.js | docker run --rm -i grafana/k6 run -