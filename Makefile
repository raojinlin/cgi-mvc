build:
	go build -i main.go
	cp main docker/htdocs/main

tests:
	@go test iogo/cgi
	@go test iogo/cgi/http
	@go test iogo/cgi/router
