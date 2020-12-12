package toolbox

import (
	"github.com/mediocregopher/radix/v3"
)

//ErrorHandler pushes an error to the error q
func ErrorHandler(err error) {
	clnt := GetPool()
	clnt.Do(radix.Cmd(nil, "rpush", "error", string(err.Error())))
}
