FROM golang:1.18-alpine AS builder

RUN apk add --update --no-cache \
  ca-certificates tzdata openssh git mercurial && update-ca-certificates \
  && rm -rf /var/cache/apk/*

WORKDIR /src

COPY go.mod* go.sum* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go install ./cmd/...

FROM alpine

RUN adduser -S -D -H -h /app appuser
USER appuser

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/* /bin/

ENV PORT=8080
EXPOSE $PORT

CMD ["octostack-api"]
