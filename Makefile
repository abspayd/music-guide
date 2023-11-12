PROJECT_PATH=.
EXE=bin/music
USER=abspayd
IMAGE_NAME=music-companion
VERSION=1.0.1
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

docker-export:
	docker buildx build --platform linux/amd64 --tag ${USER}/${IMAGE_NAME}:${VERSION} .
	docker push ${USER}/${IMAGE_NAME}:${VERSION}

clean:
	rm ${EXE}
