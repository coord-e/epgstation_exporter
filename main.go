// Copyright 2022 coord_e
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  	 http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/promlog/flag"
	"github.com/prometheus/exporter-toolkit/web"
	webflag "github.com/prometheus/exporter-toolkit/web/kingpinflag"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/coord-e/epgstation_exporter/epgstation"
	"github.com/coord-e/epgstation_exporter/exporter"
)

var (
	// see Makefile
	BuildVersion   = "devel"
	BuildCommitSha = "unknown"
)

var (
	webConfig  = webflag.AddFlags(kingpin.CommandLine, ":9112")
	metricPath = kingpin.Flag("web.telemetry-path",
		"Path under which to expose metrics.").Default("/metrics").String()
	epgstationURL = kingpin.Flag("exporter.epgstation-url",
		"URL of the EPGStation instance.").Required().String()
	fetchVersion = kingpin.Flag("exporter.version",
		"Whether to export metrics from /api/version.").Default("true").Bool()
	fetchChannels = kingpin.Flag("exporter.channels",
		"Whether to export metrics from /api/channels.").Default("true").Bool()
	fetchSchedules = kingpin.Flag("exporter.schedules",
		"Whether to export metrics from /api/schedules.").Default("true").Bool()
	fetchStorages = kingpin.Flag("exporter.storages",
		"Whether to export metrics from /api/storages.").Default("true").Bool()
	fetchStreams = kingpin.Flag("exporter.streams",
		"Whether to export metrics from /api/streams.").Default("true").Bool()
	fetchEncode = kingpin.Flag("exporter.encode",
		"Whether to export metrics from /api/encode.").Default("true").Bool()
	fetchReservesCnts = kingpin.Flag("exporter.reserves-cnts",
		"Whether to export metrics from /api/reserves/cnts.").Default("true").Bool()
)

func main() {
	promlogConfig := &promlog.Config{}
	flag.AddFlags(kingpin.CommandLine, promlogConfig)
	kingpin.Version(BuildVersion)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	logger := promlog.New(promlogConfig)

	level.Info(logger).Log("msg", "Starting epgstation_exporter", "version", BuildVersion, "commit", BuildCommitSha)

	client, err := epgstation.NewClient(*epgstationURL)
	if err != nil {
		level.Error(logger).Log("msg", "failed to create EPGStation client", "err", err)
		os.Exit(1)
	}

	config := exporter.Config{
		FetchVersion:      *fetchVersion,
		FetchChannels:     *fetchChannels,
		FetchSchedules:    *fetchSchedules,
		FetchStorages:     *fetchStorages,
		FetchStreams:      *fetchStreams,
		FetchEncode:       *fetchEncode,
		FetchReservesCnts: *fetchReservesCnts,
	}
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		registry := prometheus.NewRegistry()
		exporter := exporter.New(r.Context(), client, config, logger)
		registry.MustRegister(exporter)
		h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
		h.ServeHTTP(w, r)
	}
	http.Handle(*metricPath, promhttp.InstrumentMetricHandler(prometheus.DefaultRegisterer, handler))

	level.Info(logger).Log("msg", "Listening on", "addresses", webConfig.WebListenAddresses)
	server := &http.Server{
		ReadHeaderTimeout: 5 * time.Second,
	}
	if err := web.ListenAndServe(server, webConfig, logger); err != nil {
		level.Error(logger).Log("err", err)
		os.Exit(1)
	}
}
