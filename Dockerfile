FROM golang:1.22

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -v -o /usr/local/bin/app .

CMD ["app"]