export VERSION=1.0.0
build:
	GOOS=linux GOARCH=amd64 go build -o forecast -a -tags netgo -ldflags '-extldflags "-static"'
image: build
	cp forecast release/build/
	docker build -t forecast:$(VERSION) release/build