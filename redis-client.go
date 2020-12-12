package toolbox

import (
	"log"
	"os"

	"github.com/mediocregopher/radix/v3"
)

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

//GetPubSubConn returns only the working connection for redis
func GetPubSubConn() radix.Conn {
	conn, err := radix.Dial("tcp", os.Getenv("ERU_SE_REDIS_IP"))
	if err != nil {
		//stuff
	}
	return conn
}

//Publish publishes a message to a certain
func Publish(channel string, body string) {
	conn := GetPubSubConn()
	ps := radix.PubSub(conn)
	defer ps.Close()
	conn.Do(radix.Cmd("publish", channel, body))
}

//GetSub returns something to listen to pubsub with
func GetSub(channel string) chan radix.PubSubMessage {
	conn := GetPubSubConn()
	ps := radix.PubSub(conn)
	defer ps.Close()

	msgCh := make(chan radix.PubSubMessage)
	//
	ps.Subscribe(msgCh, channel)
	// if err := ; err != nil {
	// 	panic(err)
	// }
	return msgCh
}

//WaitForPubSub will wait when invoked
func WaitForPubSub(msgCh chan radix.PubSubMessage) radix.PubSubMessage {
	errCh := make(chan error, 1)

	// for {
	select {
	case msg := <-msgCh:
		return msg
	case err := <-errCh:
		panic(err)
	}
	// }
}

//WaitForHash will wait when invoked
func WaitForHash(aesHash string, msgCh chan radix.PubSubMessage) radix.PubSubMessage {
	errCh := make(chan error)
	for {
		select {
		case msg := <-msgCh:
			if string(AsSha256(msg.Message)) == string(aesHash) {
				return msg
			}
			log.Printf("publish to channel %q received: %q", msg.Channel, msg.Message)
		case err := <-errCh:
			panic(err)
		}
	}
}

//ClientDoReturnNil makes your client do a command
func ClientDoReturnNil(clnt *radix.Pool, cmd string, on string, val string) {
	clnt.Do(radix.Cmd(nil, cmd, on, val))
}
