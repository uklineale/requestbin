FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git g++
WORKDIR $GOPATH/src/requestbin/
COPY . .

RUN go get -d -v
RUN go test && go build -o /go/bin/requestbin

######### Build phase 2 ###########

FROM scratch

COPY --from=builder /go/bin/requestbin /go/bin/requestbin

ENTRYPOINT ["/go/bin/requestbin"]