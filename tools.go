//go:build tools
// +build tools

// https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md
package tools

import (
	_ "github.com/air-verse/air"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/swaggo/swag/cmd/swag"
)
