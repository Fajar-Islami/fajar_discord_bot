registry:=ghcr.io
username:=fajar-islami
image:=fajar_discord_bot
tags:=latest

compile:
	go build -o bin/app
		
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

dockerbuild:
	docker build --rm -t ${registry}/${username}/${image}:${tags} .
	docker image prune --filter label=stage=dockerbuilder -f

dockerun:
	docker run --name ${image} ${registry}/${username}/${image}:${tags}

dockerrm:
	docker rm ${registry}/${username}/${image}:${tags} -f
	docker rmi ${registry}/${username}/${image}:${tags}

dockerup: ## up compose image
	docker compose -f docker-compose-app.yaml up -d

dockerlogs: ## logs compose image
	docker compose -f docker-compose-app.yaml logs -f

dockerstop: ## stop compose image
	docker compose -f docker-compose-app.yaml stop

dockerdown: ## rm compose image
	docker compose -f docker-compose-app.yaml down -v

dockerrm:
	docker rm ${registry}/${username}/${image}:${tags} -f
	docker rmi ${registry}/${username}/${image}:${tags}