// +build !windows

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

package desktopcli

import (
	"net"
	"strings"
)

var (
	Socket = "/var/run/docker-cli.sock"
)

func Conn() (net.Conn, error) {
	return net.Dial("unix", Socket)
}

func ListenSocket(socket string) (net.Listener, error) {
	return net.Listen("unix", strings.TrimPrefix(socket, "unix://"))
}
