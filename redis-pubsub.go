package toolbox

import (
	"log"
	"os"

	"github.com/mediocregopher/radix/v3"
)

//GetPubSubConn returns only the working connection for redis
func GetPubSubConn() radix.PubSubConn {
	conn, err := radix.Dial("tcp", os.Getenv("ERU_SE_REDIS_IP"))
	if err != nil {
		ErrorHandler((err))
	}
	ps := radix.PubSub(conn)
	// defer close(ps)
	return ps
}

// //Publish publishes a message to a certain
// func Publish(channel string, body string, conn radix.PubSubConn) {
// 	if conn == nil {
// 		conn = GetPubSubConn()
// 	}

// 	var msg radix.PubSubMessage
// 	msg.Type = "message"
// 	msg.Channel = channel
// 	msg.Message = []byte(body)
// 	// var message bufio.Reader
// 	conn.Do(radix.Cmd(nil, "publish", msg.Channel))
// 	// err := msg.MarshalRESP()
// 	// if err != nil {
// 	// ErrorHandler(err)
// 	// }
// }

//GetSub returns something to listen to pubsub with
func GetSub(channel string, conn radix.Conn) radix.PubSubMessage {
	results := make(chan radix.PubSubMessage)
	if conn == nil {
		conn = *GetConn()
	}
	ps, err := radix.PersistentPubSubWithOpts(conn)
	if err != nil {
		//
	}
	ps.Subscribe(results, "__keyspace@0__:sentences")
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
