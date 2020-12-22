package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

var Bundle = dsl.Bundle{}

func unslice(m dsl.Matcher) {
	m.Match(`$s[:]`).Where(m["s"].Type.Is(`string`)).Suggest(`$s`)
	m.Match(`$s[:]`).Where(m["s"].Type.Is(`[]$_`)).Suggest(`$s`)
}
