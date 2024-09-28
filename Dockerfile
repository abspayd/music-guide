# Base Stage
FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS base

ARG TARGETOS
ARG TARGETARCH

WORKDIR /usr/src/app

COPY . .

RUN apk update && \
	apk add npm && npm install
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go mod download && go mod verify

ENV CGO_ENABLED=0

# Production builder Stage
FROM base AS builder

RUN npm run tailwind:build
RUN templ generate
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o /usr/local/bin/app ./cmd/server

# Developer stage (live-reload)
FROM base AS dev

RUN go install github.com/air-verse/air@latest

CMD ["air"]

# Production stage (use binary on fresh alpine install)
FROM alpine:latest AS prod
COPY --from=builder /usr/local/bin/app .
EXPOSE 3000

CMD ["./app"]
