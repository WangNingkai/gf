// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package builtin

import (
	"errors"
	"strings"

	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/util/gutil"
)

// Required if all given fields are not empty.
// Example: required-with-all:id,name

type RuleRequiredWithAll struct{}

func init() {
	Register(&RuleRequiredWithAll{})
}

func (r *RuleRequiredWithAll) Name() string {
	return "required-with-all"
}

func (r *RuleRequiredWithAll) Message() string {
	return "The {attribute} field is required"
}

func (r *RuleRequiredWithAll) Run(in RunInput) error {
	var (
		required   = true
		array      = strings.Split(in.RulePattern, ",")
		foundValue interface{}
	)
	for i := 0; i < len(array); i++ {
		_, foundValue = gutil.MapPossibleItemByKey(in.Data.Map(), array[i])
		if empty.IsEmpty(foundValue) {
			required = false
			break
		}
	}

	if required && in.Value.IsEmpty() {
		return errors.New(in.Message)
	}
	return nil
}
