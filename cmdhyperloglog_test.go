package gr_test

import (
	"log"
	"testing"
	"github.com/xzip/gr"
)

func TestHyperloglogBegin(t *testing.T) {
	log.Println("[Testing HyperLogLog]")
}

func TestPFAddWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.PFAdd("gr::hll"); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestPFAdd(t *testing.T) {
	test := func() {
		if changed, err := redis.PFAdd("gr::hll", "father", "son", "number"); err != nil || changed==0{
			t.Fail()
		}

		if changed, err := redis.PFAdd("gr::hll", "mom, grandma"); err != nil || changed==0 {
			t.Fail()
		}

		if changed, err := redis.PFAdd("gr::hll", "son"); err != nil || changed != 0 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestPFCountWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.PFCount(); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestPFCount(t *testing.T) {
	test := func() {
		redis.PFAdd("gr::hll1", "father", "son")

		if count, err := redis.PFCount("gr::hll1"); err != nil || count == 0 {
			t.Fail()
		}

		if count, err := redis.PFCount("gr::hll2"); err != nil || count > 0 {
			t.Fail()
		}

		redis.PFAdd("gr::hll2", "mother", "doughter")

		if count, err := redis.PFCount("gr::hll1", "gr::hll2"); err != nil || count == 0 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestPFMergeWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.PFMerge("gr:desthll"); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestPFMerge(t *testing.T) {
	test := func() {
		redis.PFAdd("gr::hll1", "father", "son")
		redis.PFAdd("gr::hll2", "mom", "son")

		if resp, err := redis.PFMerge("gr::hll", "gr::hll1"); err != nil || resp != "OK" {
			t.Fail()
		}

		if count, err := redis.PFCount("gr::hll"); err != nil || count == 0 {
			t.Fail()
		}

		if resp, err := redis.PFMerge("gr::hll3", "gr::hll", "gr::hll1", "gr::hll2"); err != nil || resp != "OK" {
			t.Fail()
		}

		if count, err := redis.PFCount("gr::hll3"); err != nil || count == 0 {
			t.Fail()
		}

	}

	safeTestContext(test)
}

func TestHyperloglogEnd(t *testing.T) {
	println("[OK]")
}