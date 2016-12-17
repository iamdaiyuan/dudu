package main

import (
	"github.com/beautytop/dudu/public/core"
	"github.com/beautytop/dudu"
)

func main() {
	if dudu.Local {
		core.InitConfig(dudu.Dir+"/config/usa_local_config.json", dudu.Dir+"/config/usa_log.json")
	} else {
		core.InitConfig(dudu.Dir+"/config/usa_config.json", dudu.Dir+"/config/usa_log.json")
	}
	core.AsinPool()
}
