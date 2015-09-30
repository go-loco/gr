package gr

import (
	"bufio"
	"net"
	"time"
)

type RedisChannel struct {
	Channel string
	Data    string
}

type PubSub struct {
	redis  *Redis
	conn   *net.TCPConn
	writer *bufio.Writer
	reader *bufio.Reader

	done         chan bool
	listenerDone chan bool

	Message chan RedisChannel
}

func rPublish(channel string, message string) [][]byte {
	return multiCompile("PUBLISH", channel, message)
}

func newPubSub(redis *Redis) (*PubSub, error) {

	ps := new(PubSub)

	ps.redis = redis
	ps.conn = redis.pool.get()
	ps.writer = bufio.NewWriter(ps.conn)
	ps.reader = bufio.NewReader(ps.conn)
	ps.done = make(chan bool, 1)
	ps.listenerDone = make(chan bool, 1)

	ps.Message = make(chan RedisChannel, 1000)

	return ps, nil
}

func (p *PubSub) subscribe(redisChannels ...string) {
	c := multiCompile(append([]string{"SUBSCRIBE"}, redisChannels...)...)
	p.channelListener(c)
}

func (p *PubSub) pSubscribe(redisChannels ...string) {
	c := multiCompile(append([]string{"PSUBSCRIBE"}, redisChannels...)...)
	p.channelListener(c)
}

func (p *PubSub) channelListener(pattern [][]byte) {

	write(pattern, p.writer)
	p.writer.Flush()

	readStringArray(read(p.reader))

	//SEND MESSAGE FUNCTION OR TIMEOUT
	sendMessage := func(msg []string, i, j int) {

		rc := RedisChannel{
			Channel: msg[i],
			Data:    msg[j],
		}

		select {
		case p.Message <- rc:
		case <-time.After(time.Second * 2):
		}
	}

Loop:
	for {

		//Exit loop if UNSUBSCRIBING
		select {
		case d := <-p.done:
			if d {
				close(p.Message)
				p.listenerDone <- true
				break Loop
			}
		default:
		}

		msg, _ := readStringArray(read(p.reader))

		switch msg[0] {
		case "message":
			sendMessage(msg, 1, 2)

		case "pmessage":
			sendMessage(msg, 2, 3)
		}

	} //End Loop

}

func (p *PubSub) unSubscribe() {

	p.done <- true

	c := multiCompile("UNSUBSCRIBE")
	write(c, p.writer)
	p.writer.Flush()

	//WAIT FOR LISTENER TO CLOSE OR TIMEOUT
	select {
	case <-p.listenerDone:
	case <-time.After(time.Second * 2):
	}

	p.redis.pool.put(p.conn)
	p.writer = nil
	p.reader = nil
}