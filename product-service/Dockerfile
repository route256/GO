# 1
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY secret/ ./secret/
COPY internal/ ./internal/
COPY cmd/ ./cmd/

RUN CGO_ENABLED=0 go build -a -o ./products-service ./cmd/products-server/

# 2

FROM scratch

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY init/data.json ./init/
COPY init/skus.json ./init/
COPY secret/authorized.txt ./secret/authorized.txt

ENV TZ=Europe/Moscow
ENV PORT=8082
ENV HOST=0.0.0.0
EXPOSE 8082
COPY --from=builder /app/products-service /app/products-service

LABEL domain=route256 service=products

CMD ["./products-service"]
