FROM golang:1.22.3 AS build

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN make build
RUN strip /app/api-csm

FROM scratch AS base
LABEL name="api-csm" \
    version="0.1.0"
WORKDIR /

COPY --from=build /app/api-csm ./

CMD ["/api-csm"]