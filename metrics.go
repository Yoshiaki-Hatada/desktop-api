/*
   Copyright 2021 Docker Compose e2e-test authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package desktop_api

import (
	"fmt"
)

type MetricsStatus string

const (
	// MetricsSuccessStatus command success
	MetricsSuccessStatus MetricsStatus = "success"
	// MetricsFailureStatus command failure
	MetricsFailureStatus MetricsStatus = "failure"
	// MetricsComposeParseFailureStatus failure while parsing compose file
	MetricsComposeParseFailureStatus MetricsStatus = "failure-compose-parse"
	// MetricsFileNotFoundFailureStatus failure getting compose file
	MetricsFileNotFoundFailureStatus MetricsStatus = "failure-file-not-found"
	// MetricsCommandSyntaxFailureStatus failure reading command
	MetricsCommandSyntaxFailureStatus MetricsStatus = "failure-cmd-syntax"
	// MetricsBuildFailureStatus failure building image
	MetricsBuildFailureStatus MetricsStatus = "failure-build"
	// MetricsPullFailureStatus failure pulling image
	MetricsPullFailureStatus MetricsStatus = "failure-pull"
	// MetricsCanceledStatus command canceled
	MetricsCanceledStatus MetricsStatus = "canceled"
)

func (m MetricsStatus) Validate() error {
	for _, s := range []MetricsStatus{
		MetricsSuccessStatus,
		MetricsFailureStatus,
		MetricsComposeParseFailureStatus,
		MetricsFileNotFoundFailureStatus,
		MetricsCommandSyntaxFailureStatus,
		MetricsBuildFailureStatus,
		MetricsPullFailureStatus,
		MetricsCanceledStatus,
	}{
		if s == m {
			return nil
		}
	}
	return fmt.Errorf("MetricsStatus %q is not valid", m)
}

func (m MetricsStatus) String() string {
	return string(m)
}

func NewMetricsCommand() MetricsCommand {
	return MetricsCommand{
		Status: MetricsSuccessStatus,
	}
}

type MetricsCommand struct {
	Command string        `json:"command,omitempty"`
	Context string        `json:"context,omitempty"`
	Source  string        `json:"source,omitempty"`
	Status  MetricsStatus `json:"status,omitempty"`
}

func (m MetricsCommand) Validate() error {
	return m.Status.Validate()
}