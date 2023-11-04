FROM alpine:latest
WORKDIR /app
COPY . .
# RUN apt update && apt install go
RUN apk update
RUN apk add go
EXPOSE 3000
CMD go run .
