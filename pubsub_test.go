package gr

import (
	"log"
	"strconv"
	"testing"
	"time"
)

func TestPubSubBegin(t *testing.T) {
	log.Println("[Testing PubSub]")
}

func TestPubSub(t *testing.T) {
	test := func() {
		for i := 0; i < 10; i++ {
			go pub(strconv.Itoa(i))
		}
		sub(false)
	}

	safeTestContext(test)

	print(".")
}

func TestPubSubPattern(t *testing.T) {
	test := func() {
		for i := 0; i < 10; i++ {
			go pub(strconv.Itoa(i))
		}
		sub(true)
	}

	safeTestContext(test)

	print(".")
}

func pub(goThread string) {

	time.Sleep(500 * time.Millisecond)

	for i := 0; i < 100; i++ {
		redis.Publish("gr::mychannel", "hello:"+goThread+":"+strconv.Itoa(i))
	}

}

func sub(pattern bool) {

	f := func(ps *PubSub) {
		for i := 0; i < 1000; i++ {
			<-ps.Message
		}
	}

	if pattern {
		redis.PSubscribe(f, "gr::my*")
	} else {
		redis.Subscribe(f, "gr::mychannel")
	}

}

func TestPubSubEnd(t *testing.T) {
	println("[OK]")
}