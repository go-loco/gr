package gr

import (
	"log"
	"testing"
)

var redis *Redis

func init() {

	log.Println("[Testing Connect]")

	var err error
	redis, err = New()

	if err != nil {
		panic(err)
	}

	println(".[OK]")
}

func removeKeys(t *testing.T) {
	r1, err := redis.Keys("gr::*")
	if err != nil {
		t.Fail()
	}

	r2, err := redis.Del(r1...)
	if err != nil || int(r2) != len(r1) {
		t.Fail()
	}
}
