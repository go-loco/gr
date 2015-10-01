package gr

import (
	"log"
	"testing"
	"time"
)

func TestKeysBegin(t *testing.T) {
	log.Println("[Testing Keys]")
}

func TestDump(t *testing.T) {
	test := func() {
		if _, err := redis.Set("gr::father", "Hernan"); err != nil {
			t.Fail()
		}

		if _, err := redis.Dump("gr::father"); err != nil {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestExists(t *testing.T) {
	test := func() {
		if r, err := redis.Exists("gr::father"); err != nil || !r {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestExpire(t *testing.T) {
	test := func() {
		if _, err := redis.Set("gr::expire", "bla"); err != nil {
			t.Fail()
		}

		if r, err := redis.Expire("gr::expire", 1); err != nil || !r {
			t.Fail()
		}

		time.Sleep(1100 * time.Millisecond)
		if _, err := redis.Get("gr::expire"); err != NilErr {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestExpireAt(t *testing.T) {
	test := func() {
		if _, err := redis.Set("gr::expire", "bla"); err != nil {
			t.Fail()
		}

		td := time.Now().Add(time.Second)
		if r, err := redis.ExpireAt("gr::expire", td); err != nil || !r {
			t.Fail()
		}

		time.Sleep(1100 * time.Millisecond)
		if _, err := redis.Get("gr::expire"); err != NilErr {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestPExpire(t *testing.T) {
	test := func() {
		if _, err := redis.Set("gr::expire", "bla"); err != nil {
			t.Fail()
		}

		if r, err := redis.PExpire("gr::expire", 100); err != nil || !r {
			t.Fail()
		}

		time.Sleep(200 * time.Millisecond)
		if _, err := redis.Get("gr::expire"); err != NilErr {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestPExpireAt(t *testing.T) {
	test := func() {
		if _, err := redis.Set("gr::expire", "bla"); err != nil {
			t.Fail()
		}

		td := time.Now().Add(100 * time.Millisecond)
		if r, err := redis.PExpireAt("gr::expire", td); err != nil || !r {
			t.Fail()
		}

		time.Sleep(200 * time.Millisecond)
		if _, err := redis.Get("gr::expire"); err != NilErr {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestPersist(t *testing.T) {
	test := func() {
		if _, err := redis.Set("gr::not_expire", "bla"); err != nil {
			t.Fail()
		}

		if r, err := redis.PExpire("gr::not_expire", 500); err != nil || !r {
			t.Fail()
		}

		if r, err := redis.Persist("gr::not_expire"); err != nil || !r {
			t.Fail()
		}

		time.Sleep(600 * time.Millisecond)
		if r, err := redis.Get("gr::not_expire"); err != nil || r != "bla" {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestKeys(t *testing.T) {
	test := func() {
		r, err := redis.Keys("gr::*")
		if err != nil || len(r) != 2 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestRandomKey(t *testing.T) {
	test := func() {
		r, err := redis.RandomKey()
		if err != nil || r == "" {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestTTL(t *testing.T) {
	test := func() {
		r, err := redis.TTL("gr::father")
		if err != nil || r != -1 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestPTTL(t *testing.T) {
	test := func() {
		r, err := redis.PTTL("gr::father")
		if err != nil || r != -1 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestRenameNx(t *testing.T) {
	test := func() {
		redis.Set("gr::change_me", "foo")

		r, err := redis.RenameNx("gr::change_me", "gr::changed")
		if err != nil || !r {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestRename(t *testing.T) {
	test := func() {
		r, err := redis.Rename("gr::changed", "gr::changed_2")
		if err != nil || r != "OK" {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestType(t *testing.T) {
	test := func() {
		r, err := redis.Type("gr::father")
		if err != nil || r != "string" {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestDelWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.Del(); err != NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestScan(t *testing.T) {
	test := func() {
		_, r, err := redis.Scan(0, nil)

		if len(r) == 0 || err != nil {
			t.Fail()
		}

		/////
		sp := new(ScanParams).Count(3).Match("gr\\:\\:*")

		_, rr, err := redis.Scan(0, sp)
		if err != nil || len(rr) <= 0 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestSort(t *testing.T) {
	test := func() {
		redis.RPush("gr::mylist", "3", "2", "1")

		r, err := redis.Sort("gr::mylist", nil)
		if err != nil || !(r[0] == "1" && r[1] == "2" && r[2] == "3") {
			t.Fail()
		}

		r, err = redis.Sort("gr::mylist", new(SortParams).NoSort())
		if err != nil || !(r[0] == "3" && r[1] == "2" && r[2] == "1") {
			t.Fail()
		}

		sortParams := new(SortParams).Desc()
		r, err = redis.Sort("gr::mylist", sortParams)
		if err != nil || !(r[0] == "3" && r[1] == "2" && r[2] == "1") {
			t.Fail()
		}

		sortParams = new(SortParams).Limit(0, 1)
		r, err = redis.Sort("gr::mylist", sortParams)
		if err != nil || len(r) != 1 {
			t.Fail()
		}

		sortParams = new(SortParams).By("gr::mylist").Alpha().Asc()
		r, err = redis.Sort("gr::mylist", sortParams)
		if err != nil || !(r[0] == "3" && r[1] == "2" && r[2] == "1") {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestSortStore(t *testing.T) {
	test := func() {
		sortParams := new(SortParams).By("gr::mylist").Alpha().Asc()
		r, err := redis.SortStore("gr::mylist", "gr::resultkey", sortParams)
		if err != nil || r != 3 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestObjectEncoding(t *testing.T) {
	test := func() {
		r, err := redis.Set("gr::object", "object")
		if err != nil || r != "OK" {
			t.Fail()
		}

		rr, err := redis.ObjectEncoding("gr::object")
		if err != nil || rr != "embstr" {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestObjectRefCount(t *testing.T) {
	test := func() {
		r, err := redis.ObjectRefCount("gr::object")
		if err != nil || r != 1 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestObjectIdleTime(t *testing.T) {
	test := func() {
		time.Sleep(1200 * time.Millisecond)
		r, err := redis.ObjectIdleTime("gr::object")

		if err != nil || r != 1 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestRestore(t *testing.T) {
	test := func() {
		dump, err := redis.Dump("gr::father")
		if err != nil {
			t.Fail()
		}

		if r, err := redis.Restore("gr::father", 0, dump, true); err != nil || r != "OK" {
			t.Fail()
		}

		if r, err := redis.Restore("gr::father2", 0, dump, false); err != nil || r != "OK" {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestMove(t *testing.T) {
	test := func() {
		r, err := redis.Set("gr::move_me", "foo")
		if r != "OK" || err != nil {
			t.Fail()
		}

		rr, err := redis.Move("gr::move_me", "1")
		if !rr || err != nil {
			t.Fail()
		}

		redis.Pipelined(func(p *Pipeline) {
			p.Select(1)
			p.Del("gr::move_me")
		})
	}

	safeTestContext(test)

	print(".")
}

func TestWait(t *testing.T) {
	test := func() {
		r, err := redis.Wait(1, 500)
		if err != nil || r != 1 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestMigrate(t *testing.T) {
	test := func() {
		r, err := redis.Migrate("localhost", 7000, "gr::father", "0", 500, true, true)
		if err != nil || r != "OK" {
			t.Fail()
		}

		r, err = redis.Migrate("localhost", 7000, "gr::father", "0", 500, false, true)
		if err != nil || r != "OK" {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestKeysEnd(t *testing.T) {
	println("[OK]")
}