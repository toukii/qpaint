package main

import (
	"github.com/zserge/lorca"
)

func main() {
	lorca.Embed("common/assets.go", "common", "www")
}
