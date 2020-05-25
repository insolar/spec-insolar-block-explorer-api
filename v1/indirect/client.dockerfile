FROM insolar/specification-tools AS build

ENV PACKAGE=client
ENV OUTPUT_DIR=/package

COPY . ./src/

RUN npm run export -- --target=api-exported-gen.yaml --collapse \
     && npm run export -- --target=api-exported.yaml

RUN openapi-generator generate \
      --input-spec api-exported-gen.yaml \
      --generator-name go \
      --output ${OUTPUT_DIR} \
      --package-name ${PACKAGE} \
      --skip-validate-spec

RUN cp ./api-exported.yaml ${OUTPUT_DIR} && \
    cp ./api-exported-gen.yaml ${OUTPUT_DIR}


RUN mkdir /html && \
    npx redoc-cli bundle --output /html/index.html ./api-exported.yaml

FROM golang:1.14
ENV SRC_DIR=${GOPATH}/src/github.com/insolar/insolar-block-explorer-api

WORKDIR ${SRC_DIR}
COPY --from=build /package ${SRC_DIR}
COPY --from=build /html/ /html

RUN rm go.mod go.sum && \
    GO111MODULE=on go mod init && \
    GO111MODULE=on go mod tidy && \
    GO111MODULE=on go build && \
    rm -rf go.mod go.sum && \
    rm -rf .openapi-generator && \
    rm -rf api && \
    rm -rf docs && \
    for i in $(find . -type f ! -name '*.go' ! -name 'api-exported*.yaml'); do rm -rf "$i"; done;

RUN cp -R ${SRC_DIR} /spec

