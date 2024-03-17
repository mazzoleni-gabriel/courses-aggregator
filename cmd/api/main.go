package main

import "github.com/mazzoleni-gabriel/courses-aggregator/cmd/api/modules"

func main() {
	app := modules.NewApp()
	app.Run()
}
