package gr

import (
	"log"
	"testing"
	"reflect"
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
	r, err := redis.SAdd("gr::myset::sadd", "1", "2")
	if err != nil || r != 2 {
		t.Fail()
	}

	print(".")
}

func TestSCard(t *testing.T) {
	redis.SAdd("gr::myset::scard", "1", "2", "3")
	
	r, err := redis.SCard("gr::myset::scard")
	if err != nil || r != 3 {
		t.Fail()
	}

	print(".")
}

func TestSDiffWrongParams(t *testing.T) {
	if _, err := redis.SDiff("gr::myset"); err != NotEnoughParamsErr {
		t.Fail()
	}

	print(".")
}

func TestSDiff(t *testing.T) {
	redis.SAdd("gr::myset::sdiff", "a", "b", "c", "d")
	redis.SAdd("gr::myotherset::sdiff", "a", "c", "d")

	if r, err := redis.SDiff("gr::myset::sdiff", "gr::myotherset::sdiff"); err != nil {
		t.Fail()
	
	} else {
		if !reflect.DeepEqual(r, []string{"b"}) {
			t.Fail()
		}
	}

	print(".")
}

func TestSDiffStoreWrongParams(t *testing.T) {
	if _, err := redis.SDiffStore("gr::myresultset", "gr::myset"); err != NotEnoughParamsErr {
		t.Fail()
	}

	print(".")
}

func TestSDiffStore(t *testing.T) {
	redis.SAdd("gr::myset::sdiffstore", "a", "b", "c", "d")
	redis.SAdd("gr::myotherset::sdiffstore", "a", "c")

	r, err := redis.SDiffStore("gr::myresultset::sdiffstore", "gr::myset::sdiffstore", "gr::myotherset::sdiffstore")
	if err != nil || r != 2 {
		t.Fail()
	}

	print(".")
}

func TestSInterWrongParams(t *testing.T) {
	if _, err := redis.SInter("gr::myset"); err != NotEnoughParamsErr {
		t.Fail()
	}

	print(".")
}

func TestSInter(t *testing.T) {
	redis.SAdd("gr::myset::sinter", "a", "b", "c", "d")
	redis.SAdd("gr::myotherset::sinter", "c")

	if r, err := redis.SInter("gr::myset::sinter", "gr::myotherset::sinter"); err != nil {
		t.Fail()
	
	} else {
		if !reflect.DeepEqual(r, []string{"c"}) {
			t.Fail()
		}
	}

	print(".")
}

func TestSInterStoreWrongParams(t *testing.T) {
	if _, err := redis.SInterStore("gr::myresultset", "gr::myset"); err != NotEnoughParamsErr {
		t.Fail()
	}

	print(".")
}

func TestSInterStore(t *testing.T) {
	redis.SAdd("gr::myset::sinterstore", "a", "b", "c", "d")
	redis.SAdd("gr::myotherset::sinterstore", "c")

	r, err := redis.SInterStore("gr::myresultset::sinterstore", "gr::myset::sinterstore", "gr::myotherset::sinterstore")
	if err != nil || r != 1 {
		t.Fail()
	}

	print(".")
}

func TestSIsMember(t *testing.T) {
	redis.SAdd("gr::myset::sismember", "a", "b", "c", "d")
	
	r, err := redis.SIsMember("gr::myset::sismember", "a")
	if err != nil || !r {
		t.Fail()
	}

	print(".")
}

func TestSMembers(t *testing.T) {
	redis.SAdd("gr::myset::smembers", "a", "b", "c", "d")
	
	r, err := redis.SMembers("gr::myset::smembers")
	if err != nil || len(r) != 4 {
		t.Fail()
	}

	print(".")
}

func TestSPop(t *testing.T) {
	redis.SAdd("gr::myset::spop", "a", "b", "c", "d")
	
	r, err := redis.SPop("gr::myset::spop")
	if err != nil || r == "" {
		t.Fail()
	}

	print(".")
}

func TestSRandMember(t *testing.T) {
	redis.SAdd("gr::myset::srandmember", "a", "b", "c", "d")
	
	r, err := redis.SRandMember("gr::myset::srandmember", 1)
	if err != nil || len(r) != 1 {
		t.Fail()
	}

	r, err = redis.SRandMember("gr::myset::srandmember", 4)
	if err != nil || len(r) != 4 {
		t.Fail()
	}

	print(".")
}

func TestSRemWrongParams(t *testing.T) {
	if _, err := redis.SRem("gr::myset"); err != NotEnoughParamsErr {
		t.Fail()
	}

	print(".")
}

func TestSRem(t *testing.T) {
	redis.SAdd("gr::myset::srem", "a", "b", "c", "d")
	
	r, err := redis.SRem("gr::myset::srem", "c", "d")
	if err != nil || r != 2 {
		t.Fail()
	}

	print(".")
}

func TestSScan(t *testing.T) {
	redis.SAdd("gr::myset::sscan", "a", "ab", "bc", "cd")

	_, r, err := redis.SScan("gr::myset::sscan", 0, nil)
	if len(r) == 0 || err != nil {
		t.Fail()
	}

	sp := new(ScanParams).Count(3).Match("a")

	_, rr, err := redis.SScan("gr::myset::sscan", 0, sp)
	if err != nil || len(rr) <= 0 {
		t.Fail()
	}

	print(".")
}

func TestSUnionWrongParams(t *testing.T) {
	if _, err := redis.SUnion("gr::myset"); err != NotEnoughParamsErr {
		t.Fail()
	}

	print(".")
}

func TestSUnion(t *testing.T) {
	redis.SAdd("gr::myset::sunion", "a", "b", "c")
	redis.SAdd("gr::myotherset::sunion", "c", "d")

	r, err := redis.SUnion("gr::myset::sunion", "gr::myotherset::sunion");
	if err != nil || len(r) != 4 {
		t.Fail()
	}

	print(".")
}

func TestSUnionStoreWrongParams(t *testing.T) {
	if _, err := redis.SUnionStore("gr::myresultset", "gr::myset"); err != NotEnoughParamsErr {
		t.Fail()
	}

	print(".")
}

func TestSUnionStore(t *testing.T) {
	redis.SAdd("gr::myset::sunion", "a", "b", "c")
	redis.SAdd("gr::myotherset::sunion", "c", "d")

	r, err := redis.SUnionStore("gr::myresultset::sunion", "gr::myset::sunion", "gr::myotherset::sunion")
	if err != nil || r != 4 {
		t.Fail()
	}

	print(".")
}

func TestSetsEnd(t *testing.T) {
	removeKeys(t)
	println("[OK]")
}