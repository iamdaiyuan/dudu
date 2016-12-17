package main

import (
	"github.com/beautytop/dudu/public/core"
	"github.com/beautytop/dudu"
)

func main() {
	if dudu.Local {
		core.InitConfig(dudu.Dir+"/config/de_local_config.json", dudu.Dir+"/config/de_log.json")
	} else {
		core.InitConfig(dudu.Dir+"/config/de_config.json", dudu.Dir+"/config/de_log.json")
	}
	core.AsinPool()
}
