FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN make

FROM alpine:latest

COPY --from=builder /app/spurt .

ENTRYPOINT ["./spurt"]
