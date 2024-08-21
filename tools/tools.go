//go:build tools

package tools // import "github.com/asphaltbuffet/ohm/tools"

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint" // golangci-lint
	_ "golang.org/x/tools/cmd/stringer"                     // stringer
	_ "gotest.tools/gotestsum"                              // gotestsum
)
