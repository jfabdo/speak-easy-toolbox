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

//ClientDoReturnNil makes your client do a command
func ClientDoReturnNil(clnt *radix.Pool, cmd string, on string, val string) {
	clnt.Do(radix.Cmd(nil, cmd, on, val))
}
