FROM golang:1.15.3-alpine3.12 AS build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./httpserver
RUN go install -v ./httpserver

FROM alpine:3.12.0
COPY --from=build /go/bin/httpserver /usr/local/bin/
CMD ["httpserver"]
