APP_NAME = "echo_shop"

default:
	go build -o ${APP_NAME}
	# env GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}

install:
	go mod download

c-test:
	env ECHO_SHOP_APP_RUNMODE=test go test -v ./test/controllers/...

m-test:
	env ECHO_SHOP_APP_RUNMODE=test go test -v ./test/models/...

dev:
  # go get -u github.com/Away0x/fresh
	env ECHO_SHOP_APP_RUNMODE=development fresh -c ./fresh.conf

mock:
	env ECHO_SHOP_APP_RUNMODE=development go run main.go mock

api-doc:
  # go get -u github.com/swaggo/swag/cmd/swag
	swag init

clean:
	if [ -f ${APP_NAME} ]; then rm ${APP_NAME}; fi

help:
	@echo "make - compile the source code"
	@echo "make install - install dep"
	@echo "make c-test - controllers test"
	@echo "make m-test - models test"
	@echo "make dev - run go fresh"
	@echo "make api-doc - generate swagger api docs"
	@echo "make mock - mock data"
	@echo "make clean - remove binary file"
