export VERSION=1.0.0
build:
	GOOS=linux GOARCH=amd64 go build -o fusionweather -a -tags netgo -ldflags '-extldflags "-static"'
image: build
	cp fusionweather release/build/
	docker build -t fusionweather:$(VERSION) release/build