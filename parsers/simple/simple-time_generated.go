// Code generated by go generate; DO NOT EDIT.
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
package simple

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

func (p *SimpleTime) Init() {
    p.data = nil
}

func (p *SimpleTime) GetParserName() string {
    return p.Name
}

func (p *SimpleTime) Get(createIfNotExist bool) (common.ParserData, error) {
	if p.data == nil {
		if createIfNotExist {
			p.data = &types.StringC{}
			return p.data, nil
		}
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *SimpleTime) GetOne(index int) (common.ParserData, error) {
	if index > 0 {
		return nil, errors.FetchError
	}
	if p.data == nil {
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *SimpleTime) Delete(index int) error {
	p.Init()
	return nil
}

func (p *SimpleTime) Insert(data common.ParserData, index int) error {
	return p.Set(data, index)
}

func (p *SimpleTime) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case *types.StringC:
		p.data = newValue
	case types.StringC:
		p.data = &newValue
	default:
		return errors.InvalidData
	}
	return nil
}
