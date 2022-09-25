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
	"strconv"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/coord-e/epgstation_exporter/epgstation"
)

type schedulesExporter struct {
	ctx    context.Context
	client *epgstation.Client
	logger log.Logger

	next1HourPrograms *prometheus.Desc
	next1DayPrograms  *prometheus.Desc
}

// Verify if schedulesExporter implements prometheus.Collector
var _ prometheus.Collector = (*schedulesExporter)(nil)

func newSchedulesExporter(ctx context.Context, client *epgstation.Client, logger log.Logger) *schedulesExporter {
	const subsystem = "schedules"

	return &schedulesExporter{
		ctx:    ctx,
		client: client,
		logger: logger,

		next1HourPrograms: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "next_1hour_programs"),
			"Number of programs starts in next 1 hour, found in EPGStation.",
			[]string{"service_id"}, nil),
		next1DayPrograms: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "next_1day_programs"),
			"Number of programs starts in next 1 day, found in EPGStation.",
			[]string{"service_id"}, nil),
	}
}

func (e *schedulesExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.next1HourPrograms
	ch <- e.next1DayPrograms
}

func (e *schedulesExporter) Collect(ch chan<- prometheus.Metric) {
	now := time.Now()

	schedules, err := e.client.GetSchedules(e.ctx, epgstation.GetSchedulesOpts{
		StartAt:          now.UnixMilli(),
		EndAt:            now.Add(24 * time.Hour).UnixMilli(),
		NeedsRawExtended: nil,
		IsFree:           nil,
		IsHalfWidth:      true,
		GR:               true,
		BS:               true,
		CS:               true,
		SKY:              true,
	})
	if err != nil {
		level.Error(e.logger).Log("msg", "failed to fetch EPGStation schedules", "err", err)
		return
	}

	for _, schedule := range *schedules {
		hourCount := 0
		dayCount := 0
		for _, program := range schedule.Programs {
			startAt := time.Unix(0, program.StartAt*1000*1000)

			if startAt.After(now) && startAt.Before(now.Add(time.Hour)) {
				hourCount++
			}
			if startAt.After(now) && startAt.Before(now.Add(24*time.Hour)) {
				dayCount++
			}
		}
		serviceID := strconv.Itoa(schedule.Channel.ServiceID)
		ch <- prometheus.MustNewConstMetric(e.next1HourPrograms, prometheus.GaugeValue, float64(hourCount), serviceID)
		ch <- prometheus.MustNewConstMetric(e.next1DayPrograms, prometheus.GaugeValue, float64(dayCount), serviceID)
	}
}
