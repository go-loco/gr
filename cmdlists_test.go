package gr

import (
	"log"
	"testing"
)

func TestListsBegin(t *testing.T) {
	log.Println("[Testing Lists]")
}

func TestLPushWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.LPush("gr::mylist"); err != NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
	print(".")
}

func TestLPush(t *testing.T) {
	test := func() {
		r, err := redis.LPush("gr::mylist", "3", "2")
		if err != nil || r != 2 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
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

	print(".")
}

func TestRPushWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.RPush("gr::mylist"); err != NotEnoughParamsErr {
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

	print(".")
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

	print(".")
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

	print(".")
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

	print(".")
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

	print(".")
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

	print(".")
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

	print(".")
}

func TestLInsertWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.LInsert("gr::mylist", 3, "foo", "foo"); err != ParamErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestLInsert(t *testing.T) {
	test := func() {
		redis.LPush("gr::mylist", "10", "6", "5", "4", "3", "2", "1")

		r, err := redis.LInsert("gr::mylist", Before, "10", "11")
		if err != nil || r == -1 {
			t.Fail()
		}

		r, err = redis.LInsert("gr::mylist", After, "11", "12")
		if err != nil || r == -1 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
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

	print(".")
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

	print(".")
}

func TestBLPopWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.BLPop(0); err != NotEnoughParamsErr {
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

	print(".")
}

func TestBRPopWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.BRPop(0); err != NotEnoughParamsErr {
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

	print(".")
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

	print(".")
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

	print(".")
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

	print(".")
}

func TestListsEnd(t *testing.T) {
	println("[OK]")
}
