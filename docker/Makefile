build:	
	docker image rm iogo-mvc
	docker build -t iogo-mvc . --no-cache
run:
	go build ../main.go
	mv main htdocs/main
	docker run -d --name iogo-mvc --rm -p 8081:8081 -v `pwd`/htdocs:/usr/local/apache2/htdocs/iogo iogo-mvc
