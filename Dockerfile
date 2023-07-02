FROM busybox:1.36.0-uclibc as busybox

FROM gcr.io/distroless/base-debian10

# Add libs in order to execute healthcheck from docker-compose
COPY --from=busybox /bin/sh /bin/sh
COPY --from=busybox /bin/wget /bin/wget

ARG VERSION
ENV APPLICATION_VERSION=$VERSION

WORKDIR /

COPY bin/go-amqp-amd64-linux /go-amqp-amd64-linux

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/go-amqp-amd64-linux"]
