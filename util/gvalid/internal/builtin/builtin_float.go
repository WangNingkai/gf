// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package builtin

import (
	"errors"
	"strconv"
)

type RuleFloat struct{}

func init() {
	Register(&RuleFloat{})
}

func (r *RuleFloat) Name() string {
	return "float"
}

func (r *RuleFloat) Message() string {
	return "The {attribute} value `{value}` is not a float"
}

func (r *RuleFloat) Run(in RunInput) error {
	if _, err := strconv.ParseFloat(in.Value.String(), 10); err == nil {
		return nil
	}
	return errors.New(in.Message)
}
