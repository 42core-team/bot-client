IMAGE_NAME = ghcr.io/42core-team/bot-client
TAG_NAME = dev
debug:
	go run main.go

run: build
	docker run -it --rm --env-file .env $(IMAGE_NAME):$(TAG_NAME)

push: build
	docker push $(IMAGE_NAME):$(TAG_NAME)

build:
	docker build -t $(IMAGE_NAME):$(TAG_NAME) --platform linux/amd64,linux/arm64 -f ./.github/workflows/Dockerfile .
