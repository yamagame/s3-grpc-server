FROM golang:1.20.3-alpine3.17

RUN apk add bash curl
RUN apk add --no-cache nodejs=18.14.2-r0 npm=9.1.2-r0
RUN npm install -g yarn
RUN go install github.com/ktr0731/evans@0.10.11
RUN apk add protobuf
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
