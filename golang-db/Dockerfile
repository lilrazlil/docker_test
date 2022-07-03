FROM golang:latest

ENV DB_PORT 8082

ARG ALL

ENV ALL $ALL

WORKDIR /go

COPY /docker-go-2 ./app

RUN cd app/ && go mod tidy

RUN cd /go/app/cmd/db/ && go build -o /go/app/go-db && mkdir /database

VOLUME [ "/database" ]

ENTRYPOINT [ "/go/app/go-db" ]  

EXPOSE 8082:8082