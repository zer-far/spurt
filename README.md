# Spurt
[![Go](https://github.com/zer-far/spurt/workflows/Go/badge.svg)](https://github.com/zer-far/spurt/actions?query=workflow%3A%22Go%22)

## Introduction

Spurt is a powerful stress testing tool designed to simulate high volumes of traffic and evaluate the performance and resilience of web servers. It has been tested on Linux.

Spurt first checks if the target URL is valid. A custom module is used to generate common "User-Agent" and "Referer" headers quickly, allowing it to make thousands of requests per second. Another module is used for multi-threading to send requests in parallel.

## Compile

### Prerequisites

Make sure you have the following installed:
- Go
- Make
- Git

### Steps

```bash
git clone https://github.com/zer-far/spurt
cd spurt
make
```

## Basic usage

```bash
./spurt --url https://example.com
```

Press Ctrl+C to stop Spurt.

## Features

### Target URL

- Option: --url string

- Description: Sets the target URL for requests.

### Check IP address

- Option: --check

- Description: Checks your public IP address before using it for testing.

### Custom cookie

- Option: --cookie string

- Description: Uses a custom cookie.

### Sleep time between requests

- Option: --sleep int

- Description: Sets a delay (in ms) between requests to avoid rate limits.

- Default: 1 ms

### Multi-threading

- Option: --threads int
  
- Description: Specifies the number of threads for sending requests.
  
- Default: 1 thread

### Request timeout

- Option: --timeout int
  
- Description: Sets a timeout (in ms) for each request to handle slow responses.
  
- Default: 3000 ms

### Example command

```bash
./spurt --check --cookie "sessionid=123456" --sleep 10 --threads 2 --timeout 4000 --url https://example.com
```

This command uses the IP check, a custom cookie, a sleep time of 10 ms, 2 threads and a timeout ot 4000 ms.

## Notes

- Spurt is intended for testing server performance under controlled conditions. It should not be used maliciously or for Denial of Service (DoS) attacks. Unauthorised use may lead to legal consequences.
-   Your internet connection may slow down during use.
-   Excessive requests may result in rate limiting or blocking.
-   Avoid using on connections with limited bandwidth.
