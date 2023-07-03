.PHONY: build
build:
	CGO_ENABLED=0 go build -mod=vendor -ldflags "-s" -o ./bin/fss .

.PHONY: start
start:
	mkdir -p /tmp/127.0.0.1:8081 \
		/tmp/127.0.0.1:8082 \
		/tmp/127.0.0.1:8083 \
		/tmp/127.0.0.1:8084 \
		/tmp/127.0.0.1:8085 \
		/tmp/127.0.0.1:8086
	./bin/fss api --addr='127.0.0.1:8080' &\
	./bin/fss storage --name='127.0.0.1:8081' &\
	./bin/fss storage --name='127.0.0.1:8082' &\
	./bin/fss storage --name='127.0.0.1:8083' &\
	./bin/fss storage --name='127.0.0.1:8084' &\
	./bin/fss storage --name='127.0.0.1:8085' &\
	./bin/fss storage --name='127.0.0.1:8086'

.PHONY: stop
stop:
	ps | grep fss | awk '{print $$1;}' | xargs kill

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

%:
	@:
