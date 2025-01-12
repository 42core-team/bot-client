IMAGE_NAME = registry.coregame.de/core/bot-client

debug:
	go run main.go

run: build
	docker run -it --rm --env-file .env $(IMAGE_NAME):latest

push: build
	docker push $(IMAGE_NAME):latest

build:
	docker build -t $(IMAGE_NAME):latest -f ./.github/workflows/Dockerfile .
