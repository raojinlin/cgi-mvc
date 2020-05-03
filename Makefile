build:
	go build -i main.go
	cp main docker/htdocs/main

docker_build:
	go build -i main.go
	cp -vap main docker/htdocs/main
	cd docker && docker build -t iogo-mvc . --no-cache
