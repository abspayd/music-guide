FROM golang:1.21
WORKDIR /app
COPY . .
EXPOSE 3000
CMD go run .
