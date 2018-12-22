export VERSION=1.0.0
image:
	cp -r ./app.js ./bin ./package.json ./public ./routes ./shutdown.sh ./start.sh ./views release/build
	docker build -t weathermapweb:$(VERSION) release/build