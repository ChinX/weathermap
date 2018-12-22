export VERSION=1.0.0
build:
	GOOS=linux GOARCH=amd64 go build -o weather -a -tags netgo -ldflags '-extldflags "-static"'
image: build
	cp weather release/build/
	docker build -t weather:$(VERSION) release/build