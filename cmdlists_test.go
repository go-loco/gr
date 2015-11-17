package gr_test

import (
	"log"
	"testing"

	"github.com/xzip/gr"
)

func TestListsBegin(t *testing.T) {
	log.Println("[Testing Lists]")
}

func TestLPushWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.LPush("gr::mylist"); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestLPush(t *testing.T) {
	test := func() {
		r, err := redis.LPush("gr::mylist", "3", "2")
		if err != nil || r != 2 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestLPushX(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "3", "2")

		r, err := redis.LPushX("gr::mylist", "1")
		if err != nil || r != 3 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestRPushWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.RPush("gr::mylist"); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestRPush(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "3", "2", "1")

		r, err := redis.RPush("gr::mylist", "4", "5")
		if err != nil || r != 5 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestRPushX(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "5", "4", "3", "2", "1")

		r, err := redis.RPushX("gr::mylist", "6")
		if err != nil || r != 6 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestLLen(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "6", "5", "4", "3", "2", "1")

		r, err := redis.LLen("gr::mylist")
		if err != nil || r != 6 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestLIndex(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "6", "5", "4", "3", "2", "1")

		r, err := redis.LIndex("gr::mylist", 2)
		if err != nil || r != "3" {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestLPop(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "6", "5", "4", "3", "2", "1")

		r, err := redis.LPop("gr::mylist")
		if err != nil || r != "1" {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestRPop(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "6", "5", "4", "3", "2")

		r, err := redis.RPop("gr::mylist")
		if err != nil || r != "6" {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestLSet(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "6", "5", "4", "3", "2", "1")

		_, err := redis.LSet("gr::mylist", 0, "10")
		if err != nil {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestLInsertWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.LInsert("gr::mylist", 3, "foo", "foo"); err != gr.ParamErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestLInsert(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "10", "6", "5", "4", "3", "2", "1")

		r, err := redis.LInsert("gr::mylist", gr.Before, "10", "11")
		if err != nil || r == -1 {
			t.Fail()
		}

		r, err = redis.LInsert("gr::mylist", gr.After, "11", "12")
		if err != nil || r == -1 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestRPopLPush(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "11", "10", "6", "5", "4", "3", "2", "1")

		r, err := redis.RPopLPush("gr::mylist", "my_other_list")
		if err != nil || r == "" {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestBRPopLPush(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "11", "10", "6", "5", "4", "3", "2", "1")

		r, err := redis.BRPopLPush("gr::mylist", "my_other_list", 0)
		if err != nil || r == "" {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestBLPopWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.BLPop(0); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestBLPop(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "11", "10", "6", "5", "4", "3", "2", "1")

		r, err := redis.BLPop(0, "gr::mylist")
		if err != nil || len(r) != 2 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestBRPopWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.BRPop(0); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestBRPop(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "11", "10", "6", "5", "4", "3", "2", "1")

		r, err := redis.BRPop(0, "gr::mylist")
		if err != nil || len(r) != 2 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestLRange(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "11", "10", "6", "5", "4", "3", "2", "1")

		r, err := redis.LRange("gr::mylist", 0, -1)
		if err != nil || len(r) <= 0 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestLRem(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "11", "10", "6", "5", "4", "3", "2", "1")

		r, err := redis.LRem("gr::mylist", 0, "10")
		if err != nil || r != 1 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestLTrim(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "11", "10", "6", "5", "4", "3", "2", "1")

		_, err := redis.LTrim("gr::mylist", 0, 2)
		if err != nil {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestListsPipelinedFailed(t *testing.T) {

	safeTestContext(func() {

		var lPush, rPush, lInsert *gr.RespInt
		var bLPop, bRPop *gr.RespStringArray

		//NotEnoughParamsErr
		err := redis.Pipelined(func(p *gr.Pipeline) {
			lPush = p.LPush("gr::mylist")
			rPush = p.RPush("gr::mylist")
			bLPop = p.BLPop(0)
			bRPop = p.BRPop(0)
			lInsert = p.LInsert("gr::mylist", 3, "foo", "foo")
		})

		if err == nil {
			t.Fail()
		}

		if lPush.Error != gr.NotEnoughParamsErr {
			t.Fail()
		}

		if rPush.Error != gr.NotEnoughParamsErr {
			t.Fail()
		}

		if bLPop.Error != gr.NotEnoughParamsErr {
			t.Fail()
		}

		if bRPop.Error != gr.NotEnoughParamsErr {
			t.Fail()
		}

		if lInsert.Error != gr.ParamErr {
			t.Fail()
		}

	})
}

func TestListsPipelined(t *testing.T) {

	safeTestContext(func() {

		var lPush, lPushX, rPush, rPushX, lLen, lInsert, lRem *gr.RespInt
		var lIndex, lPop, rPop, lSet, rPoplPush, brPoplPush, lTrim *gr.RespString
		var blPop, brPop, lRange *gr.RespStringArray

		err := redis.Pipelined(func(p *gr.Pipeline) {
			lPush = p.LPush("gr::mylist", "3", "2")
			lPushX = p.LPushX("gr::mylist", "1")
			rPush = p.RPush("gr::mylist", "4", "5")
			rPushX = p.RPushX("gr::mylist", "6")
			lLen = p.LLen("gr::mylist")
			lIndex = p.LIndex("gr::mylist", 2)
			lPop = p.LPop("gr::mylist")
			rPop = p.RPop("gr::mylist")
			lSet = p.LSet("gr::mylist", 0, "10")
			lInsert = p.LInsert("gr::mylist", gr.Before, "10", "11")
			rPoplPush = p.RPopLPush("gr::mylist", "my_other_list")
			brPoplPush = p.BRPopLPush("gr::mylist", "my_other_list", 0)
			blPop = p.BLPop(0, "gr::mylist")
			brPop = p.BRPop(0, "gr::mylist")
			lRange = p.LRange("gr::mylist", 0, -1)
			lRem = p.LRem("gr::mylist", 0, "10")
			lTrim = p.LTrim("gr::mylist", 0, 2)
		})

		if err != nil {
			t.Fail()
		}

		if lPush.Error != nil || lPush.Value != 2 {
			t.Fail()
		}

		if lPushX.Error != nil || lPushX.Value != 3 {
			t.Fail()
		}

		if rPush.Error != nil || rPush.Value != 5 {
			t.Fail()
		}

		if rPushX.Error != nil || rPushX.Value != 6 {
			t.Fail()
		}

		if lLen.Error != nil || lLen.Value != 6 {
			t.Fail()
		}

		if lPop.Error != nil || lPop.Value != "1" {
			t.Fail()
		}

		if rPop.Error != nil || rPop.Value != "6" {
			t.Fail()
		}

		if lSet.Error != nil {
			t.Fail()
		}

		if lInsert.Error != nil || lInsert.Value == -1 {
			t.Fail()
		}

		if rPoplPush.Error != nil || rPoplPush.Value == "" {
			t.Fail()
		}

		if brPoplPush.Error != nil || brPoplPush.Value == "" {
			t.Fail()
		}

		if blPop.Error != nil || len(blPop.Value) != 2 {
			t.Fail()
		}

		if brPop.Error != nil || len(brPop.Value) != 2 {
			t.Fail()
		}

		if lRange.Error != nil || len(lRange.Value) <= 0 {
			t.Fail()
		}

		if lRem.Error != nil || lRem.Value != 1 {
			t.Fail()
		}

		if lTrim.Error != nil {
			t.Fail()
		}

	})

}

func TestListsEnd(t *testing.T) {
	println("[OK]")
}
