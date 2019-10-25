# Build Go
FROM golang:alpine AS builder

RUN apk add --no-cache git
RUN apk add --no-cache ca-certificates
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true
# RUN /bin/sh -c "apk add --no-cache bash"
ENV GO_WORKDIR $GOPATH/src/github.com/iqbvl/login/
RUN mkdir -p src/github.com/iqbvl
# RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " > /root/.ssh/config
ADD . ${GO_WORKDIR}
WORKDIR ${GO_WORKDIR} 

RUN go get -u github.com/golang/dep/cmd/dep 
COPY ./Gopkg.toml $GOPATH/src/github.com/iqbvl/login/

RUN dep ensure -v
RUN go install 

# Minimize docker size
FROM alpine:latest
RUN set -eux; \
    apk add --no-cache ca-certificates
RUN apk add --update curl && \
    rm -rf /var/cache/apk/*
# RUN /bin/sh -c "apk add --no-cache bash"
COPY --from=builder /go/bin/login .
CMD [ "./login" ]
EXPOSE 8080 