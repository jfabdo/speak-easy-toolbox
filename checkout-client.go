package toolbox

import "github.com/mediocregopher/radix/v3"

//GetClient is stuff
func GetClient() *radix.Pool {
	pool, err := radix.NewPool("tcp", "127.0.0.1:6379", 10)
	if err != nil {
		println(err)
		panic(err)
	} else {
		return pool
	}
}

func ClientDo()
.Do(radix.Cmd(nil, "SET", "foo", "someval"))