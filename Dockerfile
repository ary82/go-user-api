FROM golang:1.22.2-alpine3.19 as base

RUN go install github.com/ary82/go-user-api/cmd/go-user-api@latest

FROM alpine:3.19

WORKDIR /usr/src

COPY --from=base /go/bin/go-user-api ./

EXPOSE 3000

CMD [ "./go-user-api" ]
