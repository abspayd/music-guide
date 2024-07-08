FROM golang:1.22
WORKDIR /app
COPY . .
EXPOSE 3000
RUN go build -o bin/music .
CMD ./bin/music
