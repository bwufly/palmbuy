package main

import (
	_ "palmbuy/boot"
	_ "palmbuy/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
