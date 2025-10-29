FROM golang:1.23-alpine

WORKDIR /app

# Install required tools for migrations
RUN apk add --no-cache bash

# Install golang-migrate CLI
RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.19.6/migrate.linux-amd64.tar.gz \
    && tar -xvzf migrate.linux-amd64.tar.gz \
    && mv migrate /usr/local/bin/ \
    && rm migrate.linux-amd64.tar.gz

# Copy Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and migrations
COPY . .

# Build Go binary
RUN go build -o main ./cmd/main.go

# Apply migrations at container start
ENTRYPOINT ["/bin/sh", "-c", "migrate -path ./migrations -database $DATABASE_URL up && ./main"]

EXPOSE 8080