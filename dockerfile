FROM golang:latest

WORKDIR /spotifyapidock

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
COPY ./cmd/statify-app-back/*.go ./


CMD [ "go", "run", "main.go"]