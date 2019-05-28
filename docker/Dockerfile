FROM golang:1.12-stretch

LABEL maintainer = "Hadi Tajallaei"

RUN apt-get update -y && apt-get install -y \
    apt-utils \
    git \
    unzip

RUN mkdir -p /usr/local/golang-oop \
    && mkdir -p /usr/local/golang-oop/.git/hooks

COPY docker/git/.gitconfig /root
COPY docker/git/pre-commit /usr/local/golang-oop/.git/hooks

ENV GO111MODULE on

WORKDIR /usr/local/golang-oop