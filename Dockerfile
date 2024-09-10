# Stage 1: Build the Go application
FROM golang:1.22.5-alpine AS builder
RUN apk add --no-cache git

RUN git help

COPY repo/go.mod /repo/go.mod
COPY repo/go.sum /repo/go.sum
WORKDIR /repo
RUN go mod download

COPY repo /repo

RUN go test -c -o test_bin

# Stage 2: Create the final image
FROM alpine:3.18.3

WORKDIR /repo
COPY --from=builder /repo/test_bin /repo/test_bin
COPY --from=builder /repo/data /repo/data

ENV DATA_PATH=/repo/data

CMD ["/repo/test_bin"]
