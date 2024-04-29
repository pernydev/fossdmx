package main

import (
	"flag"

	"github.com/pernydev/fossdmx/internal/webserver"
	"github.com/pernydev/fossdmx/pkg/global"
	"github.com/pernydev/fossdmx/pkg/registry"
	"github.com/pernydev/fossdmx/pkg/showfile"
)

func main() {
	err := showfile.CreateDataDirectory()
	if err != nil {
		panic(err)
	}

	showfilePath := flag.String("showfile", "none", "Path to showfile")
	fixturesPath := flag.String("fixtures", "fixtures.json", "Path to fixtures file")
	flag.Parse()

	err = registry.RegisterFixturesPath(*fixturesPath)
	if err != nil {
		panic(err)
	}

	if *showfilePath != "none" {
		err := showfile.LoadFromFile(*showfilePath)
		if err != nil {
			panic(err)
		}
	}

	err = showfile.InitializeDemoShowfile()
	if err != nil {
		panic(err)
	}

	go webserver.InitRouter()

	global.Log.Info("FOSSDMX Operational")

	select {}
}
