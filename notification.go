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
	"errors"
	"fmt"
)

type NotificationLevel int

const (
	NotificationLevelInfo NotificationLevel = iota
	NotificationLevelWarning
	NotificationLevelError
)

func (n NotificationLevel) Validate() error {
	if n < NotificationLevelInfo || n > NotificationLevelError {
		return fmt.Errorf("toast 'level' '%d' is invalid", n)
	}
	return nil
}

// NotificationModel is a command
type NotificationModel struct {
	Title string            `json:"title"`
	Body  string            `json:"body"`
	Level NotificationLevel `json:"level"`
}

func (n NotificationModel) Validate() error {
	if n.Title == "" {
		return errors.New("toast 'title' cannot be empty")
	}
	if n.Body == "" {
		return errors.New("toast 'body' cannot be empty")
	}
	return n.Level.Validate()
}
