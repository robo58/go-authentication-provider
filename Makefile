get-docs:
	go get -u github.com/swaggo/swag/cmd/swag

docs: get-docs
	swag init --dir cmd/api --parseDependency --output docs

build:
	go build -o bin/restapi main.go

run:
	REDIRECT_URL="http://localhost:3000/oauth/client/callback" CLIENT_ID="schoolinfoclient" CLIENT_SECRET="schoolinfosecret" go run main.go

test:
	go test -v ./test/...

build-docker: build
	docker build . -t api-rest

run-docker: build-docker
	docker run -p 3000:3000 api-rest

seed-db:
	go run ./cmd/seed.go
