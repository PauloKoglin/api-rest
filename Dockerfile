FROM golang:1.8

LABEL Maintender="pk.koglin@gmail.com"
WORKDIR /go/src/api-rest

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
#RUN go run main.go#

CMD [ "api-rest" ]