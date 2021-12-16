FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN apk --update add ca-certificates

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /go/bin/app

FROM scratch

ENV PATH=/bin

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /go/bin/app /go/bin/app

RUN apk add --no-cache ca-certificates openssl

ENTRYPOINT ["/go/bin/app"]