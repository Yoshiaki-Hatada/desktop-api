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

package e2e

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"

	desktop "github.com/ulyssessouza/desktop-api/docker-desktop"
	"github.com/ulyssessouza/desktop-api/docker-desktop/desktopcli"
	"github.com/ulyssessouza/desktop-api/utils/e2e"
)

func TestComposeMetrics(t *testing.T) {
	s := e2e.NewMetricsServer(desktopcli.Socket)
	assert.NilError(t, s.StartReady(), "Metrics mock server not available")
	defer s.Stop()

	t.Run("catch default metrics", func(t *testing.T) {
		client := desktop.NewDockerDesktopClient()
		statuses := []string{
			desktop.MetricsSuccessStatus,
			desktop.MetricsComposeParseFailureStatus,
			desktop.MetricsFileNotFoundFailureStatus,
			desktop.MetricsCommandSyntaxFailureStatus,
			desktop.MetricsPullFailureStatus,
			desktop.MetricsBuildFailureStatus,
		}
		for _, s := range statuses {
			_, err := client.SendMetrics(context.Background(), desktop.MetricsCommand{
				Command: "test",
				Context: "moby",
				Source:  "e2e-test",
				Status:  s,
			})
			assert.NilError(t, err)
		}
		assert.DeepEqual(t, []string{
			`{"command":"test","context":"moby","source":"e2e-test","status":"success"}`,
			`{"command":"test","context":"moby","source":"e2e-test","status":"failure-compose-parse"}`,
			`{"command":"test","context":"moby","source":"e2e-test","status":"failure-file-not-found"}`,
			`{"command":"test","context":"moby","source":"e2e-test","status":"failure-cmd-syntax"}`,
			`{"command":"test","context":"moby","source":"e2e-test","status":"failure-pull"}`,
			`{"command":"test","context":"moby","source":"e2e-test","status":"failure-build"}`,
		}, s.GetUsage())
	})
}