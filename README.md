# Akasio (Golang)

Akasio is a simple HTTP server that redirects traffic based on a JSON redirect table.

## Why Akasio

> What can this be used for?

Personally, I find sending long URLs like `https://gist.githubusercontent.com/k4yt3x/3b41a1a65f5d3087133e449793eb8858/raw` to people pretty annoying, since you'll either have to copy and paste the whole URL or type the whole URL out. URL shorteners like Akasio solve this issue. All that's needed to be done to send such a long URL is just to create a new mapping in the redirect table.

> What are Akasio's benefits compared to services like bit.ly?

It is self-hosted, and the redirect table is just a JSON file. This gives the users lots of flexibilities. The JSON file on the server can be symbolic-linked from a local workstation, updated by a front-end webpage, generated from a program, and so on.

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
Usages:
  -b string
        binding address (IP:port) (default "127.0.0.1:8080")
  -d    enable debugging mode, which disables security checks
  -n string
        server hostname (default "akas.io")
  -r string
        redirect table path (default "/etc/akasio.json")
```

## Building From Source

The following commands will build Akasio binary at `bin/akasio`.

```shell
git clone https://github.com/k4yt3x/akasio-go.git
cd akasio-go
make
```

After building, you may also use `make install` to install `akasio` to `/usr/local/bin/akasio`.
