# Spurt
[![Go](https://github.com/zer-far/spurt/workflows/Go/badge.svg)](https://github.com/zer-far/spurt/actions?query=workflow%3A%22Go%22)

Spurt is a powerful testing tool designed to simulate high volumes of traffic and evaluate the performance and resilience of web servers.

## Building

```bash
git clone https://github.com/zer-far/spurt
cd spurt
make
```

## Usage

```bash
./spurt --url https://example.com
```
Press Ctrl+C to stop Spurt.

## Note

Please note that the primary purpose of Spurt is to assess the performance and robustness of web servers under controlled conditions. It is not intended for malicious use or to facilitate any form of Denial of Service (DoS) attack. Unauthorized or inappropriate usage of this tool is strictly prohibited and may have legal consequences.

Additionally, it's important to consider the potential impact on your internet connection while running Spurt. The tool generates a substantial amount of traffic, which may temporarily slow down your internet connection.
