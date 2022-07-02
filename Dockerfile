FROM scratch
COPY switchbot-plug-exporter /
ENTRYPOINT ["/switchbot-plug-exporter"]