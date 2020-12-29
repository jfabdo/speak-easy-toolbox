package toolbox

import (
	"os"

	"github.com/mediocregopher/radix/v3"
)

//GetConn returns a new pool
func GetConn() *radix.Conn {
	conn, err := radix.Dial("tcp", os.Getenv("ERU_SE_REDIS_IP"))
	if err != nil {
		println(err)
		panic(err)
	} else {
		return &conn
	}
}

//GetPool returns a new pool
func GetPool() *radix.Pool {
	pool, err := radix.NewPool("tcp", os.Getenv("ERU_SE_REDIS_IP"), 10)
	if err != nil {
		println(err)
		panic(err)
	} else {
		return pool
	}
}

//GetKey returns a new pool
func GetKey(key string, topic string, conn *radix.Pool) string {
	if conn == nil {
		conn = GetPool()
	}
	var value string
	err := conn.Do(radix.Cmd(value, "hget", topic, key))
	if err != nil {
		println(err)
		panic(err)
	}
	return value
}

//GetAll gets everything at that key
func GetAll(key string, conn *radix.Pool) string {
	if conn == nil {
		conn = GetPool()
	}
	var value string
	err := conn.Do(radix.Cmd(value, "hgetall", key))
	if err != nil {
		println(err)
		panic(err)
	}
	return value
}

//SetKeyValue returns a new pool
func SetKeyValue(key string, value string, topic string, conn *radix.Pool) string {
	if conn == nil {
		conn = GetPool()
	}
	// var value string
	err := conn.Do(radix.Cmd(nil, "hset", topic, key, value))
	if err != nil {
		println(err)
		panic(err)
	}
	return value
}

//ClientDoReturnNil makes your client do a command
func ClientDoReturnNil(clnt *radix.Pool, cmd string, on string, val string) {
	clnt.Do(radix.Cmd(nil, cmd, on, val))
}
