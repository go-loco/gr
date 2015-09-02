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
	if _, err := redis.Set("gr::father", "Hernan"); err != nil {
		t.Fail()
	}

	if _, err := redis.Dump("gr::father"); err != nil {
		t.Fail()
	}

	print(".")
}

func TestExists(t *testing.T) {

	if r, err := redis.Exists("gr::father"); err != nil || !r {
		t.Fail()
	}

	print(".")
}

func TestExpire(t *testing.T) {

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

	print(".")
}

func TestExpireAt(t *testing.T) {

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

	print(".")
}

func TestPExpire(t *testing.T) {

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

	print(".")
}

func TestPExpireAt(t *testing.T) {

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

	print(".")
}

func TestPersist(t *testing.T) {

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

	print(".")
}

func TestKeys(t *testing.T) {

	r, err := redis.Keys("gr::*")
	if err != nil || len(r) != 2 {
		t.Fail()
	}

	print(".")
}

func TestRandomKey(t *testing.T) {

	r, err := redis.RandomKey()
	if err != nil || r == "" {
		t.Fail()
	}

	print(".")
}

func TestTTL(t *testing.T) {

	r, err := redis.TTL("gr::father")
	if err != nil || r != -1 {
		t.Fail()
	}

	print(".")
}

func TestPTTL(t *testing.T) {

	r, err := redis.PTTL("gr::father")
	if err != nil || r != -1 {
		t.Fail()
	}

	print(".")
}

func TestRenameNx(t *testing.T) {

	redis.Set("gr::change_me", "foo")

	r, err := redis.RenameNx("gr::change_me", "gr::changed")
	if err != nil || !r {
		t.Fail()
	}

	print(".")
}

func TestRename(t *testing.T) {

	r, err := redis.Rename("gr::changed", "gr::changed_2")
	if err != nil || r != "OK" {
		t.Fail()
	}

	print(".")
}

func TestType(t *testing.T) {

	r, err := redis.Type("gr::father")
	if err != nil || r != "string" {
		t.Fail()
	}

	print(".")
}

func TestDelWrongParams(t *testing.T) {
	if _, err := redis.Del(); err != NotEnoughParamsErr {
		t.Fail()
	}
}

func TestScan(t *testing.T) {

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

	print(".")
}

func TestSort(t *testing.T) {

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

	print(".")
}

func TestSortStore(t *testing.T) {

	sortParams := new(SortParams).By("gr::mylist").Alpha().Asc()
	r, err := redis.SortStore("gr::mylist", "gr::resultkey", sortParams)
	if err != nil || r != 3 {
		t.Fail()
	}

	print(".")
}

func TestObjectEncoding(t *testing.T) {

	r, err := redis.Set("gr::object", "object")
	if err != nil || r != "OK" {
		t.Fail()
	}

	rr, err := redis.ObjectEncoding("gr::object")
	if err != nil || rr != "embstr" {
		t.Fail()
	}

	print(".")
}

func TestObjectRefCount(t *testing.T) {

	r, err := redis.ObjectRefCount("gr::object")
	if err != nil || r != 1 {
		t.Fail()
	}

	print(".")
}

func TestObjectIdleTime(t *testing.T) {

	time.Sleep(1200 * time.Millisecond)
	r, err := redis.ObjectIdleTime("gr::object")

	if err != nil || r != 1 {
		t.Fail()
	}

}

func TestRestore(t *testing.T) {

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

func TestMove(t *testing.T) {
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

	print(".")
}

func TestWait(t *testing.T) {
	r, err := redis.Wait(1, 500)
	if err != nil || r != 1 {
		t.Fail()
	}

	print(".")
}

func TestMigrate(t *testing.T) {

	r, err := redis.Migrate("localhost", 7000, "gr::father", "0", 500, true, true)
	if err != nil || r != "OK" {
		t.Fail()
	}

	r, err = redis.Migrate("localhost", 7000, "gr::father", "0", 500, false, true)
	if err != nil || r != "OK" {
		t.Fail()
	}
}

func TestDel(t *testing.T) {
	removeKeys(t)
}

func TestKeysEnd(t *testing.T) {
	println("[OK]")
}
