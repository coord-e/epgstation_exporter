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

type encodeExporter struct {
	ctx    context.Context
	client *epgstation.Client
	logger log.Logger

	runningEncodings *prometheus.Desc
	waitingEncodings *prometheus.Desc
}

// Verify if encodeExporter implements prometheus.Collector
var _ prometheus.Collector = (*encodeExporter)(nil)

func newEncodeExporter(ctx context.Context, client *epgstation.Client, logger log.Logger) *encodeExporter {
	const subsystem = "encode"

	return &encodeExporter{
		ctx:    ctx,
		client: client,
		logger: logger,

		runningEncodings: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "running_encodings"),
			"Number of running encodings in EPGStation.",
			nil, nil),
		waitingEncodings: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "waiting_encodings"),
			"Number of waiting encodings in EPGStation.",
			nil, nil),
	}
}

func (e *encodeExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.runningEncodings
	ch <- e.waitingEncodings
}

func (e *encodeExporter) Collect(ch chan<- prometheus.Metric) {
	encode, err := e.client.GetEncode(e.ctx, epgstation.GetEncodeOpts{
		IsHalfWidth: true,
	})
	if err != nil {
		level.Error(e.logger).Log("msg", "failed to fetch EPGStation encode", "err", err)
		return
	}

	ch <- prometheus.MustNewConstMetric(e.runningEncodings, prometheus.GaugeValue, float64(len(encode.RunningItems)))
	ch <- prometheus.MustNewConstMetric(e.waitingEncodings, prometheus.GaugeValue, float64(len(encode.WaitItems)))
}
