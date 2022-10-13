// Copyright 2022 coord_e
// Licensed under the Apache License, Streams 2.0 (the "License");
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

type reservesCntsExporter struct {
	ctx    context.Context
	client *epgstation.Client
	logger log.Logger

	normalReserves   *prometheus.Desc
	conflictReserves *prometheus.Desc
	skipReserves     *prometheus.Desc
	overlapReserves  *prometheus.Desc
}

// Verify if reservesCntsExporter implements prometheus.Collector
var _ prometheus.Collector = (*reservesCntsExporter)(nil)

func newReservesCntsExporter(ctx context.Context, client *epgstation.Client, logger log.Logger) *reservesCntsExporter {
	const subsystem = "reservescnts"

	return &reservesCntsExporter{
		ctx:    ctx,
		client: client,
		logger: logger,

		normalReserves: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "normal_reserves"),
			"Number of normal reserves in EPGStation.",
			nil, nil),
		conflictReserves: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "conflict_reserves"),
			"Number of conflicted reserves in EPGStation.",
			nil, nil),
		skipReserves: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "skip_reserves"),
			"Number of skipped reserves in EPGStation.",
			nil, nil),
		overlapReserves: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "overlap_reserves"),
			"Number of overlapped reserves in EPGStation.",
			nil, nil),
	}
}

func (e *reservesCntsExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.normalReserves
	ch <- e.conflictReserves
	ch <- e.skipReserves
	ch <- e.overlapReserves
}

func (e *reservesCntsExporter) Collect(ch chan<- prometheus.Metric) {
	cnts, err := e.client.GetReservesCnts(e.ctx)
	if err != nil {
		level.Error(e.logger).Log("msg", "failed to fetch EPGStation reserve counts", "err", err)
		return
	}

	ch <- prometheus.MustNewConstMetric(e.normalReserves, prometheus.GaugeValue, float64(cnts.Normal))
	ch <- prometheus.MustNewConstMetric(e.conflictReserves, prometheus.GaugeValue, float64(cnts.Conflicts))
	ch <- prometheus.MustNewConstMetric(e.skipReserves, prometheus.GaugeValue, float64(cnts.Skips))
	ch <- prometheus.MustNewConstMetric(e.overlapReserves, prometheus.GaugeValue, float64(cnts.Overlaps))
}
