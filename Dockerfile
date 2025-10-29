FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target="/root/.cache/go-build" --mount=type=cache,target="/go/pkg/mod" go build -a -o service ./cmd/service

COPY . .

RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]
