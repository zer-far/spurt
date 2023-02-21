# Getblaze

Getblaze is a [DoS (Denial of Service)](https://en.wikipedia.org/wiki/Denial-of-service_attack) tool that can send over 1k GET requests per second and overload web servers.

## Building

```bash
  git clone https://github.com/zer-far/getblaze
  cd getblaze
  make
```

## Usage

```bash
./getblaze --hostname https://example.com
```
Press control+c to stop.

## Note

This program might slow your Internet connection down because it sends many requests and the target will send responses back.
