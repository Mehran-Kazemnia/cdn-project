FROM golang:1.15.6-alpine3.12  AS build

RUN apk add git

RUN apk add gcc musl-dev

WORKDIR /tmp/go

COPY go.mod .

RUN  GOPROXY=https://goproxy.io/ go mod download

COPY . .

RUN go build -tags musl -o ./out/out.o .

FROM alpine:latest

COPY --from=build /tmp/go/out/out.o /out.o

CMD ["/out.o"]