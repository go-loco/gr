package gr_test

import (
	"log"
	"testing"

	"github.com/xzip/gr"
)

func TestHashesBegin(t *testing.T) {
	log.Println("[Testing Hashes]")
}

func TestHSet(t *testing.T) {
	test := func() {
		if r, err := redis.HSet("gr::myhash", "father", "Darth"); err != nil || !r {
			t.Fail()
		}

		if r, err := redis.HSet("gr::myhash", "father", "Darth Vader"); err != nil || r {
			t.Fail()
		}

		if r, err := redis.HSet("gr::myhash", "son", "Luke Skywalker"); err != nil || !r {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestHGet(t *testing.T) {
	test := func() {
		redis.HSet("gr::myhash", "father", "Darth Vader")

		r, err := redis.HGet("gr::myhash", "father")
		if err != nil || r != "Darth Vader" {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestHGetAll(t *testing.T) {
	test := func() {
		redis.HSet("gr::myhash", "father", "Darth")
		redis.HSet("gr::myhash", "father", "Darth Vader")
		redis.HSet("gr::myhash", "son", "Luke")
		redis.HSet("gr::myhash", "son", "Luke Skywalker")

		r, err := redis.HGetAll("gr::myhash")
		if err != nil || len(r) != 4 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestHIncrBy(t *testing.T) {
	test := func() {
		redis.HSet("gr::myhash", "number", "2")

		r, err := redis.HIncrBy("gr::myhash", "number", 2)
		if err != nil || r != 4 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestHIncrByFloat(t *testing.T) {
	test := func() {
		redis.HSet("gr::myhash", "number", "4")

		r, err := redis.HIncrByFloat("gr::myhash", "number", 2.2)
		if err != nil || r != 6.2 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestHExists(t *testing.T) {
	test := func() {
		redis.HSet("gr::myhash", "father", "Darth")

		if r, err := redis.HExists("gr::myhash", "father"); err != nil || !r {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestHKeys(t *testing.T) {
	test := func() {
		redis.HSet("gr::myhash", "father", "Darth")
		redis.HSet("gr::myhash", "son", "Luke")
		redis.HSet("gr::myhash", "number", "4")

		r, err := redis.HKeys("gr::myhash")
		if err != nil {
			t.Fail()
		}

		fields := map[string]bool{"father": true, "son": true, "number": true}
		counter := 0

		for _, f := range r {
			counter++

			if !fields[f] {
				t.Fail()
			}

		}

		if counter != 3 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestHLen(t *testing.T) {
	test := func() {
		redis.HSet("gr::myhash", "father", "Darth")
		redis.HSet("gr::myhash", "son", "Luke")
		redis.HSet("gr::myhash", "number", "4")

		if r, err := redis.HLen("gr::myhash"); err != nil || r != 3 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestHMGetWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.HMGet("gr::myhash"); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestHMGet(t *testing.T) {
	test := func() {
		redis.HSet("gr::myhash", "father", "Darth")
		redis.HSet("gr::myhash", "son", "Luke")

		r, err := redis.HMGet("gr::myhash", "father", "son")
		if err != nil || len(r) != 2 {
			t.Fail()
		}

		if r[0] != "Darth" || r[1] != "Luke" {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestHMSetWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.HMSet("gr::myhash"); err != gr.NotEnoughParamsErr {
			t.Fail()
		}

		if _, err := redis.HMSet("gr::myhash", "foo"); err != gr.ParamsNotTuplesErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestHMSet(t *testing.T) {
	test := func() {
		_, err := redis.HMSet("gr::numbers", "one", "1", "two", "2")
		if err != nil {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestHSetNx(t *testing.T) {
	test := func() {
		r, err := redis.HSetNx("gr::new_hash_key", "one", "1")
		if err != nil || !r {
			t.Fail()
		}
	}

	safeTestContext(test)
}

//func TestHStrLen(t *testing.T) {
//	r, err := redis.HStrLen("gr::myhash", "father")
//	if err != nil || r != 11 {
//		t.Fail()
//	}
//
//	print(".")
//}

func TestHVals(t *testing.T) {
	test := func() {
		redis.HSet("gr::myhash", "father", "Darth")
		redis.HSet("gr::myhash", "son", "Luke")
		redis.HSet("gr::myhash", "number", "4")

		r, err := redis.HVals("gr::myhash")
		if err != nil || len(r) != 3 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestHScan(t *testing.T) {
	test := func() {
		redis.HSet("gr::myhash", "father", "Darth")
		redis.HSet("gr::myhash", "son", "Luke")
		redis.HSet("gr::myhash", "number", "4")

		_, r, err := redis.HScan("gr::myhash", 0, nil)

		if len(r) == 0 || err != nil {
			t.Fail()
		}

		/////
		sp := new(gr.ScanParams).Count(3).Match("father")

		_, rr, err := redis.HScan("gr::myhash", 0, sp)
		if err != nil || len(rr) <= 0 {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestHDelWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.HDel("gr::myhash"); err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestHDel(t *testing.T) {
	test := func() {
		redis.HSet("gr::myhash", "father", "Darth")
		redis.HSet("gr::myhash", "son", "Luke")
		redis.HSet("gr::myhash", "number", "4")

		if _, err := redis.HDel("gr::myhash", "father"); err != nil {
			t.Fail()
		}

		if _, err := redis.HDel("gr::myhash", "son"); err != nil {
			t.Fail()
		}

		if _, err := redis.HDel("gr::myhash", "number"); err != nil {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestHashesPipelinedFail(t *testing.T) {

	safeTestContext(func() {

		var hDel *gr.RespInt
		var hMGet *gr.RespStringArray
		var hMSet, hMSet2 *gr.RespString

		err := redis.Pipelined(func(p *gr.Pipeline) {
			hDel = p.HDel("gr::myhash")
			hMGet = p.HMGet("gr::myhash")
			hMSet = p.HMSet("gr::myhash")
			hMSet2 = p.HMSet("gr::myhash", "foo")
		})

		if err == nil {
			t.Fail()
		}

		for _, e := range err[0:2] {
			if e != gr.NotEnoughParamsErr {
				t.Fail()
			}
		}

		if err[3] != gr.ParamsNotTuplesErr {
			t.Fail()
		}

	})

}

func TestHashesPipelined(t *testing.T) {

	safeTestContext(func() {

		var hIncr *gr.RespInt
		var hIncrFloat *gr.RespFloat
		var hSet, hSet2, hSet3, hSetNum, hSetNum2, hSetNx *gr.RespBool
		var hGet, hMSet *gr.RespString
		var hGetAll *gr.RespStringArray

		var hExists *gr.RespBool
		var hKeys, hMGet, hVals *gr.RespStringArray
		var hLen, hDel *gr.RespInt

		err := redis.Pipelined(func(p *gr.Pipeline) {
			hSet = p.HSet("gr::myhash", "father", "Darth")

			hSet2 = p.HSet("gr::myhash", "father", "Darth Vader")

			hSet3 = p.HSet("gr::myhash", "son", "Luke Skywalker")
			hGet = p.HGet("gr::myhash", "father")
			hGetAll = p.HGetAll("gr::myhash")

			hSetNum = p.HSet("gr::myhash", "number", "2")
			hIncr = p.HIncrBy("gr::myhash", "number", 2)

			hSetNum2 = p.HSet("gr::myhash", "number", "4")
			hIncrFloat = p.HIncrByFloat("gr::myhash", "number", 2.2)

			hExists = p.HExists("gr::myhash", "father")

			hKeys = p.HKeys("gr::myhash")
			hLen = p.HLen("gr::myhash")
			hMGet = p.HMGet("gr::myhash", "father", "son")

			hMSet = p.HMSet("gr::numbers", "one", "1", "two", "2")
			hSetNx = p.HSetNx("gr::new_hash_key", "one", "1")

			hVals = p.HVals("gr::myhash")

			hDel = p.HDel("gr::myhash", "father")
		})

		if err != nil {
			t.Fail()
		}

		if hSet.Error != nil || !hSet.Value {
			t.Fail()
		}

		if hSet2.Error != nil || hSet2.Value {
			t.Fail()
		}

		if hSet3.Error != nil || !hSet3.Value {
			t.Fail()
		}

		if hSetNum.Error != nil || !hSetNum.Value {
			t.Fail()
		}

		if hSetNum2.Error != nil || hSetNum2.Value {
			t.Fail()
		}

		if hExists.Error != nil || !hExists.Value {
			t.Fail()
		}

		if hKeys.Error != nil || len(hKeys.Value) < 1 {
			t.Fail()
		}

		if hLen.Error != nil || hLen.Value < 1 {
			t.Fail()
		}

		if hMGet.Error != nil || len(hMGet.Value) < 1 {
			t.Fail()
		}

		if hMSet.Error != nil || hMSet.Value != "OK" {
			t.Fail()
		}

		if hSetNx.Error != nil || !hSetNx.Value {
			t.Fail()
		}

		if hVals.Error != nil || len(hVals.Value) < 1 {
			t.Fail()
		}

		if hDel.Error != nil || hDel.Value != 1 {
			t.Fail()
		}

	})
}

func TestHashesEnd(t *testing.T) {
	println("[OK]")
}
