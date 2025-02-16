FROM golang:1.24.0-alpine3.21 AS build

ENV CGO_ENABLED=1

RUN apk add --no-cache gcc musl-dev git

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -v -x -o loadept.com cmd/loadept/main.go

# RUN STAGE
FROM alpine:3.21

WORKDIR /app

COPY --from=build /app/loadept.com ./
COPY --from=build /app/web web

ENTRYPOINT [ "./loadept.com" ]
