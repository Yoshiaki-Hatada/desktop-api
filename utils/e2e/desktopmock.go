/*
   Copyright 2021 Docker Compose CLI authors

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
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/docker/desktop-api"
	"github.com/labstack/echo"
)

// DesktopMockServer a mock registering all invocations
type DesktopMockServer struct {
	socket        string
	requestBodies []string
	e             *echo.Echo
}

// NewDesktopServer instantiate a new DesktopMockServer
func NewDesktopServer(socket string) *DesktopMockServer {
	return &DesktopMockServer{
		socket: socket,
		e:      echo.New(),
	}
}

func (s *DesktopMockServer) handlePost(c echo.Context) error {
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	body := strings.TrimSpace(string(b))
	s.requestBodies = append(s.requestBodies, body)
	return c.String(http.StatusOK, "")
}

// GetUsage get requestBodies
func (s *DesktopMockServer) GetUsage() []string {
	return s.requestBodies
}

// ResetUsage reset requestBodies
func (s *DesktopMockServer) ResetUsage() {
	s.requestBodies = []string{}
}

// Stop stop the mock server
func (s *DesktopMockServer) Stop() {
	_ = s.e.Close()
}

// Start start the mock server
func (s *DesktopMockServer) Start() {
	go func() {
		listener, err := net.Listen("unix", strings.TrimPrefix(s.socket, "unix://"))
		if err != nil {
			log.Fatal(err)
		}
		s.e.Listener = listener
		s.e.POST("/usage", s.handlePost)
		s.e.GET("/usage", s.handlePost)
		_ = s.e.Start(":1323")
	}()
}

// StartReady start the mock server
func (s *DesktopMockServer) StartReady() error {
	s.Start()
	client := desktop_api.NewDockerDesktopClient()
	metricsCommand := desktop_api.MetricsCommand{
		Command: "test",
		Source:  "e2e-test",
		Context: "default",
		Status:  desktop_api.MetricsStatusSuccess,
	}
	for i := 0; i < 10; i++ {
		_, err := client.SendMetrics(context.Background(), metricsCommand)
		if err != nil {
			continue
		}
		if len(s.GetUsage()) > 0 {
			s.ResetUsage()
			return nil
		}
		time.Sleep(time.Second)
	}
	return nil
}
