ARG TAG_NAME="dev"

# Build the go helper
FROM golang:alpine AS build-helper
WORKDIR /helper
COPY ./ ./
RUN go mod download
RUN go build -o helper

# Build the bot
FROM ghcr.io/42core-team/my-core-bot:${TAG_NAME}
WORKDIR /bot
COPY --from=build-helper /helper/helper ./helper

CMD ["./helper"]
