FROM ubuntu:22.04 as builder

# Avoid x509 errors
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY switchbot-plug-exporter /
ENTRYPOINT ["/switchbot-plug-exporter"]
