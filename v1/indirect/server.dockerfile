FROM golang:1.14

COPY api-exported.yaml file.yaml

RUN go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen && ls -al .
RUN mkdir /server && \
    oapi-codegen -package server -generate types,server ./file.yaml > /server/generated.go
