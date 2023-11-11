PATH=.
EXE=bin/music

build:
	go build -o ${EXE} ${PATH}

run: build
	${EXE}

test:
	go test ./...

clean:
	rm ${EXE}
