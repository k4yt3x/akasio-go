# Akasio (Golang)

## Description

Akasio is a simple HTTP server that redirects traffic based on a JSON redirect table.

If you own a domain and wish to self-host a URL-shortening service, then this is the right tool for you.

Originally, Akasio is a backend server for website [akas.io](akas.io) (akas stands for "also known as") written in Python and Flask. After rewriting the server with Golang, I decided to open-source it for anyone else that might be interested by it. It is both very easy to deploy and very easy to maintain, thanks to Golang's handy single-binary release (let's not talk about its size here).

## Why Akasio

> What can this be used for?

Personally, I find sending long URLs like `https://gist.githubusercontent.com/k4yt3x/3b41a1a65f5d3087133e449793eb8858/raw` to people pretty annoying, since you'll either have to copy and paste the whole URL or type the whole URL out. URL shorteners like Akasio solve this issue. All that's needed to be done to send such a long URL is just to create a new mapping in the redirect table (akas.io/z).

> What are Akasio's benefits compared to services like bit.ly?

Akasio is self-hosted, and the redirect table is just a JSON file. This gives the users lots of flexibilities. The JSON file on the server can be symbolic-linked from a local workstation, updated by a front-end webpage, generated from a program, and so on.

## Usages

This section covers Akasio's fundamental concepts, basic usages and setup guide.

### Redirect Table

Akasio redirects incoming requests based on what's called a "redirect table". This table is essentially a JSON file with a simple source-to-target mapping. You can find an example `akasio.json` under the `configs` directory. By default, Akasio reads the redirect table from `/etc/akasio.json`.

```json
{
    "/": "http://k4yt3x.com/akasio-go/",
    "/g": "https://www.google.com",
    "/k4yt3x": "https://k4yt3x.com"
}
```

This example redirect table does the following mappings:

- `/` to http://k4yt3x.com/akasio-go/
- `/g` to https://www.google.com
- `/k4yt3x` to https://k4yt3x.com

Taking the `/g` mapping for example, when a user visits `https://yourwebsite.com/g`, the user will be redirected to https://www.google.com via a HTTP 301 (moved permanently) response.

### Website Setup

The recommended setup is to start Akasio as a service behind reverse proxy web server like Apache, Nginx or Caddy. You can find an example service file at `configs/akasio.service`.

A typical stand-alone setup process will look like the following.

1. Build the `akasio` binary or download the `akasio` binary from [releases](https://github.com/k4yt3x/akasio-go/releases).
1. Move the `akasio` binary to `/usr/local/bin/akasio`.
1. Move the service file to `/etc/systemd/system/akasio.service`.
1. Reload systemd with `systemctl daemon-reload`.
1. Enable and start the service with `systemctl enable --now akasio`.
1. Verify that the service has been started successfully via `curl -v 127.0.0.1:8000`.
1. Configure front-end web server to reverse proxy to http://127.0.0.1:8000.

### Binary Usages

The binary's usage is as following. You can also invoke `akasio -h` to see the usages.

```console
Usage:
  -b string
        binding address (IP:port) (default "127.0.0.1:8000")
  -d    enable debugging mode, which disables security checks
  -n value
        server hostname, can be specified multiple times
  -r string
        redirect table path (default "/etc/akasio.json")
  -v    print Akasio version and exit
```

The command below, for instance, launches Akasio, reads configurations from the file `/etc/akasio.json`, and serves domains `akas.io` and `ffg.gg`.

```shell
/usr/local/bin/akasio -r /etc/akasio.json -n akas.io -n ffg.gg
```

### Running from Docker

Akasio is also available on Docker Hub. Below is an example how you can run Akasio with Docker. Be sure to create the redirect table and change the redirect table's path in the command below. You'll also need to change the server's hostname.

```shell
docker run -it -p 8000:8000 -v $PWD/akasio.json:/etc/akasio.json -h akasio --name akasio k4yt3x/akasio-go:1.1.1 -n akas.io

docker run -it \                                            # interactive
           -p 8000:8000 \                                   # bind container port to host's port 8000
           -v $PWD/akasio.json:/etc/akasio.json \           # bind mount host's akasio.json file under the current directory to container's /etc/akasio.json
           -h akasio \                                      # set container hostname akasio
           --name akasio \                                  # set container name akasio
           k4yt3x/akasio-go:1.1.1 \                         # container name
           -n akas.io                                       # listening hostnames
```

After spinning the container up, you can verify that it's running correctly by making a query with `curl` or any other tool of your preference.

## Building From Source

The following commands will build Akasio binary at `bin/akasio`.

```shell
git clone https://github.com/k4yt3x/akasio-go.git
cd akasio-go
make
```

After building, you may also use `sudo make install` to install `akasio` to `/usr/local/bin/akasio`.
