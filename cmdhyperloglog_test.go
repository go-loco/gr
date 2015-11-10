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

func TestHyperLogLogPipelineFailFirst(t *testing.T) {

	test := func() {

		var pfAdd1 *gr.RespInt
		var pfAdd2 *gr.RespInt
		var pfMerge *gr.RespString
		var pfCount *gr.RespInt

		err := redis.Pipelined(func(p *gr.Pipeline) {
			pfAdd1 = p.PFAdd("gr::hll1")
			pfAdd2 = p.PFAdd("gr::hll2","dad")
			pfMerge = p.PFMerge("gr::all", "gr::hll1", "gr::hll2")
			pfCount = p.PFCount("gr::all")
		})

		if err == nil || len(err)!=1 {
			t.Fail()
		}

		if err[0] != gr.NotEnoughParamsErr {
			t.Fail()
		}

	}

	safeTestContext(test)

}

func TestHyperLogLogPipelineFailLast(t *testing.T) {

	test := func() {

		var pfAdd1 *gr.RespInt
		var pfAdd2 *gr.RespInt
		var pfMerge *gr.RespString
		var pfCount *gr.RespInt

		err := redis.Pipelined(func(p *gr.Pipeline) {
			pfAdd1 = p.PFAdd("gr::hll1", "mom")
			pfAdd2 = p.PFAdd("gr::hll2","dad")
			pfMerge = p.PFMerge("gr::all", "gr::hll1", "gr::hll2")
			pfCount = p.PFCount()
		})

		if err == nil || len(err)!=1 {
			t.Fail()
		}

		if err[0] != gr.NotEnoughParamsErr {
			t.Fail()
		}

	}

	safeTestContext(test)

}

func TestHyperLogLogPipelineFail(t *testing.T) {

	test := func() {

		var pfAdd1 *gr.RespInt
		var pfAdd2 *gr.RespInt
		var pfMerge *gr.RespString
		var pfCount *gr.RespInt

		err := redis.Pipelined(func(p *gr.Pipeline) {
			pfAdd1 = p.PFAdd("gr::hll1")
			pfAdd2 = p.PFAdd("gr::hll2")
			pfMerge = p.PFMerge("gr::all", "gr::hll1", "gr::hll2")
			pfCount = p.PFCount()
		})

		if err == nil || len(err)!=3 {
			t.Fail()
		}

		for _, e := range err[0:2] {
			if e != gr.NotEnoughParamsErr {
				t.Fail()
			}
		}

	}

	safeTestContext(test)

}

func TestHyperLogLogPipeline(t *testing.T) {

	test := func() {

		var pfAdd1, pfAdd2, pfAdd3 *gr.RespInt
		var pfMerge *gr.RespString
		var pfCount1, pfCount2, pfCount3 *gr.RespInt

		err := redis.Pipelined(func(p *gr.Pipeline) {
			pfAdd1 = p.PFAdd("gr::hll1","mom")
			pfAdd2 = p.PFAdd("gr::hll2","dad")
			pfCount1 = p.PFCount("gr::hll1")
			pfCount2 = p.PFCount("gr::hll2")
			pfMerge = p.PFMerge("gr::all", "gr::hll1", "gr::hll2")
			pfAdd3 = p.PFAdd("gr::all","son")
			pfCount3 = p.PFCount("gr::all")
		})

		if err != nil {
			t.Fail()
		}

		if pfAdd1.Error != nil || pfAdd1.Value != 1 {
			t.Fail()
		}

		if pfAdd2.Error != nil || pfAdd2.Value != 1{
			t.Fail()
		}
		
		if pfAdd3.Error != nil || pfAdd3.Value != 1{
			t.Fail()
		}
		
		if pfMerge.Error != nil || pfMerge.Value!="OK" {
			t.Fail()
		}

		if pfCount1.Error != nil || pfCount1.Value < 1 {
			t.Fail()
		}

		if pfCount2.Error != nil || pfCount2.Value < 1 {
			t.Fail()
		}

		if pfCount3.Error != nil || pfCount3.Value < 1 {
			t.Fail()
		}
	}

	safeTestContext(test)

}

func TestHyperloglogEnd(t *testing.T) {
	println("[OK]")
}