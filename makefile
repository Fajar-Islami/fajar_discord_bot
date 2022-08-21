compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app
		
run:
	GOFLAGS=-mod=mod go run main.go

bot:
	./bin/gopher-bot-discord

heroku:
	git push heroku master

herokuconfig:
	cat .env | xargs heroku config:set