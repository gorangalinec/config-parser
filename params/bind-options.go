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

package params

import (
	"fmt"
	"strings"
)

//BindOption ...
type BindOption interface {
	Parse(options []string, currentIndex int) (int, error)
	Valid() bool
	String() string
}

//BindOptionWord ...
type BindOptionWord struct {
	Name string
}

//Parse ...
func (b *BindOptionWord) Parse(options []string, currentIndex int) (int, error) {
	if currentIndex < len(options) {
		if options[currentIndex] == b.Name {
			return 1, nil
		}
		return 0, &ErrNotFound{Have: options[currentIndex], Want: b.Name}
	}
	return 0, &ErrNotEnoughParams{}
}

//Valid ...
func (b *BindOptionWord) Valid() bool {
	return b.Name != ""
}

//String ...
func (b *BindOptionWord) String() string {
	return b.Name
}

//BindOptionDoubleWord ...
type BindOptionDoubleWord struct {
	Name  string
	Value string
}

//Parse ...
func (b *BindOptionDoubleWord) Parse(options []string, currentIndex int) (int, error) {
	if currentIndex+1 < len(options) {
		if options[currentIndex] == b.Name && b.Value == options[currentIndex+1] {
			return 2, nil
		}
		return 0, &ErrNotFound{
			Have: fmt.Sprintf("%s %s", options[currentIndex], options[currentIndex]),
			Want: fmt.Sprintf("%s %s", b.Name, b.Value),
		}
	}
	return 0, &ErrNotEnoughParams{}
}

//Valid ...
func (b *BindOptionDoubleWord) Valid() bool {
	return b.Name != "" && b.Value != ""
}

//String ...
func (b *BindOptionDoubleWord) String() string {
	if b.Name == "" || b.Value == "" {
		return ""
	}
	return fmt.Sprintf("%s %s", b.Name, b.Value)
}

//BindOptionValue ...
type BindOptionValue struct {
	Name  string
	Value string
}

//Parse ...
func (b *BindOptionValue) Parse(options []string, currentIndex int) (int, error) {
	if currentIndex+1 < len(options) {
		if options[currentIndex] == b.Name {
			b.Value = options[currentIndex+1]
			return 2, nil
		}
		return 0, &ErrNotFound{Have: options[currentIndex], Want: b.Name}
	}
	return 0, &ErrNotEnoughParams{}
}

//Valid ...
func (b *BindOptionValue) Valid() bool {
	return b.Value != ""
}

//String ...
func (b *BindOptionValue) String() string {
	if b.Name == "" || b.Value == "" {
		return ""
	}
	return fmt.Sprintf("%s %s", b.Name, b.Value)
}

func getBindOptions() []BindOption {
	return []BindOption{
		&BindOptionWord{Name: "accept-proxy"},
		&BindOptionWord{Name: "allow-0rtt"},
		&BindOptionWord{Name: "defer-accept"},
		&BindOptionWord{Name: "force-sslv3"},
		&BindOptionWord{Name: "force-tlsv10"},
		&BindOptionWord{Name: "force-tlsv11"},
		&BindOptionWord{Name: "force-tlsv12"},
		&BindOptionWord{Name: "force-tlsv13"},
		&BindOptionWord{Name: "generate-certificates"},
		&BindOptionWord{Name: "prefer-client-ciphers"},
		&BindOptionWord{Name: "ssl"},
		&BindOptionWord{Name: "strict-sni"},
		&BindOptionWord{Name: "tfo"},
		&BindOptionWord{Name: "transparent"},
		&BindOptionWord{Name: "v4v6"},
		&BindOptionWord{Name: "v6only"},

		&BindOptionDoubleWord{Name: "expose-fd", Value: "listeners"},

		&BindOptionValue{Name: "accept-netscaler-cip"},
		&BindOptionValue{Name: "alpn"},
		&BindOptionValue{Name: "backlog"},
		&BindOptionValue{Name: "curves"},
		&BindOptionValue{Name: "ecdhe"},
		&BindOptionValue{Name: "ca-file"},
		&BindOptionValue{Name: "ca-ignore-err"},
		&BindOptionValue{Name: "ca-sign-file"},
		&BindOptionValue{Name: "ca-sign-pass"},
		&BindOptionValue{Name: "ciphers"},
		&BindOptionValue{Name: "crl-file"},
		&BindOptionValue{Name: "crt"},
		&BindOptionValue{Name: "crt-ignore-err"},
		&BindOptionValue{Name: "crt-list"},
		&BindOptionValue{Name: "gid"},
		&BindOptionValue{Name: "group"},
		&BindOptionValue{Name: "id"},
		&BindOptionValue{Name: "interface"},
		&BindOptionValue{Name: "level"},
		&BindOptionValue{Name: "severity-output"},
		&BindOptionValue{Name: "maxconn"},
		&BindOptionValue{Name: "mode"},
		&BindOptionValue{Name: "mss"},
		&BindOptionValue{Name: "name"},
		&BindOptionValue{Name: "namespace"},
		&BindOptionValue{Name: "nice"},
		&BindOptionValue{Name: "no-ca-names"},
		&BindOptionValue{Name: "no-sslv3"},
		&BindOptionValue{Name: "no-tlsv10"},
		&BindOptionValue{Name: "no-tlsv11"},
		&BindOptionValue{Name: "no-tlsv12"},
		&BindOptionValue{Name: "no-tlsv13"},
		&BindOptionValue{Name: "npm"},
		&BindOptionValue{Name: "process"},
		&BindOptionValue{Name: "ssl-max-ver"},
		&BindOptionValue{Name: "ssl-min-ver"},
		&BindOptionValue{Name: "tcp-ut"},
		&BindOptionValue{Name: "tls-ticket-keys"},
		&BindOptionValue{Name: "uid"},
		&BindOptionValue{Name: "user"},
		&BindOptionValue{Name: "verify"},
	}
}

//Parse ...
func ParseBindOptions(options []string) []BindOption {

	result := []BindOption{}
	currentIndex := 0
	for currentIndex < len(options) {
		found := false
		for _, parser := range getBindOptions() {
			if size, err := parser.Parse(options, currentIndex); err == nil {
				result = append(result, parser)
				found = true
				currentIndex += size
			}
		}
		if !found {
			currentIndex++
		}
	}
	return result
}

func BindOptionsString(options []BindOption) string {
	var sb strings.Builder
	first := true
	for _, parser := range options {
		if parser.Valid() {
			if !first {
				sb.WriteString(" ")
			} else {
				first = false
			}
			sb.WriteString(parser.String())
		}
	}
	return sb.String()
}

//CreateBindOptionWord creates valid one word value
func CreateBindOptionWord(name string) (BindOptionWord, ErrParseBindOption) {
	b := BindOptionWord{
		Name: name,
	}
	_, err := b.Parse([]string{name}, 0)
	return b, err
}

//CreateBindOptionDoubleWord creates valid two word value
func CreateBindOptionDoubleWord(name1 string, name2 string) (BindOptionDoubleWord, ErrParseBindOption) {
	b := BindOptionDoubleWord{
		Name:  name1,
		Value: name2,
	}
	_, err := b.Parse([]string{name1, name2}, 0)
	return b, err
}

//CreateBindOptionValue creates valid option with value
func CreateBindOptionValue(name string, value string) (BindOptionValue, ErrParseBindOption) {
	b := BindOptionValue{
		Name:  name,
		Value: value,
	}
	_, err := b.Parse([]string{name, value}, 0)
	return b, err
}
