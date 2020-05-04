build:	
	docker build -t iogo-mvc . --no-cache
run:
	go build ../main.go
	mv main htdocs/main
	docker run -d --name iogo-mvc --rm -p 8081:8081\
 -v `pwd`/cigo-vhost.conf:/usr/local/apache2/conf/extra/cigo-vhost.conf\
 -v `pwd`/htdocs:/usr/local/apache2/htdocs/iogo iogo-mvc

clean:
	@if docker ps --format '{{.Names}}' | grep iogo-mvc > /dev/null; \
		then docker image rm iogo-mvc; \
		else echo "Not found container 'iogo-mvc'"; fi
	@if docker images iogo-mvc | grep iogo-mvc > /dev/null; \
		then docker image rm iogo-mvc ; \
		else echo "Not found image 'iogo-mvc'"; fi

restart:
	@docker exec -it iogo-mvc apachectl -k restart