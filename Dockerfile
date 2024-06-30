FROM golang:1.22.4-alpine as build
LABEL stage=dockerbuilder

WORKDIR /app
COPY . .

# Build the binary
RUN go build -o fajar-discord-bot

FROM alpine:latest

COPY --from=build /app/fajar-discord-bot /app/fajar-discord-bot

ENTRYPOINT [ "/app/fajar-discord-bot" ]