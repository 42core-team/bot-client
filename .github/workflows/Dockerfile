ARG TAG_NAME="dev"
ARG GO_VERSION="1.22.3"

# Build the go helper
FROM golang:${GO_VERSION}-alpine AS build-helper
WORKDIR /helper
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o helper

# Build the bot
FROM ghcr.io/42core-team/my-core-bot:${TAG_NAME}
WORKDIR /bot
COPY --from=build-helper /helper/helper ./helper

CMD ["./helper"]
