FROM golang:1.17 AS builder

ENV CGO_ENABLED 0
WORKDIR /app
COPY . /app

RUN go build -mod vendor -a -ldflags "-w -s" -installsuffix cgo -o ./bin/server ./cmd/server/main.go

FROM alpine:latest

WORKDIR /opt

COPY --from=builder /app/bin/server /opt/server

COPY ./entrypoint.sh /opt/entrypoint.sh

ENTRYPOINT ["/opt/entrypoint.sh"]

CMD ["/opt/server"]