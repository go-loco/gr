package gr

import (
	"log"
	"testing"
)

func TestSetsBegin(t *testing.T) {
	log.Println("[Testing Sets]")
}

func TestSAddWrongParams(t *testing.T) {
	if _, err := redis.SAdd("gr::myset"); err != NotEnoughParamsErr {
		t.Fail()
	}
}

func TestSAdd(t *testing.T) {
	r, err := redis.SAdd("gr::myset", "3", "2")
	if err != nil || r != 2 {
		t.Fail()
	}

	print(".")
}

func TestSetsEnd(t *testing.T) {
	redis.Del("gr::myset")
	println("[OK]")
}