# GETBLAZE: HTTP GET flooder

This program can send more than 1,000 requests per second from a single machine.

# Building
```
git clone https://github.com/zer-far/getblaze
cd getblaze
go build
```

# Usage
```
# ./getblaze --hostname https://example.com
```
Press control+c to stop.

# Note
This program might slow your Internet connection down because it sends many requests and the target will send responses back.
