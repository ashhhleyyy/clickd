package tracker

import (
	_ "embed"
	"strings"

	"github.com/ashhhleyyy/clickd/config"
)

//go:generate yarn build

//go:embed build/tracker.js
var script []byte

var replacedScript []byte

func init() {
	replacedScript = []byte(strings.ReplaceAll(string(script), "%CLICKD_HOST%", config.Configuration.Host))
}

func Script() []byte {
	return replacedScript
}
