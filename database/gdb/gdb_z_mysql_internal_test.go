// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gdb

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

var (
	db  DB
	ctx = context.TODO()
)

func Test_parseConfigNodeLink_WithType(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `mysql:root:CxzhD*624:27jh@tcp(9.135.69.119:3306)/khaos_oss?loc=Local&parseTime=true&charset=latin`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.Type, `mysql`)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, `khaos_oss`)
		t.Assert(newNode.Extra, `loc=Local&parseTime=true&charset=latin`)
		t.Assert(newNode.Charset, `latin`)
		t.Assert(newNode.Protocol, `tcp`)
	})
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `mysql:root:CxzhD*624:27jh@tcp(9.135.69.119:3306)/khaos_oss?`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.Type, `mysql`)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, `khaos_oss`)
		t.Assert(newNode.Extra, ``)
		t.Assert(newNode.Charset, defaultCharset)
		t.Assert(newNode.Protocol, `tcp`)
	})
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `mysql:root:CxzhD*624:27jh@tcp(9.135.69.119:3306)/khaos_oss`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.Type, `mysql`)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, `khaos_oss`)
		t.Assert(newNode.Extra, ``)
		t.Assert(newNode.Charset, defaultCharset)
		t.Assert(newNode.Protocol, `tcp`)
	})
	// empty database preselect.
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `mysql:root:CxzhD*624:27jh@tcp(9.135.69.119:3306)/?loc=Local&parseTime=true&charset=latin`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.Type, `mysql`)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, ``)
		t.Assert(newNode.Extra, `loc=Local&parseTime=true&charset=latin`)
		t.Assert(newNode.Charset, `latin`)
		t.Assert(newNode.Protocol, `tcp`)
	})
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `mysql:root:CxzhD*624:27jh@tcp(9.135.69.119:3306)?loc=Local&parseTime=true&charset=latin`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.Type, `mysql`)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, ``)
		t.Assert(newNode.Extra, `loc=Local&parseTime=true&charset=latin`)
		t.Assert(newNode.Charset, `latin`)
		t.Assert(newNode.Protocol, `tcp`)
	})
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `mysql:root:CxzhD*624:27jh@tcp(9.135.69.119:3306)/`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.Type, `mysql`)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, ``)
		t.Assert(newNode.Extra, ``)
		t.Assert(newNode.Charset, defaultCharset)
		t.Assert(newNode.Protocol, `tcp`)
	})
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `mysql:root:CxzhD*624:27jh@tcp(9.135.69.119:3306)`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.Type, `mysql`)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, ``)
		t.Assert(newNode.Extra, ``)
		t.Assert(newNode.Charset, defaultCharset)
		t.Assert(newNode.Protocol, `tcp`)
	})
	// udp.
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `mysql:root:CxzhD*624:27jh@udp(9.135.69.119:3306)`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.Type, `mysql`)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, ``)
		t.Assert(newNode.Extra, ``)
		t.Assert(newNode.Charset, defaultCharset)
		t.Assert(newNode.Protocol, `udp`)
	})
}

func Test_parseConfigNodeLink_WithoutType(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `root:CxzhD*624:27jh@tcp(9.135.69.119:3306)/khaos_oss?loc=Local&parseTime=true&charset=latin`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.Type, ``)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, `khaos_oss`)
		t.Assert(newNode.Extra, `loc=Local&parseTime=true&charset=latin`)
		t.Assert(newNode.Charset, `latin`)
		t.Assert(newNode.Protocol, `tcp`)
	})
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `root:CxzhD*624:27jh@tcp(9.135.69.119:3306)/khaos_oss?`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.Type, ``)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, `khaos_oss`)
		t.Assert(newNode.Extra, ``)
		t.Assert(newNode.Charset, defaultCharset)
		t.Assert(newNode.Protocol, `tcp`)
	})
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `root:CxzhD*624:27jh@tcp(9.135.69.119:3306)/khaos_oss`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.Type, ``)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, `khaos_oss`)
		t.Assert(newNode.Extra, ``)
		t.Assert(newNode.Charset, defaultCharset)
		t.Assert(newNode.Protocol, `tcp`)
	})
	// empty database preselect.
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `root:CxzhD*624:27jh@tcp(9.135.69.119:3306)/?loc=Local&parseTime=true&charset=latin`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, ``)
		t.Assert(newNode.Extra, `loc=Local&parseTime=true&charset=latin`)
		t.Assert(newNode.Charset, `latin`)
		t.Assert(newNode.Protocol, `tcp`)
	})
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `root:CxzhD*624:27jh@tcp(9.135.69.119:3306)?loc=Local&parseTime=true&charset=latin`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, ``)
		t.Assert(newNode.Extra, `loc=Local&parseTime=true&charset=latin`)
		t.Assert(newNode.Charset, `latin`)
		t.Assert(newNode.Protocol, `tcp`)
	})
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `root:CxzhD*624:27jh@tcp(9.135.69.119:3306)/`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, ``)
		t.Assert(newNode.Extra, ``)
		t.Assert(newNode.Charset, defaultCharset)
		t.Assert(newNode.Protocol, `tcp`)
	})
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `root:CxzhD*624:27jh@tcp(9.135.69.119:3306)`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, ``)
		t.Assert(newNode.Extra, ``)
		t.Assert(newNode.Charset, defaultCharset)
		t.Assert(newNode.Protocol, `tcp`)
	})
	// protocol.
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `root:CxzhD*624:27jh@udp(9.135.69.119:3306)`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, ``)
		t.Assert(newNode.Extra, ``)
		t.Assert(newNode.Charset, defaultCharset)
		t.Assert(newNode.Protocol, `udp`)
	})
}

func Test_parseConfigNodeLink_WithoutType_Clickhouse(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		node := ConfigNode{
			Link: `clickhouse://default@127.0.0.1:9000,127.0.0.1:9000/default?dial_timeout=200ms&max_execution_time=60`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.Type, ``)
		t.Assert(newNode.User, `root`)
		t.Assert(newNode.Pass, `CxzhD*624:27jh`)
		t.Assert(newNode.Host, `9.135.69.119`)
		t.Assert(newNode.Port, `3306`)
		t.Assert(newNode.Name, `khaos_oss`)
		t.Assert(newNode.Extra, `loc=Local&parseTime=true&charset=latin`)
		t.Assert(newNode.Charset, `latin`)
		t.Assert(newNode.Protocol, `tcp`)
	})
}

func Test_Func_doQuoteWord(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := map[string]string{
			"user":                   "`user`",
			"user u":                 "user u",
			"user_detail":            "`user_detail`",
			"user,user_detail":       "user,user_detail",
			"user u, user_detail ut": "user u, user_detail ut",
			"u.id asc":               "u.id asc",
			"u.id asc, ut.uid desc":  "u.id asc, ut.uid desc",
		}
		for k, v := range array {
			t.Assert(doQuoteWord(k, "`", "`"), v)
		}
	})
}

func Test_Func_doQuoteString(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := map[string]string{
			"user":                             "`user`",
			"user u":                           "`user` u",
			"user,user_detail":                 "`user`,`user_detail`",
			"user u, user_detail ut":           "`user` u,`user_detail` ut",
			"u.id, u.name, u.age":              "`u`.`id`,`u`.`name`,`u`.`age`",
			"u.id asc":                         "`u`.`id` asc",
			"u.id asc, ut.uid desc":            "`u`.`id` asc,`ut`.`uid` desc",
			"user.user u, user.user_detail ut": "`user`.`user` u,`user`.`user_detail` ut",
			// mssql global schema access with double dots.
			"user..user u, user.user_detail ut": "`user`..`user` u,`user`.`user_detail` ut",
		}
		for k, v := range array {
			t.Assert(doQuoteString(k, "`", "`"), v)
		}
	})
}

func Test_Func_addTablePrefix(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		prefix := ""
		array := map[string]string{
			"user":                         "`user`",
			"user u":                       "`user` u",
			"user as u":                    "`user` as u",
			"user,user_detail":             "`user`,`user_detail`",
			"user u, user_detail ut":       "`user` u,`user_detail` ut",
			"`user`.user_detail":           "`user`.`user_detail`",
			"`user`.`user_detail`":         "`user`.`user_detail`",
			"user as u, user_detail as ut": "`user` as u,`user_detail` as ut",
			"UserCenter.user as u, UserCenter.user_detail as ut": "`UserCenter`.`user` as u,`UserCenter`.`user_detail` as ut",
			// mssql global schema access with double dots.
			"UserCenter..user as u, user_detail as ut": "`UserCenter`..`user` as u,`user_detail` as ut",
		}
		for k, v := range array {
			t.Assert(doQuoteTableName(k, prefix, "`", "`"), v)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		prefix := "gf_"
		array := map[string]string{
			"user":                         "`gf_user`",
			"user u":                       "`gf_user` u",
			"user as u":                    "`gf_user` as u",
			"user,user_detail":             "`gf_user`,`gf_user_detail`",
			"user u, user_detail ut":       "`gf_user` u,`gf_user_detail` ut",
			"`user`.user_detail":           "`user`.`gf_user_detail`",
			"`user`.`user_detail`":         "`user`.`gf_user_detail`",
			"user as u, user_detail as ut": "`gf_user` as u,`gf_user_detail` as ut",
			"UserCenter.user as u, UserCenter.user_detail as ut": "`UserCenter`.`gf_user` as u,`UserCenter`.`gf_user_detail` as ut",
			// mssql global schema access with double dots.
			"UserCenter..user as u, user_detail as ut": "`UserCenter`..`gf_user` as u,`gf_user_detail` as ut",
		}
		for k, v := range array {
			t.Assert(doQuoteTableName(k, prefix, "`", "`"), v)
		}
	})
}

func Test_isSubQuery(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(isSubQuery("user"), false)
		t.Assert(isSubQuery("user.uid"), false)
		t.Assert(isSubQuery("u, user.uid"), false)
		t.Assert(isSubQuery("select 1"), true)
	})
}
