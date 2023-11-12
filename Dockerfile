FROM golang:1.21
WORKDIR /app
COPY . .
EXPOSE 3000
RUN go build -o bin/music .
CMD ./bin/music
