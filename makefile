compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app
		
run: exportconfig
	GOFLAGS=-mod=mod go run main.go

bot:
	./bin/go_discord_bot

heroku:
	git push heroku master

herokuconfig:
	cat .env.prod | xargs heroku config:set

exportconfig:
	export $(cat .env | xargs)

push:
	git push
	make heroku

log:
	heroku logs --tail