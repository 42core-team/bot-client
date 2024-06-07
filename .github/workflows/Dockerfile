# Build the go helper
FROM go:latest AS build-helper
WORKDIR /helper
COPY ./ ./
RUN go mod download
RUN go build -o helper

# Build the alpine gcc image
FROM alpine:latest AS alpine-gcc
RUN apk --no-cache add \
	gcc musl-dev \
	make

# Build the library
FROM alpine-gcc AS build-library
COPY ./connection /connection
RUN make -C connection

# Create the release image
FROM alpine-gcc AS release
WORKDIR /bot
COPY ./my-core-bot

COPY --from=build-helper /helper/helper ./helper
COPY --from=build-library /connection/con_lib.a /my-core-bot/connection/con_lib.a

CMD ["/bot/helper"]
