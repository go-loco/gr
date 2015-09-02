package gr

import (
	"log"
	"testing"
)

func TestListsBegin(t *testing.T) {
	log.Println("[Testing Lists]")
}

func TestLPushWrongParams(t *testing.T) {
	if _, err := redis.LPush("gr::mylist"); err != NotEnoughParamsErr {
		t.Fail()
	}
}

func TestLPush(t *testing.T) {
	r, err := redis.LPush("gr::mylist", "3", "2")
	if err != nil || r != 2 {
		t.Fail()
	}

	print(".")
}

func TestLPushX(t *testing.T) {
	r, err := redis.LPushX("gr::mylist", "1")
	if err != nil || r != 3 {
		t.Fail()
	}

	print(".")
}

func TestRPushWrongParams(t *testing.T) {
	if _, err := redis.RPush("gr::mylist"); err != NotEnoughParamsErr {
		t.Fail()
	}
}

func TestRPush(t *testing.T) {
	r, err := redis.RPush("gr::mylist", "4", "5")
	if err != nil || r != 5 {
		t.Fail()
	}

	print(".")
}

func TestRPushX(t *testing.T) {
	r, err := redis.RPushX("gr::mylist", "6")
	if err != nil || r != 6 {
		t.Fail()
	}

	print(".")
}

func TestLLen(t *testing.T) {
	r, err := redis.LLen("gr::mylist")
	if err != nil || r != 6 {
		t.Fail()
	}

	print(".")
}

func TestLIndex(t *testing.T) {
	r, err := redis.LIndex("gr::mylist", 2)
	if err != nil || r != "3" {
		t.Fail()
	}

	print(".")
}

func TestLPop(t *testing.T) {
	r, err := redis.LPop("gr::mylist")
	if err != nil || r != "1" {
		t.Fail()
	}

	print(".")
}

func TestRPop(t *testing.T) {
	r, err := redis.RPop("gr::mylist")
	if err != nil || r != "6" {
		t.Fail()
	}

	print(".")
}

func TestLSet(t *testing.T) {
	_, err := redis.LSet("gr::mylist", 0, "10")
	if err != nil {
		t.Fail()
	}

	print(".")
}

func TestLInsertWrongParams(t *testing.T) {
	if _, err := redis.LInsert("gr::mylist", 3, "foo", "foo"); err != ParamErr {
		t.Fail()
	}
}

func TestLInsert(t *testing.T) {
	r, err := redis.LInsert("gr::mylist", Before, "10", "11")
	if err != nil || r == -1 {
		t.Fail()
	}

	r, err = redis.LInsert("gr::mylist", After, "11", "12")
	if err != nil || r == -1 {
		t.Fail()
	}

	print(".")
}

func TestRPopLPush(t *testing.T) {
	r, err := redis.RPopLPush("gr::mylist", "my_other_list")
	if err != nil || r == "" {
		t.Fail()
	}

	print(".")
}

func TestBRPopLPush(t *testing.T) {
	r, err := redis.BRPopLPush("gr::mylist", "my_other_list", 0)
	if err != nil || r == "" {
		t.Fail()
	}

	print(".")
}

func TestBLPopWrongParams(t *testing.T) {
	if _, err := redis.BLPop(0); err != NotEnoughParamsErr {
		t.Fail()
	}
}

func TestBLPop(t *testing.T) {
	r, err := redis.BLPop(0, "gr::mylist")
	if err != nil || len(r) != 2 {
		t.Fail()
	}

	print(".")
}

func TestBRPopWrongParams(t *testing.T) {
	if _, err := redis.BRPop(0); err != NotEnoughParamsErr {
		t.Fail()
	}
}

func TestBRPop(t *testing.T) {
	r, err := redis.BRPop(0, "gr::mylist")
	if err != nil || len(r) != 2 {
		t.Fail()
	}

	print(".")
}

func TestLRange(t *testing.T) {
	r, err := redis.LRange("gr::mylist", 0, -1)
	if err != nil || len(r) <= 0 {
		t.Fail()
	}

	print(".")
}

func TestLRem(t *testing.T) {
	r, err := redis.LRem("gr::mylist", 0, "10")
	if err != nil || r != 1 {
		t.Fail()
	}

	print(".")
}

func TestLTrim(t *testing.T) {
	_, err := redis.LTrim("gr::mylist", 0, 2)
	if err != nil {
		t.Fail()
	}

	print(".")
}

func TestListsEnd(t *testing.T) {
	redis.Del("gr::mylist", "my_other_list")
	println("[OK]")
}
