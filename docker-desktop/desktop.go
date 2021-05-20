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

package desktop

import (
	"context"
	"net"
	"net/http"

	"github.com/ulyssessouza/desktop-api/docker-desktop/desktopcli"
	dockercliapi "github.com/ulyssessouza/desktop-api/internal/generated/docker-cli"
)

func NewDockerDesktopClient() *DockerDesktopClient {
	return &DockerDesktopClient{
		APIClient: dockercliapi.NewAPIClient(getDockerCliConfiguration()),
	}
}

type DockerDesktopClient struct {
	*dockercliapi.APIClient
}

func NewMetricsCommand() MetricsCommand {
	return MetricsCommand{
		Status: MetricsSuccessStatus,
	}
}

type MetricsCommand struct {
	Command string `json:"command,omitempty"`
	Context string `json:"context,omitempty"`
	Source  string `json:"source,omitempty"`
	Status  string `json:"status,omitempty"`
}

func (d DockerDesktopClient) SendMetrics(ctx context.Context, metrics MetricsCommand) (*http.Response, error) {
	r := d.MetricsApi.SubmitMetrics(ctx)
	m := dockercliapi.MetricsCommand{
		Command: &metrics.Command,
		Context: &metrics.Context,
		Source:  &metrics.Source,
		Status:  &metrics.Status,
	}
	return r.MetricsCommand(m).Execute()
}

func getDockerCliConfiguration() *dockercliapi.Configuration {
	dockercliapiCfg := dockercliapi.NewConfiguration()
	dockercliapiCfg.Host = "localhost"
	dockercliapiCfg.Scheme = "http"
	dockercliapiCfg.UserAgent = userAgent
	dockercliapiCfg.Servers = dockercliapi.ServerConfigurations{
		dockercliapi.ServerConfiguration{
			URL:         "",
		},
	}
	dockercliapiCfg.HTTPClient = &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return desktopcli.Conn()
			},
		},
	}
	return dockercliapiCfg
}
