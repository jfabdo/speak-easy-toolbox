package toolbox

import (
	"evilrus.c/speak-easy/toolbox/toolbox"
	"github.com/mediocregopher/radix/v3"
)

//ErrorHandler pushes an error to the error q
func ErrorHandler(err error) {
	clnt := toolbox.GetClient()
	clnt.Do(radix.Cmd(nil, "rpush", "error", string(err.Error())))
}
