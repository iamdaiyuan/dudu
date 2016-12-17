package main

import (
	"github.com/beautytop/dudu"
	"github.com/beautytop/dudu/public/core"
)

func main() {
	if dudu.Local {
		core.InitConfig(dudu.Dir+"/config/uk_local_config.json", dudu.Dir+"/config/uk_log.json")
	} else {
		core.InitConfig(dudu.Dir+"/config/uk_config.json", dudu.Dir+"/config/uk_log.json")
	}
	core.InitDB()
}
