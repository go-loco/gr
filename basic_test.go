package gr_test

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/xzip/gr"
)

type testCase func()

var redis *gr.Redis

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
	redis, err = gr.New()

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

func removeKeys() error {
	r1, err := redis.Keys("*")
	if err != nil {
		return err
	}

	if len(r1) > 0 {
		r2, err := redis.Del(r1...)
		if err != nil || int(r2) != len(r1) {
			return errors.New("Unexpected fail in removeKeys method")
		}
	}

	return nil
}

func safeTestContext(fn testCase) {
	fn()

	if err := removeKeys(); err != nil {
		panic(err)
	}
}
