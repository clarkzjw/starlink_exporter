FROM golang:1.24.0 AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum vendor ./
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 go build -v -o /usr/local/bin/exporter ./cmd/starlink_exporter/main.go

FROM alpine:3.21.3

COPY --from=builder /usr/local/bin/exporter /usr/bin/exporter

CMD ["/usr/bin/exporter"]
