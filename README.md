# VMShare [![MIT License](https://img.shields.io/badge/License-MIT-a10b31)](https://github.com/notwithering/vmshare/blob/main/LICENSE)

**VMShare** is a program that just starts an http file server at the 1st argument. VMShare is meant to be used inbetween a host computer and a virtual environment without using shared folders.

VMShare acts a read-only shared folder over http.

![GIF](https://github.com/notwithering/vmshare/assets/124115470/a3cc59ed-6fb9-4ae0-b982-dd65daa720a1)

## Installing

```bash
go install github.com/notwithering/vmshare@latest
```

## Example Usage

```bash
$ echo "heyy kiddo" > test.txt
$ vmshare &
Sharing at http://127.0.0.1:8080

$ curl 127.0.0.1:8080/test.txt
heyy kiddo
```