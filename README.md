# EPGStation Exporter

[![ci](https://github.com/coord-e/epgstation_exporter/actions/workflows/ci.yml/badge.svg)](https://github.com/coord-e/epgstation_exporter/actions/workflows/ci.yml)

Prometheus exporter for [EPGStation](https://github.com/l3tnun/EPGStation/) metrics.
Pre-built binaries are available at [the releases](https://github.com/coord-e/epgstation_exporter/releases).
Container images are available at [the packages](https://github.com/coord-e?tab=packages&repo_name=epgstation_exporter).

## Usage

```console
$ epgstation_exporter -h
usage: epgstation_exporter --exporter.epgstation-url=EXPORTER.EPGSTATION-URL [<flags>]

Flags:
  -h, --help                    Show context-sensitive help (also try --help-long and --help-man).
      --web.systemd-socket      Use systemd socket activation listeners instead of port listeners
                                (Linux only).
      --web.listen-address=:9112 ...
                                Addresses on which to expose metrics and web interface. Repeatable
                                for multiple addresses.
      --web.config.file=""      [EXPERIMENTAL] Path to configuration file that can enable TLS or
                                authentication.
      --web.telemetry-path="/metrics"
                                Path under which to expose metrics.
      --exporter.epgstation-url=EXPORTER.EPGSTATION-URL
                                URL of the EPGStation instance.
      --exporter.version        Whether to export metrics from /api/version.
      --exporter.channels       Whether to export metrics from /api/channels.
      --exporter.schedules      Whether to export metrics from /api/schedules.
      --exporter.storages       Whether to export metrics from /api/storages.
      --exporter.streams        Whether to export metrics from /api/streams.
      --exporter.encode         Whether to export metrics from /api/encode.
      --exporter.reserves-cnts  Whether to export metrics from /api/reserves/cnts.
      --log.level=info          Only log messages with the given severity or above. One of: [debug, info, warn, error]
      --log.format=logfmt       Output format of log messages. One of: [logfmt, json]
      --version                 Show application version.
```

### Example

To run against a EPGStation instance running at `localhost:8888`:

```console
$ epgstation_exporter --exporter.epgstation-url=http://localhost:8888/
```

## Build

```console
$ make build
```
