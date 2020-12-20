package rules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

var Bundle = fluent.Bundle{
    Version: "0.0.1",
}

func boolComparison(m fluent.Matcher) {
	m.Match(`$x == true`,
		`$x != true`,
		`$x == false`,
		`$x != false`).
		Report(`omit bool literal in expression`)
}
