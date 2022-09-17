// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package builtin

import (
	"errors"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

type RuleDifferent struct{}

func init() {
	Register(&RuleDifferent{})
}

func (r *RuleDifferent) Name() string {
	return "different"
}

func (r *RuleDifferent) Message() string {
	return "The {attribute} value `{value}` must be different from field {pattern}"
}

func (r *RuleDifferent) Run(in RunInput) error {
	var (
		ok    = true
		value = in.Value.String()
	)
	_, foundValue := gutil.MapPossibleItemByKey(in.Data.Map(), in.RulePattern)
	if foundValue != nil {
		if in.CaseInsensitive {
			ok = !strings.EqualFold(value, gconv.String(foundValue))
		} else {
			ok = strings.Compare(value, gconv.String(foundValue)) != 0
		}
	}
	if !ok {
		return errors.New(in.Message)
	}
	return nil
}
