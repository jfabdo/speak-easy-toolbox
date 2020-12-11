package extras

import "github.com/mediocregopher/radix/v3"

func getClient() *radix.Pool {
	pool, err := radix.NewPool("tcp", "172.17.0.1:6379", 10)
	if err != nil {
		println(err)
		panic(err)
	} else {
		return pool
	}
}
