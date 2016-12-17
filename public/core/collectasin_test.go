package core

import (
	"github.com/hunterhug/go_tool/util"
	"testing"
)

func TestCollectAsin(t *testing.T) {
	files, err := util.ListDir(DataDir+"/list/20161110", "html")
	if err != nil {
		panic(err)
	}
	CollectAsin(files)
}
