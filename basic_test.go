package gr

import (
	"log"
	"testing"
	"errors"
)

type testCase func()

var redis *Redis

func TestMain(m *testing.M) {
    log.Println("Init test")

    setup()

    code := m.Run()

    teardown()

    os.Exit(code)
}

func setup() {
	log.Println("[Testing Connect]")

	var err error
	redis, err = New()

	if err != nil {
		panic(err)
	}

	println(".[OK]")
}

func teardown() {
	if err := removeKeys(); err != nil {
		panic(err)
	}
}

func removeKeys() err {
	r1, err := redis.Keys("*")
	if err != nil {
		return err
	}

	r2, err := redis.Del(r1...)
	if err != nil || int(r2) != len(r1) {
		return errors.New("Unexpected fail in removeKeys method")
	}
}

func safeTestContext(fn testCase) {
	fn()

	if err := removeKeys(); err != nil {
		panic(err)
	}
}