FROM golang:alpine as builder

WORKDIR /app
ENV GO111MODULE=on

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /bookcrud

EXPOSE 8080

CMD [ "/bookcrud" ]