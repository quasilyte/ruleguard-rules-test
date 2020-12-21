package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

var Bundle = fluent.Bundle{}

func unslice(m fluent.Matcher) {
	m.Match(`$s[:]`).Where(m["s"].Type.Is(`string`)).Suggest(`$s`)
	m.Match(`$s[:]`).Where(m["s"].Type.Is(`[]$_`)).Suggest(`$s`)
}
