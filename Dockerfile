FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.* ./
COPY *.go ./
RUN go build -o /go-myip


FROM alpine

LABEL maintainer "quentinb69"

RUN apk update \
	&& apk upgrade --no-cache \
	&& apk add --no-cache curl

WORKDIR /opt/gmi
COPY --from=builder /go-myip ./gmi

RUN adduser -D gmi && chown -R gmi:gmi .

USER gmi:gmi

HEALTHCHECK --timeout=2s --start-period=5s \
	CMD curl -k https://localhost:8000/health

CMD [ "/opt/gmi/gmi" ]
