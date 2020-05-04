ACCESS_LOG="/usr/local/apache2/logs/cgi-mvc-access_log"
ERROR_LOG="/usr/local/apache2/logs/cgi-mvc-error_log"
IMAGE_NAME="iogo-mvc"
CONTAINER_NAME="iogo-mvc"

build:
	docker build -t "${IMAGE_NAME}" . --no-cache
run:
	go build ../main.go
	mv main htdocs/main
	docker run -d --name "${CONTAINER_NAME}" --rm -p 8081:8081\
 -v `pwd`/cigo-vhost.conf:/usr/local/apache2/conf/extra/cigo-vhost.conf\
 -v `pwd`/htdocs:/usr/local/apache2/htdocs/iogo "${IMAGE_NAME}"

clean:
	@if docker ps --format '{{.Names}}' | grep "${CONTAINER_NAME}" > /dev/null; \
		then docker stop "${CONTAINER_NAME}"; \
		else echo "Not found container '${CONTAINER_NAME}'"; fi
	@if docker images iogo-mvc | grep "${IMAGE_NAME}" > /dev/null; \
		then docker image rm "${IMAGE_NAME}" ; \
		else echo "Not found image '${IMAGE_NAME}'"; fi

restart:
	@docker exec "${CONTAINER_NAME}" apachectl -k restart

access_logs:
	@docker exec "${CONTAINER_NAME}" cat "${ACCESS_LOG}"

error_logs:
	@docker exec "${CONTAINER_NAME}" cat "${ERROR_LOG}"
