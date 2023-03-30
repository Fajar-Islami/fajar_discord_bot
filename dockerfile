FROM golang:alpine3.17 as build
LABEL stage=dockerbuilder

WORKDIR /app
COPY . .

# Build the binary
RUN go build -o apps

FROM alpine:latest

COPY --from=build /app/apps /app/apps

ENTRYPOINT [ "/app" ]