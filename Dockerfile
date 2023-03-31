FROM golang:alpine3.17 as build
LABEL stage=dockerbuilder

WORKDIR /app
COPY . .

# Build the binary
RUN go build -o fajar-discord-bot

FROM alpine:latest

RUN mkdir bin
COPY --from=build /app/fajar-discord-bot /app/bin/fajar-discord-bot

ENTRYPOINT [ "/app/bin/fajar-discord-bot" ]