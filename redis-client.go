package toolbox

import (
	"os"

	"github.com/mediocregopher/radix/v3"
)

//GetClient is stuff
func GetClient() *radix.Pool {
	pool, err := radix.NewPool("tcp", os.Getenv("ERU_SE_REDIS_IP"), 10)
	if err != nil {
		println(err)
		panic(err)
	} else {
		return pool
	}
}

//ClientDo makes your client do a command
func ClientDo(clnt *radix.Pool, to interface{}, cmd string, on string, val string) {
	clnt.Do(radix.Cmd(to, cmd, on, val))
}
