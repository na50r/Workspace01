FROM golang:1.22.5

COPY repo/go.mod /repo/go.mod
COPY repo/go.sum /repo/go.sum
WORKDIR /repo
RUN go mod download

COPY repo /repo

ENV DATA_PATH=/repo/data

CMD ["sh", "-c", "sleep infinity"]
