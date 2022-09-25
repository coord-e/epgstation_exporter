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

type channelsExporter struct {
	ctx    context.Context
	client *epgstation.Client
	logger log.Logger

	channels *prometheus.Desc
}

// Verify if channelsExporter implements prometheus.Collector
var _ prometheus.Collector = (*channelsExporter)(nil)

func newChannelsExporter(ctx context.Context, client *epgstation.Client, logger log.Logger) *channelsExporter {
	const subsystem = "channels"

	return &channelsExporter{
		ctx:    ctx,
		client: client,
		logger: logger,

		channels: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "services"),
			"Number of services found in EPGStation.",
			[]string{"type", "channel"}, nil),
	}
}

func (e *channelsExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.channels
}

func (e *channelsExporter) Collect(ch chan<- prometheus.Metric) {
	channels, err := e.client.GetChannels(e.ctx)
	if err != nil {
		level.Error(e.logger).Log("msg", "failed to fetch EPGStation channels", "err", err)
		return
	}

	type countKey struct {
		Type    string
		Channel string
	}
	counts := map[countKey]int{}
	for _, channel := range *channels {
		counts[countKey{Type: channel.ChannelType, Channel: channel.Channel}]++
	}

	for key, count := range counts {
		ch <- prometheus.MustNewConstMetric(e.channels, prometheus.GaugeValue, float64(count), key.Type, key.Channel)
	}
}
