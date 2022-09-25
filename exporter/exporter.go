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
	"github.com/prometheus/client_golang/prometheus"

	"github.com/coord-e/epgstation_exporter/epgstation"
)

const namespace = "epgstation"

type Config struct {
	FetchVersion   bool
	FetchChannels  bool
	FetchSchedules bool
}

type Exporter struct {
	ctx    context.Context
	logger log.Logger

	version   *versionExporter
	channels  *channelsExporter
	schedules *schedulesExporter
}

// Verify if Exporter implements prometheus.Collector
var _ prometheus.Collector = (*Exporter)(nil)

func New(ctx context.Context, client *epgstation.Client, config Config, logger log.Logger) *Exporter {
	var versionExporter *versionExporter
	if config.FetchVersion {
		versionExporter = newVersionExporter(ctx, client, logger)
	}

	var channelsExporter *channelsExporter
	if config.FetchChannels {
		channelsExporter = newChannelsExporter(ctx, client, logger)
	}

	var schedulesExporter *schedulesExporter
	if config.FetchSchedules {
		schedulesExporter = newSchedulesExporter(ctx, client, logger)
	}

	return &Exporter{
		ctx:       ctx,
		logger:    logger,
		version:   versionExporter,
		channels:  channelsExporter,
		schedules: schedulesExporter,
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	if e.version != nil {
		e.version.Describe(ch)
	}
	if e.channels != nil {
		e.channels.Describe(ch)
	}
	if e.schedules != nil {
		e.schedules.Describe(ch)
	}
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	if e.version != nil {
		e.version.Collect(ch)
	}
	if e.channels != nil {
		e.channels.Collect(ch)
	}
	if e.schedules != nil {
		e.schedules.Collect(ch)
	}
}