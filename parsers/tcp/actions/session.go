/*
Copyright 2019 HAProxy Technologies

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

package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type Session struct {
	Action   []string
	Cond     string
	CondTest string
	Comment  string
}

func (f *Session) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 3 {
		command, condition := common.SplitRequest(parts[2:])
		if len(command) > 0 {
			f.Action = command
		} else {
			return fmt.Errorf("not enough params")
		}
		if len(condition) > 1 {
			f.Cond = condition[0]
			f.CondTest = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *Session) String() string {
	var result strings.Builder
	result.WriteString("content ")

	result.WriteString(strings.Join(f.Action, " "))

	if f.Cond != "" {
		result.WriteString(" ")
		result.WriteString(f.Cond)
		result.WriteString(" ")
		result.WriteString(f.CondTest)
	}
	if f.Comment != "" {
		result.WriteString(" # ")
		result.WriteString(f.Comment)
	}
	return result.String()
}

func (f *Session) GetComment() string {
	return f.Comment
}
