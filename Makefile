IMAGE_NAME = bot-client

run: build
	docker run -it --rm --env-file .env $(IMAGE_NAME):latest

build:
	docker build -t $(IMAGE_NAME):latest -f ./.github/workflows/Dockerfile .
