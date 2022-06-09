get: $(repo)
	docker run --rm -v "$(PWD)":/usr/src/myapp -w /usr/src/myapp -p 8080:8080 golang:1.18 /bin/bash -c "go get -u $(repo)"
install: $(repo)
	docker run --rm -v "$(PWD)":/usr/src/myapp -w /usr/src/myapp -p 8080:8080 golang:1.18 /bin/bash -c "go install $(repo)"
run:
	docker run --rm -v "$(PWD)":/usr/src/myapp -w /usr/src/myapp -p 8080:8080 golang:1.18 /bin/bash -c "go run ."
swagger:
	docker run --rm -v "$(PWD)":/usr/src/myapp -w /usr/src/myapp -p 8080:8080 golang:1.18 /bin/bash -c "go install github.com/swaggo/swag/cmd/swag; $$GOPATH/bin/swag init"