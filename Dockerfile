# syntax=docker.io/docker/dockerfile:1

ARG BASE_IMAGE
FROM $BASE_IMAGE

ARG BIN_DIR
ARG TARGETARCH
COPY $BIN_DIR/$TARGETARCH/epgstation_exporter /usr/bin/epgstation_exporter

EXPOSE 9112
CMD ["/usr/bin/epgstation_exporter"]
