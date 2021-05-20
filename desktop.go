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
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/ulyssessouza/desktop-api/desktopcli"
	dockercliapi "github.com/ulyssessouza/desktop-api/internal/generated/docker-cli"
	"github.com/ulyssessouza/desktop-api/utils"
)

type DockerDesktopClient struct {
	*dockercliapi.APIClient
}

func NewDockerDesktopClient() *DockerDesktopClient {
	return &DockerDesktopClient{
		APIClient: dockercliapi.NewAPIClient(getDockerCliConfiguration()),
	}
}

func (d DockerDesktopClient) SendMetrics(ctx context.Context, metrics MetricsCommand) (*http.Response, error) {
	r := d.MetricsApi.PostMetrics(ctx)
	m := dockercliapi.MetricsCommand{
		Command: &metrics.Command,
		Context: &metrics.Context,
		Source:  &metrics.Source,
		Status:  &metrics.Status,
	}
	return r.MetricsCommand(m).Execute()
}

func (d DockerDesktopClient) GetUUID(ctx context.Context) (string, error) {
	r, err := d.UuidApi.GetUuid(ctx).Execute()
	if err != nil {
		return "", err
	}
	if r.StatusCode != 200 {
		return "", fmt.Errorf("status %d while getting uuid", r.StatusCode)
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return "", err
	}
	result := string(body)
	if result == "" {
		return "", fmt.Errorf("empty UUID")
	}
	return result, nil
}

func getDockerCliConfiguration() *dockercliapi.Configuration {
	dockercliapiCfg := dockercliapi.NewConfiguration()
	dockercliapiCfg.Scheme = "http"
	dockercliapiCfg.Host = utils.DummyHost
	dockercliapiCfg.UserAgent = utils.UserAgent
	dockercliapiCfg.Servers = dockercliapi.ServerConfigurations{
		dockercliapi.ServerConfiguration{
			URL: "",
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
