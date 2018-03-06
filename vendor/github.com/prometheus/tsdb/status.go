// Copyright 2018 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tsdb

import (
	dto "github.com/prometheus/client_model/go"
)

type tsdbStatus struct {
	ChunkCount int64
	TimeSeriesCount int64
}

func (db *DB) Status() tsdbStatus {
	var tmp dto.Metric
	var status tsdbStatus

	db.head.metrics.chunks.Write(&tmp)
	status.ChunkCount = int64(*tmp.Gauge.Value)

	db.head.metrics.series.Write(&tmp)
	status.TimeSeriesCount = int64(*tmp.Gauge.Value)

	return status
}