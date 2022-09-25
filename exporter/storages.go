// Copyright 2022 coord_e
// Licensed under the Apache License, Storages 2.0 (the "License");
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

package exporter

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/coord-e/epgstation_exporter/epgstation"
)

type storagesExporter struct {
	ctx    context.Context
	client *epgstation.Client
	logger log.Logger

	used      *prometheus.Desc
	available *prometheus.Desc
	total     *prometheus.Desc
}

// Verify if storagesExporter implements prometheus.Collector
var _ prometheus.Collector = (*storagesExporter)(nil)

func newStoragesExporter(ctx context.Context, client *epgstation.Client, logger log.Logger) *storagesExporter {
	const subsystem = "storages"

	return &storagesExporter{
		ctx:    ctx,
		client: client,
		logger: logger,

		used: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "used_bytes"),
			"Used storage size of EPGStation in bytes.",
			[]string{"name"}, nil),
		available: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "available_bytes"),
			"Available storage size of EPGStation in bytes.",
			[]string{"name"}, nil),
		total: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_bytes"),
			"Total storage size of EPGStation in bytes.",
			[]string{"name"}, nil),
	}
}

func (e *storagesExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.used
	ch <- e.available
	ch <- e.total
}

func (e *storagesExporter) Collect(ch chan<- prometheus.Metric) {
	storages, err := e.client.GetStorages(e.ctx)
	if err != nil {
		level.Error(e.logger).Log("msg", "failed to fetch EPGStation storages", "err", err)
		return
	}

	for _, storage := range storages.Items {
		ch <- prometheus.MustNewConstMetric(e.used, prometheus.GaugeValue, float64(storage.Used), storage.Name)
		ch <- prometheus.MustNewConstMetric(e.available, prometheus.GaugeValue, float64(storage.Available), storage.Name)
		ch <- prometheus.MustNewConstMetric(e.total, prometheus.GaugeValue, float64(storage.Total), storage.Name)
	}
}
