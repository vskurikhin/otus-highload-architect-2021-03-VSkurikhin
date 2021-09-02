build:
	go clean
	GOOS=linux GOARCH=amd64 go build
	docker build -t vskurikhin/app-main-amd64 .

all:	build
