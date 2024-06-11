# Build the go helper
FROM golang:alpine AS build-helper
WORKDIR /helper
COPY ./ ./
RUN go mod download
RUN go build -o helper

# Build the bot
FROM --platform=$TARGETPLATFORM registry.coregame.de/core/core-event-0:latest
WORKDIR /bot
COPY --from=build-helper /helper/helper ./helper
COPY ./my-core-bot/Makefile ./Makefile

CMD ["./helper"]
