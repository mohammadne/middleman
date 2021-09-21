FROM golang:1.17 AS builder

WORKDIR /app

COPY . .

RUN go mod download && make code-gen

RUN go build -o /bin/app ./cmd/root.go

FROM alpine:latest

RUN apk add --no-cache libc6-compat 

WORKDIR /bin/

COPY --from=builder /bin/app .

LABEL org.opencontainers.image.source="https://github.com/mohammadne/middleman"

ENTRYPOINT ["/bin/app"]

CMD ["server", "--env=prod"]
