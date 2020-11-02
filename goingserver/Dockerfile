FROM golang:1.15.3-alpine3.12 AS build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./goingserver
RUN go install -v ./goingserver

FROM alpine:3.12.0
COPY --from=build /go/bin/goingserver /usr/local/bin/
CMD ["goingserver"]