build:
	GOFLAGS=-mod=mod go build -o bin/gopher-bot-discord main.go 
	
run:
	GOFLAGS=-mod=mod go run main.go

bot:
	./bin/gopher-bot-discord -t $BOT_TOKEN