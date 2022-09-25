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

type streamsExporter struct {
	ctx    context.Context
	client *epgstation.Client
	logger log.Logger

	streams *prometheus.Desc
}

// Verify if streamsExporter implements prometheus.Collector
var _ prometheus.Collector = (*streamsExporter)(nil)

func newStreamsExporter(ctx context.Context, client *epgstation.Client, logger log.Logger) *streamsExporter {
	const subsystem = "streams"

	return &streamsExporter{
		ctx:    ctx,
		client: client,
		logger: logger,

		streams: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "streams"),
			"Number of streams in EPGStation.",
			[]string{"type"}, nil),
	}
}

func (e *streamsExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.streams
}

func (e *streamsExporter) Collect(ch chan<- prometheus.Metric) {
	streams, err := e.client.GetStreams(e.ctx, epgstation.GetStreamsOpts{
		IsHalfWidth: true,
	})
	if err != nil {
		level.Error(e.logger).Log("msg", "failed to fetch EPGStation streams", "err", err)
		return
	}

	counts := map[string]int{
		"LiveStream":     0,
		"LiveHLS":        0,
		"RecordedStream": 0,
		"RecordedHLS":    0,
	}
	for _, stream := range streams.Items {
		counts[stream.Type]++
	}

	for ty, count := range counts {
		ch <- prometheus.MustNewConstMetric(e.streams, prometheus.GaugeValue, float64(count), ty)
	}
}
