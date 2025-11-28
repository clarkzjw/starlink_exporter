FROM docker.io/library/golang:1 AS builder

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go vet -v cmd/starlink_exporter/main.go
RUN go test -v cmd/starlink_exporter/main.go

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags='-s -w' -o starlink_exporter cmd/starlink_exporter/main.go

FROM gcr.io/distroless/static:nonroot AS final
COPY --from=builder /go/src/app/starlink_exporter /
ENTRYPOINT ["/starlink_exporter"]
