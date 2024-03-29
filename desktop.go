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
	"time"

	"github.com/docker/desktop-api/desktopcli"
	dockercliapi "github.com/docker/desktop-api/internal/generated/docker-cli"
)

type DockerDesktopClient struct {
	*dockercliapi.APIClient
}

func NewDockerDesktopClient() *DockerDesktopClient {
	return &DockerDesktopClient{
		APIClient: dockercliapi.NewAPIClient(getDockerCliConfiguration()),
	}
}

func (d *DockerDesktopClient) WithSocketPath(path string) *DockerDesktopClient {
	desktopcli.Socket = path
	d.APIClient.GetConfig().HTTPClient.Transport = getHttpTransport()
	return d
}

func (d *DockerDesktopClient) WithTimeout(t time.Duration) *DockerDesktopClient {
	if d.APIClient == nil || d.APIClient.GetConfig().HTTPClient == nil {
		return d
	}
	d.APIClient.GetConfig().HTTPClient.Timeout = t
	return d
}

func (d DockerDesktopClient) SendMetrics(ctx context.Context, metrics MetricsCommand) (*http.Response, error) {
	err := metrics.Validate()
	if err != nil {
		return nil, err
	}
	r := d.MetricsApi.PostMetrics(ctx)
	var status = metrics.Status.String()
	m := dockercliapi.MetricsCommand{
		Command: &metrics.Command,
		Context: &metrics.Context,
		Source:  &metrics.Source,
		Status:  &status,
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

// SendToast notification to the Docker Desktop server
func (d DockerDesktopClient) SendToast(ctx context.Context, notification NotificationModel) (*http.Response, error) {
	err := notification.Validate()
	if err != nil {
		return nil, err
	}
	var level int32 = int32(notification.Level)
	n := dockercliapi.NotificationModel{
		Title: &notification.Title,
		Body:  &notification.Body,
		Level: &level,
	}
	r := d.NotificationApi.ShowToast(ctx)
	return r.NotificationModel(n).Execute()
}

// IsProUser checks if it's a Pro User
func (d DockerDesktopClient) IsProUser(ctx context.Context) (bool, error) {
	f, r, err := d.FeaturesApi.GetFeatures(ctx).Execute()
	if err != nil {
		return false, err
	}
	if r.StatusCode != 200 {
		return false, fmt.Errorf("status %d while checking if it's a Pro User", r.StatusCode)
	}

	return f.Prouser != nil && f.Prouser.Enabled, nil
}

func getDockerCliConfiguration() *dockercliapi.Configuration {
	dockercliapiCfg := dockercliapi.NewConfiguration()
	dockercliapiCfg.Scheme = "http"
	dockercliapiCfg.Host = DummyHost
	dockercliapiCfg.UserAgent = UserAgent
	dockercliapiCfg.Servers = dockercliapi.ServerConfigurations{
		dockercliapi.ServerConfiguration{
			URL: "",
		},
	}
	dockercliapiCfg.HTTPClient = &http.Client{
		Timeout:   500 * time.Millisecond,
		Transport: getHttpTransport(),
	}
	return dockercliapiCfg
}

func getHttpTransport() *http.Transport {
	return &http.Transport{
		DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
			return desktopcli.Conn()
		},
	}
}
