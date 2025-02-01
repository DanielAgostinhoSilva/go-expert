STACK_NAME=go-expert-apis

start:
	docker compose -p ${STACK_NAME} up -d

build:
	docker compose -p ${STACK_NAME} up --build -d

stop:
	docker compose -p ${STACK_NAME} stop

restart: stop start

clean:
	docker compose -p ${STACK_NAME} down

logs:
	docker compose -p ${STACK_NAME} logs -f


