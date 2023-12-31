FROM golang:1.20.6-alpine

RUN apk update && apk add git

WORKDIR /app

COPY . .

CMD ["tail", "-f", "/dev/null"]
