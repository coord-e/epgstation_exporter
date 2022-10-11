# syntax=docker.io/docker/dockerfile:1

FROM gcr.io/distroless/static-debian11:nonroot

ARG BIN_DIR
ARG TARGETARCH
COPY $BIN_DIR/$TARGETARCH/epgstation_exporter /usr/bin/epgstation_exporter

EXPOSE 9112
ENTRYPOINT ["/usr/bin/epgstation_exporter"]
