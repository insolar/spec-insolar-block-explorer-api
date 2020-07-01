FROM golang:1.14

COPY api-exported.yaml file.yaml

RUN go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen && \
    go get golang.org/x/tools/cmd/goimports
RUN mkdir /server && \
    oapi-codegen -package server -generate types,server ./file.yaml > /server/generated.go

RUN sed -i 's/SortByPulse_pulse_number_asc, jet_id_desc/SortByPulse_pulse_number_asc_jet_id_desc/g' /server/generated.go
RUN sed -i 's/SortByPulse_pulse_number_desc, jet_id_asc/SortByPulse_pulse_number_desc_jet_id_asc/g' /server/generated.go
RUN str=$(goimports /server/generated.go) && \
    echo "$str" > /server/generated.go
