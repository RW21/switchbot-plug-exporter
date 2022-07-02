# FROM ubuntu:22.04 as builder

# # Avoid x509 errors
# RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

FROM alpine:3.16
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
RUN apk --no-cache add ca-certificates

COPY switchbot-plug-exporter /
ENTRYPOINT ["/switchbot-plug-exporter"]
