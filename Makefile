


default: build-all
# билд консоли
build-cli:
	go build ./cmd/cli

# билд основного приложения с консолью
build-app:
	go build ./cmd/app
#билд фонового исполнения
build-app-d:
	go build -ldflags -H=windowsgui ./cmd/app

build-all:
	go build ./cmd/cli ./cmd/app

# нагрузочное тестирование k6
http-k6:
	cat script.js | docker run --rm -i grafana/k6 run -
