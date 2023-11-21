

docker-build:
	docker build -t defer-runtime .   

go-build:
	cd runtime && GOOS=linux GOARCH=amd64 go build        


EXEC_SHELL_CMD := docker exec -it $$(docker-compose ps -q  defer-runtime) /bin/bash

run: 
	docker-compose up --build -d
	$(EXEC_SHELL_CMD)
	

test:
	curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'
