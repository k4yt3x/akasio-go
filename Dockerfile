# Name: Akasio Go Dockerfile
# Creator: K4YT3X
# Date Created: June 16, 2020
# Last Modified: January 25, 2022

# build akasio binary
FROM alpine:3.15 AS builder
COPY . /akasio-go
RUN apk add --no-cache go make \
    && cd /akasio-go \
    && make -j $(nproc)

# make final container
FROM alpine:3.15
LABEL maintainer="K4YT3X <i@k4yt3x.com>"
COPY --from=builder /akasio-go/bin/akasio /usr/local/bin/akasio

USER nobody:nogroup
WORKDIR /

ENTRYPOINT ["/usr/local/bin/akasio"]
CMD ["-b", "0.0.0.0:8000", "-r", "/etc/akasio.json"]
