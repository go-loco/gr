package gr

import (
	"log"
	"testing"
)

func TestHashesBegin(t *testing.T) {
	log.Println("[Testing Hashes]")
}

func TestHSet(t *testing.T) {
	if r, err := redis.HSet("gr::myhash", "father", "Darth"); err != nil || r != 1 {
		t.Fail()
	}

	if r, err := redis.HSet("gr::myhash", "father", "Darth Vader"); err != nil || r != 0 {
		t.Fail()
	}

	if r, err := redis.HSet("gr::myhash", "son", "Luke Skywalker"); err != nil || r != 1 {
		t.Fail()
	}

	print(".")
}

func TestHGet(t *testing.T) {
	r, err := redis.HGet("gr::myhash", "father")

	if err != nil || r != "Darth Vader" {
		t.Fail()
	}

	print(".")
}

func TestHGetAll(t *testing.T) {
	r, err := redis.HGetAll("gr::myhash")

	if err != nil || len(r) != 4 {
		t.Fail()
	}

	print(".")
}

func TestHIncrBy(t *testing.T) {
	redis.HSet("gr::myhash", "number", "2")
	r, err := redis.HIncrBy("gr::myhash", "number", 2)

	if err != nil || r != 4 {
		t.Fail()
	}

	print(".")
}

func TestHIncrByFloat(t *testing.T) {
	r, err := redis.HIncrByFloat("gr::myhash", "number", 2.2)

	if err != nil || r != 6.2 {
		t.Fail()
	}

	print(".")
}

func TestHExists(t *testing.T) {

	if r, err := redis.HExists("gr::myhash", "father"); err != nil || !r {
		t.Fail()
	}

	print(".")
}

func TestHKeys(t *testing.T) {

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

	print(".")
}

func TestHLen(t *testing.T) {
	if r, err := redis.HLen("gr::myhash"); err != nil || r != 3 {
		t.Fail()
	}

	print(".")
}

func TestHMGetWrongParams(t *testing.T) {
	if _, err := redis.HMGet("gr::myhash"); err != NotEnoughParamsErr {
		t.Fail()
	}
}

func TestHMGet(t *testing.T) {

	r, err := redis.HMGet("gr::myhash", "father", "son")
	if err != nil || len(r) != 2 {
		t.Fail()
	}

	if r[0] != "Darth Vader" || r[1] != "Luke Skywalker" {
		t.Fail()
	}

	print(".")
}

func TestHMSetWrongParams(t *testing.T) {
	if _, err := redis.HMSet("gr::myhash"); err != NotEnoughParamsErr {
		t.Fail()
	}

	if _, err := redis.HMSet("gr::myhash", "foo"); err != ParamsNotTuplesErr {
		t.Fail()
	}
}

func TestHMSet(t *testing.T) {

	_, err := redis.HMSet("gr::numbers", "one", "1", "two", "2")
	if err != nil {
		t.Fail()
	}
	print(".")
}

func TestHSetNx(t *testing.T) {
	r, err := redis.HSetNx("gr::new_hash_key", "one", "1")
	if err != nil || !r {
		t.Fail()
	}

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
	r, err := redis.HVals("gr::myhash")
	if err != nil || len(r) != 3 {
		t.Fail()
	}

	print(".")
}

func TestHScan(t *testing.T) {

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

	print(".")
}

func TestHDelWrongParams(t *testing.T) {
	if _, err := redis.HDel("gr::myhash"); err != NotEnoughParamsErr {
		t.Fail()
	}
}

func TestHDel(t *testing.T) {

	if _, err := redis.HDel("gr::myhash", "father"); err != nil {
		t.Fail()
	}

	if _, err := redis.HDel("gr::myhash", "son"); err != nil {
		t.Fail()
	}

	if _, err := redis.HDel("gr::myhash", "number"); err != nil {
		t.Fail()
	}

	print(".")
}

func TestRemoveKeysHashes(t *testing.T) {
	removeKeys(t)
}

func TestHashesEnd(t *testing.T) {
	println("[OK]")
}
