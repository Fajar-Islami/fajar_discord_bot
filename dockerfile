FROM golang:1.17 as build

WORKDIR /app
COPY . .

# Build the binary
RUN make compile

FROM alpine:latest


COPY --from=build /app/bin/app /app
COPY .env /.env

ENTRYPOINT [ "/app" ]