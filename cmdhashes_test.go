package gr

import (
	"log"
	"testing"
)

func TestHashesBegin(t *testing.T) {
	log.Println("[Testing Hashes]")
}

func TestHSet(t *testing.T) {
	test := func() {
		if r, err := redis.HSet("gr::myhash", "father", "Darth"); err != nil || r != 1 {
			t.Fail()
		}

		if r, err := redis.HSet("gr::myhash", "father", "Darth Vader"); err != nil || r != 0 {
			t.Fail()
		}

		if r, err := redis.HSet("gr::myhash", "son", "Luke Skywalker"); err != nil || r != 1 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
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

	print(".")
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

	print(".")
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

	print(".")
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

	print(".")
}

func TestHExists(t *testing.T) {
	test := func() {
		redis.HSet("gr::myhash", "father", "Darth")

		if r, err := redis.HExists("gr::myhash", "father"); err != nil || !r {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
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

	print(".")
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

	print(".")
}

func TestHMGetWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.HMGet("gr::myhash"); err != NotEnoughParamsErr {
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

	print(".")
}

func TestHMSetWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.HMSet("gr::myhash"); err != NotEnoughParamsErr {
			t.Fail()
		}

		if _, err := redis.HMSet("gr::myhash", "foo"); err != ParamsNotTuplesErr {
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

	print(".")
}

func TestHSetNx(t *testing.T) {
	test := func() {
		r, err := redis.HSetNx("gr::new_hash_key", "one", "1")
		if err != nil || !r {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
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

	print(".")
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
		sp := new(ScanParams).Count(3).Match("father")

		_, rr, err := redis.HScan("gr::myhash", 0, sp)
		if err != nil || len(rr) <= 0 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestHDelWrongParams(t *testing.T) {
	test := func() {
		if _, err := redis.HDel("gr::myhash"); err != NotEnoughParamsErr {
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

	print(".")
}

func TestHashesEnd(t *testing.T) {
	println("[OK]")
}