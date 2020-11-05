# Name: Akasio Dockerfile
# Creator: K4YT3X
# Date Created: June 16, 2020
# Last Modified: November 5, 2020

FROM alpine:latest

LABEL maintainer="K4YT3X <k4yt3x@k4yt3x.com>"

COPY . akasio-go
RUN apk add --no-cache --virtual .build-deps go make \
    && cd akasio-go \
    && make -j $(nproc) \
    && make install \
    && rm -rf /akasio-go \
    && apk del .build-deps

# run the Akasio binary with user nobody and group nogroup by default
USER nobody:nogroup

WORKDIR /
ENTRYPOINT ["/usr/local/bin/akasio", "-b", "0.0.0.0:8000", "-r", "/etc/akasio.json"]
