FROM golang:alpine3.8
RUN apk update && apk add --upgrade git

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/platinum

FROM alpine:3.8
RUN apk --no-cache update

WORKDIR /
ENTRYPOINT ["/app/platinum"]
COPY --from=0 /app/bin/platinum /app/platinum