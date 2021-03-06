package gr_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/xzip/gr"
)

func TestSetsBegin(t *testing.T) {
	log.Println("[Testing Sets]")
}

func TestSAddWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.SAdd("gr::myset"); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSAdd(t *testing.T) {
	test := func() {
		r, err := redis.SAdd("gr::myset::sadd", "1", "2")
		if err != nil || r != 2 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSCard(t *testing.T) {
	test := func() {
		redis.SAdd("gr::myset::scard", "1", "2", "3")
		r, err := redis.SCard("gr::myset::scard")
		if err != nil || r != 3 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSDiffWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.SDiff("gr::myset"); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSDiff(t *testing.T) {
	test := func() {
		redis.SAdd("gr::myset::sdiff", "a", "b", "c", "d")
		redis.SAdd("gr::myotherset::sdiff", "a", "c", "d")

		if r, err := redis.SDiff("gr::myset::sdiff", "gr::myotherset::sdiff"); err != nil {
			t.Fail()

		} else {
			if !reflect.DeepEqual(r, []string{"b"}) {
				t.Fail()
			}
		}
	}

	safeTestContext(test)
}

func TestSDiffStoreWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.SDiffStore("gr::myresultset", "gr::myset"); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSDiffStore(t *testing.T) {
	test := func() {
		redis.SAdd("gr::myset::sdiffstore", "a", "b", "c", "d")
		redis.SAdd("gr::myotherset::sdiffstore", "a", "c")

		r, err := redis.SDiffStore("gr::myresultset::sdiffstore", "gr::myset::sdiffstore", "gr::myotherset::sdiffstore")
		if err != nil || r != 2 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSInterWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.SInter("gr::myset"); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSInter(t *testing.T) {
	test := func() {
		redis.SAdd("gr::myset::sinter", "a", "b", "c", "d")
		redis.SAdd("gr::myotherset::sinter", "c")

		if r, err := redis.SInter("gr::myset::sinter", "gr::myotherset::sinter"); err != nil {
			t.Fail()

		} else {
			if !reflect.DeepEqual(r, []string{"c"}) {
				t.Fail()
			}
		}
	}

	safeTestContext(test)
}

func TestSInterStoreWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.SInterStore("gr::myresultset", "gr::myset"); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSInterStore(t *testing.T) {
	test := func() {
		redis.SAdd("gr::myset::sinterstore", "a", "b", "c", "d")
		redis.SAdd("gr::myotherset::sinterstore", "c")

		r, err := redis.SInterStore("gr::myresultset::sinterstore", "gr::myset::sinterstore", "gr::myotherset::sinterstore")
		if err != nil || r != 1 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSIsMember(t *testing.T) {
	test := func() {
		redis.SAdd("gr::myset::sismember", "a", "b", "c", "d")

		r, err := redis.SIsMember("gr::myset::sismember", "a")
		if err != nil || !r {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSMembers(t *testing.T) {
	test := func() {
		redis.SAdd("gr::myset::smembers", "a", "b", "c", "d")

		r, err := redis.SMembers("gr::myset::smembers")
		if err != nil || len(r) != 4 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSMove(t *testing.T) {
	test := func() {
		redis.SAdd("gr::myset::smove", "a", "b")
		redis.SAdd("gr::myotherset::smove", "c", "d")

		r, err := redis.SMove("gr::myset::smove", "gr::myotherset::smove", "a")
		if err != nil || !r {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSPop(t *testing.T) {
	test := func() {
		redis.SAdd("gr::myset::spop", "a", "b", "c", "d")

		r, err := redis.SPop("gr::myset::spop")
		if err != nil || r == "" {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSRandMember(t *testing.T) {
	test := func() {
		redis.SAdd("gr::myset::srandmember", "a", "b", "c", "d")

		r, err := redis.SRandMember("gr::myset::srandmember", 1)
		if err != nil || len(r) != 1 {
			t.Fail()
		}

		r, err = redis.SRandMember("gr::myset::srandmember", 4)
		if err != nil || len(r) != 4 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSRemWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.SRem("gr::myset"); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSRem(t *testing.T) {
	test := func() {
		redis.SAdd("gr::myset::srem", "a", "b", "c", "d")

		r, err := redis.SRem("gr::myset::srem", "c", "d")
		if err != nil || r != 2 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSScan(t *testing.T) {
	test := func() {
		redis.SAdd("gr::myset::sscan", "a", "ab", "bc", "cd")

		_, r, err := redis.SScan("gr::myset::sscan", 0, nil)
		if len(r) == 0 || err != nil {
			t.Fail()
		}

		sp := new(gr.ScanParams).Count(3).Match("a")

		_, rr, err := redis.SScan("gr::myset::sscan", 0, sp)
		if err != nil || len(rr) <= 0 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSUnionWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.SUnion("gr::myset"); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSUnion(t *testing.T) {
	test := func() {
		redis.SAdd("gr::myset::sunion", "a", "b", "c")
		redis.SAdd("gr::myotherset::sunion", "c", "d")

		r, err := redis.SUnion("gr::myset::sunion", "gr::myotherset::sunion")
		if err != nil || len(r) != 4 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSUnionStoreWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.SUnionStore("gr::myresultset", "gr::myset"); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSUnionStore(t *testing.T) {
	test := func() {
		redis.SAdd("gr::myset::sunion", "a", "b", "c")
		redis.SAdd("gr::myotherset::sunion", "c", "d")

		r, err := redis.SUnionStore("gr::myresultset::sunion", "gr::myset::sunion", "gr::myotherset::sunion")
		if err != nil || r != 4 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSetsPipelinedFailed(t *testing.T) {

	var sAdd, sDiffStore, sInterStore, sRem, sUnionStore *gr.RespInt
	var sDiff, sInter, sUnion *gr.RespStringArray

	safeTestContext(func() {
		err := redis.Pipelined(func(p *gr.Pipeline) {
			sAdd = p.SAdd("gr::myset")
			sDiff = p.SDiff("gr::myset")
			sDiffStore = p.SDiffStore("gr::myresultset", "gr::myset")
			sInter = p.SInter("gr::myset")
			sInterStore = p.SInterStore("gr::myresultset", "gr::myset")
			sRem = p.SRem("gr::myset")
			sUnion = p.SUnion("gr::myset")
			sUnionStore = p.SUnionStore("gr::myresultset", "gr::myset")
		})

		if err == nil {
			t.Fail()
		}

		if sAdd.Error != gr.NotEnoughParamsErr {
			t.Fail()
		}

		if sDiff.Error != gr.NotEnoughParamsErr {
			t.Fail()
		}

		if sDiffStore.Error != gr.NotEnoughParamsErr {
			t.Fail()
		}

		if sInter.Error != gr.NotEnoughParamsErr {
			t.Fail()
		}

		if sInterStore.Error != gr.NotEnoughParamsErr {
			t.Fail()
		}

		if sRem.Error != gr.NotEnoughParamsErr {
			t.Fail()
		}

		if sUnion.Error != gr.NotEnoughParamsErr {
			t.Fail()
		}

		if sUnionStore.Error != gr.NotEnoughParamsErr {
			t.Fail()
		}

	})
}

func TestSetsPipelined(t *testing.T) {

	test := func() {
		var sAdd, sCard, sDiffStore, sInterStore, sRem, sUnionStore *gr.RespInt
		var sDiff, sInter, sMembers, sRandMember, sUnion *gr.RespStringArray
		var sIsMember, sMove *gr.RespBool
		var sPop *gr.RespString

		err := redis.Pipelined(func(p *gr.Pipeline) {
			sAdd = p.SAdd("gr::pipeline::myset::sadd", "1", "2")

			p.SAdd("gr::pipeline::myset::scard", "1", "2", "3")
			sCard = p.SCard("gr::pipeline::myset::scard")

			p.SAdd("gr::pipeline::myset::sdiff", "a", "b", "c", "d")
			p.SAdd("gr::pipeline::myotherset::sdiff", "a", "c", "d")
			sDiff = p.SDiff("gr::pipeline::myset::sdiff", "gr::pipeline::myotherset::sdiff")

			p.SAdd("gr::pipeline::myset::sdiffstore", "a", "b", "c", "d")
			p.SAdd("gr::pipeline::myotherset::sdiffstore", "a", "c")
			sDiffStore = p.SDiffStore("gr::pipeline::myresultset::sdiffstore", "gr::pipeline::myset::sdiffstore", "gr::pipeline::myotherset::sdiffstore")

			p.SAdd("gr::pipeline::myset::sinter", "a", "b", "c", "d")
			p.SAdd("gr::pipeline::myotherset::sinter", "c")
			sInter = p.SInter("gr::pipeline::myset::sinter", "gr::pipeline::myotherset::sinter")

			p.SAdd("gr::pipeline::myset::sinterstore", "a", "b", "c", "d")
			p.SAdd("gr::pipeline::myotherset::sinterstore", "c")
			sInterStore = p.SInterStore("gr::pipeline::myresultset::sinterstore", "gr::pipeline::myset::sinterstore", "gr::pipeline::myotherset::sinterstore")

			p.SAdd("gr::pipeline::myset::sismember", "a", "b", "c", "d")
			sIsMember = p.SIsMember("gr::pipeline::myset::sismember", "a")

			p.SAdd("gr::pipeline::myset::smembers", "a", "b", "c", "d")
			sMembers = p.SMembers("gr::pipeline::myset::smembers")

			p.SAdd("gr::pipeline::myset::smove", "a", "b")
			p.SAdd("gr::pipeline::myotherset::smove", "c", "d")
			sMove = p.SMove("gr::pipeline::myset::smove", "gr::pipeline::myotherset::smove", "a")

			p.SAdd("gr::pipeline::myset::spop", "a", "b", "c", "d")
			sPop = p.SPop("gr::pipeline::myset::spop")

			p.SAdd("gr::pipeline::myset::srandmember", "a", "b", "c", "d")
			sRandMember = p.SRandMember("gr::pipeline::myset::srandmember", 4)

			p.SAdd("gr::pipeline::myset::srem", "a", "b", "c", "d")
			sRem = p.SRem("gr::pipeline::myset::srem", "c", "d")

			p.SAdd("gr::pipeline::myset::sunion", "a", "b", "c")
			p.SAdd("gr::pipeline::myotherset::sunion", "c", "d")
			sUnion = p.SUnion("gr::pipeline::myset::sunion", "gr::pipeline::myotherset::sunion")

			p.SAdd("gr::pipeline::myset::sunion", "a", "b", "c")
			p.SAdd("gr::pipeline::myotherset::sunion", "c", "d")
			sUnionStore = p.SUnionStore("gr::pipeline::myresultset::sunion", "gr::pipeline::myset::sunion", "gr::pipeline::myotherset::sunion")
		})

		if err != nil {
			t.Fail()
		}

		if sAdd.Error != nil || sAdd.Value != 2 {
			t.Fail()
		}

		if sCard.Error != nil || sCard.Value != 3 {
			t.Fail()
		}

		if sDiff.Error != nil {
			t.Fail()

		} else {
			if !reflect.DeepEqual(sDiff.Value, []string{"b"}) {
				t.Fail()
			}
		}

		if sDiffStore.Error != nil || sDiffStore.Value != 2 {
			t.Fail()
		}

		if sInter.Error != nil {
			t.Fail()

		} else {
			if !reflect.DeepEqual(sInter.Value, []string{"c"}) {
				t.Fail()
			}
		}

		if sInterStore.Error != nil || sInterStore.Value != 1 {
			t.Fail()
		}

		if sIsMember.Error != nil || !sIsMember.Value {
			t.Fail()
		}

		if sMembers.Error != nil || len(sMembers.Value) != 4 {
			t.Fail()
		}

		if sMove.Error != nil || !sMove.Value {
			t.Fail()
		}

		if sPop.Error != nil || sPop.Value == "" {
			t.Fail()
		}

		if sRandMember.Error != nil || len(sRandMember.Value) != 4 {
			t.Fail()
		}

		if sRem.Error != nil || sRem.Value != 2 {
			t.Fail()
		}

		if sUnion.Error != nil || len(sUnion.Value) != 4 {
			t.Fail()
		}

		if sUnionStore.Error != nil || sUnionStore.Value != 4 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSetsEnd(t *testing.T) {
	println("[OK]")
}
