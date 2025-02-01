# Build the go helper
FROM golang:alpine AS build-helper
WORKDIR /helper
COPY ./ ./
RUN go mod download
RUN go build -o helper

# Build the bot
FROM registry.coregame.de/core/core-event-01:latest
WORKDIR /bot
COPY --from=build-helper /helper/helper ./helper

CMD ["./helper"]
