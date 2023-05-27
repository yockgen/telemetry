# SPDX-FileCopyrightText: (C) 2023 Intel Corporation
# SPDX-License-Identifier: LicenseRef-Intel
VERSION 0.7

LOCALLY
ARG http_proxy=$(echo $http_proxy)
ARG https_proxy=$(echo $https_proxy)
ARG no_proxy=$(echo $no_proxy)
ARG HTTP_PROXY=$(echo $HTTP_PROXY)
ARG HTTPS_PROXY=$(echo $HTTPS_PROXY)
ARG NO_PROXY=$(echo $NO_PROXY)

FROM golang:1.20.2-alpine3.17
ENV http_proxy=$http_proxy
ENV https_proxy=$https_proxy
ENV no_proxy=$no_proxy
ENV HTTP_PROXY=$HTTP_PROXY
ENV HTTPS_PROXY=$HTTPS_PROXY
ENV NO_PROXY=$NO_PROXY

test:
    RUN go install github.com/magefile/mage@latest && \
        go install github.com/onsi/ginkgo/v2/ginkgo@latest
    WORKDIR /work
    COPY . .
    RUN mage -v test:go

lint:
    RUN go install github.com/magefile/mage@latest
    WORKDIR /work
    COPY . .
    RUN mage -v lint:openapi

build-api:
    ARG version='0.0.0-unknown'
    WORKDIR /work
    COPY . .
    RUN --ssh CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
        go build -trimpath -o build/exe \
            -ldflags "-s -w -extldflags '-static' -X main.Version=$version" \
            ./cmd/api
    SAVE ARTIFACT build/exe AS LOCAL ./build/exe

docker-api:
    FROM scratch
    ARG version='0.0.0-unknown'
    COPY (+build-api/exe --version=$version) .
    ENTRYPOINT ["/exe"]
    SAVE IMAGE edge-iaas-api:latest
