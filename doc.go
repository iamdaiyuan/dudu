// a go spider!
package dudu

import "github.com/hunterhug/go_tool/util"

var Dir = util.CurDir()
var Local = true

func init() {
	if util.FileExist(Dir + "/远程.txt") {
		Local = false
	}
}
