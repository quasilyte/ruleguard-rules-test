package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func boolExprSimplify(m dsl.Matcher) {
	m.Match(`!!$x`).Suggest(`$x`)
	m.Match(`!($x != $y)`).Suggest(`$x == $y`)
	m.Match(`!($x == $y)`).Suggest(`$x != $y`)
}
