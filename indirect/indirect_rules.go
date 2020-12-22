package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl/fluent"
	extrarules "github.com/quasilyte/ruleguard-rules-test"
)

var Bundle = fluent.Bundle{}

func init() {
	fluent.ImportRules("extra", extrarules.Bundle)
}
