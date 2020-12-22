package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
	extrarules "github.com/quasilyte/ruleguard-rules-test"
)

var Bundle = dsl.Bundle{}

func init() {
	dsl.ImportRules("extra", extrarules.Bundle)
}
