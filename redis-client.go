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

//GetPubSub returns something to listen to pubsub with
func GetPubSub(aesHash string, channel string) chan radix.PubSubMessage {
	conn, err := radix.Dial("tcp", os.Getenv("ERU_SE_REDIS_IP"))
	if err != nil {
		println(err)
		panic(err)
	}
	ps := radix.PubSub(conn)
	defer ps.Close()

	msgCh := make(chan radix.PubSubMessage)
	//
	if err := ps.Subscribe(msgCh, channel); err != nil {
		panic(err)
	}
	return msgCh
}

//WaitForPubSub will wait when invoked
func WaitForPubSub(aesHash string, msgCh chan radix.PubSubMessage) radix.PubSubMessage {
	errCh := make(chan error)
	for {
		select {
		case msg := <-msgCh:
			if string(msg.Message) == string(aesHash) {
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
