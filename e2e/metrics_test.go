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

	"github.com/docker/desktop-api"
	"gotest.tools/v3/assert"

	"github.com/docker/desktop-api/desktopcli"
	"github.com/docker/desktop-api/utils/e2e"
)

func TestComposeMetrics(t *testing.T) {
	s := e2e.NewMetricsServer(desktopcli.Socket)
	assert.NilError(t, s.StartReady(), "Metrics mock server not available")
	defer s.Stop()

	t.Run("catch default metrics", func(t *testing.T) {
		client := desktop_api.NewDockerDesktopClient()
		statuses := []string{
			desktop_api.MetricsStatusSuccess,
			desktop_api.MetricsStatusComposeParseFailure,
			desktop_api.MetricsStatusFileNotFoundFailure,
			desktop_api.MetricsStatusCommandSyntaxFailure,
			desktop_api.MetricsStatusPullFailure,
			desktop_api.MetricsStatusBuildFailure,
		}
		for _, s := range statuses {
			_, err := client.SendMetrics(context.Background(), desktop_api.MetricsCommand{
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
