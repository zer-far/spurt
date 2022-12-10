# GETBLAZE: HTTP GET flooder

getblaze can send more than 1k requests per second from a single machine.

# Building
```
git clone https://github.com/zer-far/getblaze
cd getblaze
go mod init getblaze
go mod tidy
go build
```

# Usage
```
# ./getblaze --hostname https://example.com
```
Press control+c to stop.

# Note
This program might slow your Internet connection down because it sends many requests and the target will send responses back.
