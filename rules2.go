package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

func boolExprSimplify(m fluent.Matcher) {
	m.Match(`!!$x`).Suggest(`$x`)
	m.Match(`!($x != $y)`).Suggest(`$x == $y`)
	m.Match(`!($x == $y)`).Suggest(`$x != $y`)
}
