FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go get

RUN go build api/app.go

EXPOSE 8090

CMD ./app
