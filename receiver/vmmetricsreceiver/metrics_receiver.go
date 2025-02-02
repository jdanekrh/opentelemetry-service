// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vmmetricsreceiver

import (
	"errors"
	"sync"

	"github.com/open-telemetry/opentelemetry-service/receiver"
)

var (
	errAlreadyStarted = errors.New("already started")
	errAlreadyStopped = errors.New("already stopped")
)

// Receiver is the type used to handle metrics from VM metrics.
type Receiver struct {
	mu sync.Mutex

	vmc *VMMetricsCollector

	stopOnce  sync.Once
	startOnce sync.Once
}

const metricsSource string = "VMMetrics"

// MetricsSource returns the name of the metrics data source.
func (vmr *Receiver) MetricsSource() string {
	return metricsSource
}

// StartMetricsReception scrapes VM metrics based on the OS platform.
func (vmr *Receiver) StartMetricsReception(host receiver.Host) error {
	vmr.mu.Lock()
	defer vmr.mu.Unlock()

	var err = errAlreadyStarted
	vmr.startOnce.Do(func() {
		vmr.vmc.StartCollection()
		err = nil
	})
	return err
}

// StopMetricsReception stops and cancels the underlying VM metrics scrapers.
func (vmr *Receiver) StopMetricsReception() error {
	vmr.mu.Lock()
	defer vmr.mu.Unlock()

	var err = errAlreadyStopped
	vmr.stopOnce.Do(func() {
		vmr.vmc.StopCollection()
		err = nil
	})
	return err
}
