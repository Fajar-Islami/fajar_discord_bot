FROM golang:alpine3.17 as build
LABEL stage=dockerbuilder

WORKDIR /app
COPY . .

# Build the binary
RUN go build -o fajar-discord-bot

FROM alpine:latest

COPY --from=build /app/fajar-discord-bot /app/fajar-discord-bot

ENTRYPOINT [ "/app/fajar-discord-bot" ]