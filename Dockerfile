FROM registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.18-openshift-4.11 AS builder
WORKDIR /go/src/github.com/rvanderp3/hairpin-test
COPY . .
ENV GO_PACKAGE github.com/rvanderp3/hairpin-test
RUN go mod tidy
RUN go mod vendor
RUN go build

