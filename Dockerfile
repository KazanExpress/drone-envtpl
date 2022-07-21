FROM golang:1.13.4-alpine as builder

RUN apk add --no-cache ca-certificates curl git

WORKDIR /app
COPY . .
WORKDIR /app/cmd


RUN CGO_ENABLED=0 GOOS=`go env GOHOSTOS` GOARCH=`go env GOHOSTARCH` go build -o drone-envtpl

FROM python:3.8-alpine
RUN apk add --no-cache \
		git \
        openssh-client \
        ca-certificates \
        bash \
        curl perl && \
    rm -rf /var/cache/apk/*

RUN pip install envtpl

COPY --from=builder /app/cmd/drone-envtpl /bin
ENTRYPOINT ["/bin/drone-envtpl"]
