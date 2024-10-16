FROM golang:1.22.4-alpine AS build
LABEL stage=dockerbuilder
WORKDIR /app
# Copy go.mod and go.sum first to take advantage of Docker caching.
COPY go.mod go.sum ./

# Download dependencies.
RUN go mod download

# copy code
COPY . .

# Build the binary
RUN go build -o fajar-discord-bot

FROM alpine:latest

COPY --from=build /app/fajar-discord-bot /fajar-discord-bot

ENTRYPOINT [ "/fajar-discord-bot" ]