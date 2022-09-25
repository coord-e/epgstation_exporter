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

package exporter

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/coord-e/epgstation_exporter/epgstation"
)

type versionExporter struct {
	ctx    context.Context
	client *epgstation.Client
	logger log.Logger

	info *prometheus.Desc
}

// Verify if versionExporter implements prometheus.Collector
var _ prometheus.Collector = (*versionExporter)(nil)

func newVersionExporter(ctx context.Context, client *epgstation.Client, logger log.Logger) *versionExporter {
	const subsystem = "version"

	return &versionExporter{
		ctx:    ctx,
		client: client,
		logger: logger,

		info: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "info"),
			"A metric with a constant '1' value labeled by metadata of EPGStation.",
			[]string{"version"}, nil),
	}
}

func (e *versionExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.info
}

func (e *versionExporter) Collect(ch chan<- prometheus.Metric) {
	version, err := e.client.GetVersion(e.ctx)
	if err != nil {
		level.Error(e.logger).Log("msg", "failed to fetch EPGStation version", "err", err)
		return
	}

	ch <- prometheus.MustNewConstMetric(e.info, prometheus.UntypedValue, 1.0, version.Version)
}
