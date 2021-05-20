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

const (
	DummyHost = "localhost"
	UserAgent = "Docker Desktop API"
)

const (
	// MetricsSuccessStatus command success
	MetricsSuccessStatus = "success"
	// MetricsFailureStatus command failure
	MetricsFailureStatus = "failure"
	// MetricsComposeParseFailureStatus failure while parsing compose file
	MetricsComposeParseFailureStatus = "failure-compose-parse"
	// MetricsFileNotFoundFailureStatus failure getting compose file
	MetricsFileNotFoundFailureStatus = "failure-file-not-found"
	// MetricsCommandSyntaxFailureStatus failure reading command
	MetricsCommandSyntaxFailureStatus = "failure-cmd-syntax"
	// MetricsBuildFailureStatus failure building image
	MetricsBuildFailureStatus = "failure-build"
	// MetricsPullFailureStatus failure pulling image
	MetricsPullFailureStatus = "failure-pull"
	// MetricsCanceledStatus command canceled
	MetricsCanceledStatus = "canceled"
)
