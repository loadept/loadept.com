# Building static files
FROM docker.io/node:22-alpine3.21 AS build-static

WORKDIR /app

RUN npm i -g pnpm

ENV API_URL=/
ENV CI=true

COPY web/package.json web/pnpm-lock.yaml ./

RUN pnpm i

COPY web .

RUN pnpm run build

# Build go server
FROM docker.io/golang:1.24.0-alpine3.21 AS build

ENV CGO_ENABLED=1

RUN apk add --no-cache gcc musl-dev git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o loadept.com cmd/loadept/main.go

# Execution stage
FROM docker.io/alpine:3.21

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=build /app/loadept.com ./
COPY --from=build-static /app/dist web/dist

ENTRYPOINT [ "./loadept.com" ]
