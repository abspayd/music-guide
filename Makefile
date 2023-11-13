PROJECT_PATH=.
EXE=bin/music
USER=abspayd
IMAGE_NAME=music-companion
VERSION=latest
PORT=3000

build:
	go build -o ${EXE} ${PROJECT_PATH}

run: build
	${EXE}

test:
	go test ./...

docker-build:
	docker build --tag ${USER}/${IMAGE_NAME}:${VERSION} .

docker-run:
	docker run -p ${PORT}:${PORT} -ti --rm --name ${IMAGE_NAME} ${USER}/${IMAGE_NAME}:${VERSION}

docker-clean:
	docker rm ${IMAGE_NAME}

docker-buildx:
	docker buildx create --driver=docker-container \
		--name=builder --bootstrap --use
	docker buildx build --platform linux/amd64,linux/arm64 --tag ${USER}/${IMAGE_NAME}:${VERSION} --push .
	docker buildx rm

deploy:
	doctl compute ssh music-companion --ssh-command \
		"docker pull ${USER}/${IMAGE_NAME}:${VERSION} && \
		docker stop ${IMAGE_NAME}; \
		docker run -itd -p ${PORT}:${PORT} --rm --name ${IMAGE_NAME} ${USER}/${IMAGE_NAME}:${VERSION}"

clean:
	rm ${EXE}
