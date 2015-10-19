package gr_test

import (
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/xzip/gr"
)

func TestKeysBegin(t *testing.T) {
	log.Println("[Testing Keys]")
}

func TestDump(t *testing.T) {
	test := func() {
		if _, err := redis.Set("gr::father", "Vader"); err != nil {
			t.Fail()
		}

		if _, err := redis.Dump("gr::father"); err != nil {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestExists(t *testing.T) {
	test := func() {
		redis.Set("gr::father", "Vader")

		if r, err := redis.Exists("gr::father"); err != nil || !r {
			t.Fail()
		}
	}

	safeTestContext(test)
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
		if _, err := redis.Get("gr::expire"); err != gr.NilErr {
			t.Fail()
		}
	}

	safeTestContext(test)
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
		if _, err := redis.Get("gr::expire"); err != gr.NilErr {
			t.Fail()
		}
	}

	safeTestContext(test)
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
		if _, err := redis.Get("gr::expire"); err != gr.NilErr {
			t.Fail()
		}
	}

	safeTestContext(test)
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
		if _, err := redis.Get("gr::expire"); err != gr.NilErr {
			t.Fail()
		}
	}

	safeTestContext(test)
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
}

func TestKeys(t *testing.T) {
	test := func() {
		redis.Set("gr::father", "Darth")
		redis.Set("gr::son", "Luke")

		r, err := redis.Keys("gr::*")
		if err != nil || len(r) != 2 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestRandomKey(t *testing.T) {
	test := func() {
		redis.Set("gr::father", "Darth")
		redis.Set("gr::son", "Luke")

		r, err := redis.RandomKey()
		if err != nil || r == "" {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestTTL(t *testing.T) {
	test := func() {
		redis.Set("gr::father", "Darth")

		r, err := redis.TTL("gr::father")
		if err != nil || r != -1 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestPTTL(t *testing.T) {
	test := func() {
		redis.Set("gr::father", "Darth")

		r, err := redis.PTTL("gr::father")
		if err != nil || r != -1 {
			t.Fail()
		}
	}

	safeTestContext(test)
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
}

func TestRename(t *testing.T) {
	test := func() {
		redis.Set("gr::changed", "foo")

		r, err := redis.Rename("gr::changed", "gr::changed_2")
		if err != nil || r != "OK" {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestType(t *testing.T) {
	test := func() {
		redis.Set("gr::father", "Darth")

		r, err := redis.Type("gr::father")
		if err != nil || r != "string" {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestDelWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.Del(); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestScan(t *testing.T) {
	test := func() {
		redis.Set("gr::father", "Darth")
		redis.Set("gr::son", "Luke")
		redis.Set("gr::otherson", "unknown")

		_, r, err := redis.Scan(0, nil)

		if len(r) == 0 || err != nil {
			t.Fail()
		}

		/////
		sp := new(gr.ScanParams).Count(3).Match("gr\\:\\:*")

		_, rr, err := redis.Scan(0, sp)
		if err != nil || len(rr) <= 0 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSort(t *testing.T) {
	test := func() {
		redis.RPush("gr::mylist", "3", "2", "1")

		r, err := redis.Sort("gr::mylist", nil)
		if err != nil || !reflect.DeepEqual(r, []string{"1", "2", "3"}) {
			t.Fail()
		}

		r, err = redis.Sort("gr::mylist", new(gr.SortParams).NoSort())
		if err != nil || !reflect.DeepEqual(r, []string{"3", "2", "1"}) {
			t.Fail()
		}

		sortParams := new(gr.SortParams).Desc()
		r, err = redis.Sort("gr::mylist", sortParams)
		if err != nil || !reflect.DeepEqual(r, []string{"3", "2", "1"}) {
			t.Fail()
		}

		sortParams = new(gr.SortParams).Limit(0, 1)
		r, err = redis.Sort("gr::mylist", sortParams)
		if err != nil || len(r) != 1 {
			t.Fail()
		}

		sortParams = new(gr.SortParams).By("gr::mylist").Alpha().Asc()
		r, err = redis.Sort("gr::mylist", sortParams)
		if err != nil || !reflect.DeepEqual(r, []string{"3", "2", "1"}) {
			t.Fail()
		}

		sortParams = new(gr.SortParams).Get("#")
		r, err = redis.Sort("gr::mylist", sortParams)
		if err != nil || !reflect.DeepEqual(r, []string{"1", "2", "3"}) {
			t.Fail()
		}

	}

	safeTestContext(test)
}

func TestSortStore(t *testing.T) {

	test := func() {
		redis.RPush("gr::mylist", "3", "2", "1")

		r, err := redis.SortStore("gr::mylist", "gr::resultkey", nil)
		if err != nil || r != 3 {
			t.Fail()
		}

		rr, err := redis.LRange("gr::resultkey", 0, -1)
		if err != nil || rr[0] != "1" {
			t.Fail()
		}

		sortParams := new(gr.SortParams).By("gr::mylist").Alpha().Asc()
		r, err = redis.SortStore("gr::mylist", "gr::resultkey", sortParams)
		if err != nil || r != 3 {
			t.Fail()
		}

		rr, err = redis.LRange("gr::resultkey", 0, -1)
		if err != nil || rr[0] != "3" {
			t.Fail()
		}

	}

	safeTestContext(test)
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
}

func TestObjectRefCount(t *testing.T) {
	test := func() {
		redis.Set("gr::object", "object")

		r, err := redis.ObjectRefCount("gr::object")
		if err != nil || r != 1 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestObjectIdleTime(t *testing.T) {
	test := func() {
		redis.Set("gr::object", "object")

		time.Sleep(1100 * time.Millisecond)

		r, err := redis.ObjectIdleTime("gr::object")
		if err != nil || r < 1 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestRestore(t *testing.T) {
	test := func() {
		redis.Set("gr::father", "Darth")

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

		redis.Pipelined(func(p *gr.Pipeline) {
			p.Select(1)
			p.Del("gr::move_me")
		})
	}

	safeTestContext(test)
}

func TestWait(t *testing.T) {
	test := func() {
		r, err := redis.Wait(1, 500)
		if err != nil || r != 1 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestMigrate(t *testing.T) {

	test := func() {

		r2, err2 := gr.NewWithConfig(gr.Config{
			Port:           7000,
			Address:        "localhost",
			MinConnections: 1,
		})

		if err2 != nil {
			t.Fail()
		}

		r2.Del("gr::father")
		redis.Set("gr::father", "Darth")

		r, err := redis.Migrate("localhost", 7000, "gr::father", "0", 500, true, false)
		if err != nil || r != "OK" {
			t.Fail()
		}

		r, err = redis.Migrate("localhost", 7000, "gr::father", "0", 500, true, true)
		if err != nil || r != "OK" {
			t.Fail()
		}

		r2.Del("gr::father")

		r, err = redis.Migrate("localhost", 7000, "gr::father", "0", 500, false, false)
		if err != nil || r != "OK" {
			t.Fail()
		}

		redis.Set("gr::father", "Darth")

		r, err = redis.Migrate("localhost", 7000, "gr::father", "0", 500, false, true)
		if err != nil || r != "OK" {
			t.Fail()
		}

		d, err := r2.Del("gr::father")
		if err != nil || d != 1 {
			t.Fail()
		}

	}

	safeTestContext(test)
}

func TestKeysPipelined(t *testing.T) {

	safeTestContext(func() {

		var dump *gr.RespString

		//setup
		err := redis.Pipelined(func(p *gr.Pipeline) {
			p.Set("gr::expire", "bla")
			p.Expire("gr::expire", 1)

			p.Set("gr::expire_at", "bla")
			td := time.Now().Add(time.Second)
			p.ExpireAt("gr::expire_at", td)

			p.Set("gr::p_expire_at", "bla")
			td = time.Now().Add(100 * time.Millisecond)
			p.PExpireAt("gr::p_expire_at", td)

			p.Set("gr::not_expire", "bla")
			p.PExpire("gr::not_expire", 500)
			p.Persist("gr::not_expire")

			p.Set("gr::change_me_nx", "foo")
			p.Set("gr::change_me", "foo")

			p.RPush("gr::mylist", "3", "2", "1")

			p.Set("gr::object", "object")
			p.Set("gr::move_me", "foo")

			dump = p.Dump("gr::father")
		})

		if err != nil {
			t.Fail()
		}

		time.Sleep(1200 * time.Millisecond)

		var s [10]*gr.RespString
		var b [2]*gr.RespBool
		var i [5]*gr.RespInt
		var sa [7]*gr.RespStringArray

		err = redis.Pipelined(func(p *gr.Pipeline) {

			p.Set("gr::father", "Vader")
			b[0] = p.Exists("gr::father")

			s[0] = p.Get("gr::expire")
			s[1] = p.Get("gr::expire_at")
			s[2] = p.Get("gr::p_expire_at")
			s[3] = p.Get("gr::not_expire")

			sa[0] = p.Keys("gr::*")

			s[4] = p.RandomKey()

			i[0] = p.TTL("gr::father")
			i[1] = p.PTTL("gr::father")

			b[1] = p.RenameNx("gr::change_me_nx", "gr::changed")
			s[5] = p.Rename("gr::change_me", "gr::changed")

			s[6] = p.Type("gr::father")

			///p.Del()

			sa[1] = p.Sort("gr::mylist", nil)

			sa[2] = p.Sort("gr::mylist", new(gr.SortParams).NoSort())

			sortParams := new(gr.SortParams).Desc()
			sa[3] = p.Sort("gr::mylist", sortParams)

			sortParams = new(gr.SortParams).Limit(0, 1)
			sa[4] = p.Sort("gr::mylist", sortParams)

			sortParams = new(gr.SortParams).By("gr::mylist").Alpha().Asc()
			sa[5] = p.Sort("gr::mylist", sortParams)

			sortParams = new(gr.SortParams).Get("#")
			sa[6] = p.Sort("gr::mylist", sortParams)

			sortParams = new(gr.SortParams).By("gr::mylist").Alpha().Asc()
			i[2] = p.SortStore("gr::mylist", "gr::resultkey", sortParams)

			s[7] = p.ObjectEncoding("gr::object")
			i[3] = p.ObjectRefCount("gr::object")
			i[4] = p.ObjectIdleTime("gr::object")

			s[8] = p.Restore("gr::father", 0, dump.Value, true)
			s[9] = p.Restore("gr::father2", 0, dump.Value, false)

			p.Move("gr::move_me", "1")
			p.Select(1)
			p.Del("gr::move_me")
			p.Select(0)

			p.Migrate("localhost", 7000, "gr::father", "0", 500, true, true)
			p.Migrate("localhost", 7000, "gr::father", "0", 500, false, true)
		})

		if err != nil {
			t.Fail()
		}

		if b[0].Error != nil || !b[0].Value {
			t.Fail()
		}

		if s[0].Error != gr.NilErr {
			t.Fail()
		}

		if s[1].Error != gr.NilErr {
			t.Fail()
		}

		if s[2].Error != gr.NilErr {
			t.Fail()
		}

		if s[3].Error != nil || s[3].Value != "bla" {
			t.Fail()
		}

		if s[4].Error != nil || s[4].Value == "" {
			t.Fail()
		}

		if i[0].Error != nil || i[0].Value != -1 {
			t.Fail()
		}

		if i[1].Error != nil || i[1].Value != -1 {
			t.Fail()
		}

		if b[1].Error != nil || !b[1].Value {
			t.Fail()
		}

		if s[5].Error != nil || s[5].Value != "OK" {
			t.Fail()
		}

		if s[6].Error != nil || s[6].Value != "string" {
			t.Fail()
		}

		if sa[1].Error != nil || !reflect.DeepEqual(sa[1].Value, []string{"1", "2", "3"}) {
			t.Fail()
		}

		if sa[2].Error != nil || !reflect.DeepEqual(sa[2].Value, []string{"3", "2", "1"}) {
			t.Fail()
		}

		if sa[3].Error != nil || !reflect.DeepEqual(sa[3].Value, []string{"3", "2", "1"}) {
			t.Fail()
		}

		if sa[4].Error != nil || len(sa[4].Value) != 1 {
			t.Fail()
		}

		if sa[5].Error != nil || !reflect.DeepEqual(sa[5].Value, []string{"3", "2", "1"}) {
			t.Fail()
		}

		if sa[6].Error != nil || !reflect.DeepEqual(sa[6].Value, []string{"1", "2", "3"}) {
			t.Fail()
		}

		if s[7].Error != nil || s[7].Value != "embstr" {
			t.Fail()
		}

		if i[3].Error != nil || i[3].Value != 1 {
			t.Fail()
		}

		if i[4].Error != nil || i[4].Value < 1 {
			t.Fail()
		}

		if s[8].Error != nil || s[8].Value != "OK" {
			t.Fail()
			println(s[8].Value)
		}

		/*if s[9].Error != nil || s[9].Value != "OK" {
			t.Fail()
		}*/

		r2, err2 := gr.NewWithConfig(gr.Config{
			Port:           7000,
			Address:        "localhost",
			MinConnections: 1,
		})

		if err2 != nil {
			t.Fail()
		}

		r2.Del("gr::father")

	})

}

func TestKeysEnd(t *testing.T) {
	println("[OK]")
}
