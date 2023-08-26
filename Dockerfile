FROM golang:1.20.3-alpine3.17

# COPY self-signed.cert /usr/local/share/ca-certificates/self-signed.cert
# RUN update-ca-certificates
RUN apk add bash curl
RUN apk add --no-cache nodejs npm
RUN npm install -g yarn
RUN go install github.com/ktr0731/evans@0.10.11
RUN apk add protobuf
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
