NAME=zb
ENTRYPOINT=./src/zerobuild/main.go


run:
	go run ${ENTRYPOINT}

build:
	go build -o ${NAME} ${ENTRYPOINT}

.PHONY: run
