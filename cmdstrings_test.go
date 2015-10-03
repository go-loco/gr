package gr

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestStringsBegin(t *testing.T) {
	log.Println("[Testing Strings]")
}

func TestSet(t *testing.T) {
	test := func() {
		r, err := redis.Set("gr::father", "Hernan")

		if err != nil || r != "OK" {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestSetNx(t *testing.T) {
	test := func() {
		r, err := redis.SetNx("gr::it doesn't exist", "??")
		if err != nil || !r {
			t.Fail()
		}

		redis.Del("it doesn't exist")
	}

	safeTestContext(test)

	print(".")
}

func TestGet(t *testing.T) {
	test := func() {
		redis.Set("gr::father", "Hernan")

		r, err := redis.Get("gr::father")
		if err != nil || r != "Hernan" {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestSetGet(t *testing.T) {
	test := func() {
		redis.Set("gr::father", "Hernan")

		r, err := redis.GetSet("gr::father", "Hernán Di Chello")
		if err != nil || r != "Hernan" {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestSetX(t *testing.T) {
	test := func() {
		k := KeyExpiration{2, Seconds}
		q := MustNotExist

		r, err := redis.SetX("gr::A-Key-That-Not-Exists", "THE VALUE", &k, &q)

		if err != nil || r != "OK" {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestSetXFail(t *testing.T) {
	test := func() {
		k := KeyExpiration{2, 3}
		q := MustExist

		_, err := redis.SetX("gr::A-Key-That-Not-Exists", "THE VALUE", &k, &q)

		if err != ParamErr {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestGetNil(t *testing.T) {
	test := func() {
		_, err := redis.Get("gr::i am sure this is not a key")

		if err != NilErr {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestIncr(t *testing.T) {
	test := func() {
		_, err := redis.Set("gr::number", "1")

		r, err := redis.Incr("gr::number")
		if err != nil || r != 2 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestIncrBy(t *testing.T) {
	test := func() {
		_, err := redis.Set("gr::number", "2")

		r, err := redis.IncrBy("gr::number", 2)
		if err != nil || r != 4 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestIncrByfloat(t *testing.T) {
	test := func() {
		r, err := redis.IncrByFloat("gr::number:float", 0.5)
		if err != nil || r != 0.5 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestDecr(t *testing.T) {
	test := func() {
		_, err := redis.Set("gr::number", "4")

		r, err := redis.Decr("gr::number")
		if err != nil || r != 3 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestDecrBy(t *testing.T) {
	test := func() {
		_, err := redis.Set("gr::number", "3")

		r, err := redis.DecrBy("gr::number", 2)
		if err != nil || r != 1 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestMSet(t *testing.T) {
	test := func() {
		keyValues := []string{"gr::one", "1", "gr::two", "2", "gr::three", "3"}

		r, err := redis.MSet(keyValues...)
		if err != nil || r != "OK" {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestMSetFail(t *testing.T) {
	test := func() {
		_, err := redis.MSet()
		if err != NotEnoughParamsErr {
			t.Fail()
			fmt.Println(err)
		}

		_, err = redis.MSet("foo")
		if err != ParamsNotTuplesErr {
			t.Fail()
			fmt.Println(err)
		}
	}

	safeTestContext(test)

	print(".")
}

func TestMSetNx(t *testing.T) {
	test := func() {
		r, err := redis.MSetNx("gr::four", "4")
		if err != nil || !r {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestMSetNxFail(t *testing.T) {
	test := func() {
		_, err := redis.MSetNx()
		if err != NotEnoughParamsErr {
			t.Fail()
			fmt.Println(err)
		}

		_, err = redis.MSetNx("foo")
		if err != ParamsNotTuplesErr {
			t.Fail()
			fmt.Println(err)
		}
	}

	safeTestContext(test)

	print(".")
}

func TestMGet(t *testing.T) {
	test := func() {
		redis.MSet([]string{"gr::one", "1", "gr::two", "2", "gr::three", "3"}...)

		r, err := redis.MGet("gr::one", "gr::two", "gr::three")
		if err != nil {
			t.Fail()
		}

		i := 1
		for _, s := range r {
			if s != strconv.Itoa(i) {
				t.Fail()
			}
			i++
		}
	}

	safeTestContext(test)

	print(".")
}

func TestMGetFail(t *testing.T) {
	test := func() {
		_, err := redis.MGet()
		if err != NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestAppend(t *testing.T) {
	test := func() {
		redis.MSet([]string{"gr::one", "1", "gr::two", "2", "gr::three", "3"}...)

		r, err := redis.Append("gr::two", "-dos")
		if err != nil || r != 5 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestBitCount(t *testing.T) {
	test := func() {
		redis.MSet([]string{"gr::one", "1", "gr::two", "2", "gr::three", "3"}...)

		r, err := redis.BitCount("gr::one")
		if err != nil || r != 3 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestBitOp(t *testing.T) {
	test := func() {
		redis.MSet([]string{"gr::one", "1", "gr::two", "2-dos", "gr::three", "3"}...)

		_, err := redis.BitOp(AND, "gr::one", "gr::two")
		if err != nil {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestBitOpFail(t *testing.T) {
	test := func() {
		_, err := redis.BitOp(AND, "gr::one")
		if err != NotEnoughParamsErr {
			t.Fail()
		}

		_, err = redis.BitOp(10, "gr::one", "gr::two")
		if err != ParamErr {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestBitPos(t *testing.T) {
	test := func() {
		redis.MSet([]string{"gr::one", "1", "gr::two", "2-dos", "gr::three", "3"}...)

		_, err := redis.BitPos("gr::one", true)
		if err != nil {
			t.Fail()
		}

		_, err = redis.BitPos("gr::one", true, 0)
		if err != nil {
			t.Fail()
		}

		_, err = redis.BitPos("gr::one", true, 0, -1)
		if err != nil {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestGetBit(t *testing.T) {
	test := func() {
		redis.MSet([]string{"gr::one", "1", "gr::two", "2-dos", "gr::three", "3"}...)

		r, err := redis.GetBit("gr::one", 2)
		if err != nil || r != 1 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestGetRange(t *testing.T) {
	test := func() {
		redis.Set("gr::father", "Hernan")

		r, err := redis.GetRange("gr::father", 0, 2)
		if err != nil || r != "Her" {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestPSetEx(t *testing.T) {
	test := func() {
		r, err := redis.PSetEx("gr::volatile", 100, "ninja")
		if err != nil {
			t.Fail()
		}

		r, err = redis.Get("gr::volatile")
		if err != nil || r != "ninja" {
			t.Fail()
		}

		time.Sleep(200 * time.Millisecond)

		r, err = redis.Get("gr::volatile")
		if err != NilErr {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestSetEx(t *testing.T) {
	test := func() {
		r, err := redis.SetEx("gr::volatile", 1, "ninja")
		if err != nil {
			t.Fail()
		}

		r, err = redis.Get("gr::volatile")
		if err != nil || r != "ninja" {
			t.Fail()
		}

		time.Sleep(1100 * time.Millisecond)

		r, err = redis.Get("gr::volatile")
		if err != NilErr {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestSetBit(t *testing.T) {
	test := func() {
		r, err := redis.SetBit("gr::one", 1, true)
		if err != nil || r != 0 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestSetRange(t *testing.T) {
	test := func() {
		redis.Set("gr::one", "test")
		
		r, err := redis.SetRange("gr::one", 4, "s")
		if err != nil || r != 5 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestStrLen(t *testing.T) {
	test := func() {
		redis.Set("gr::one", "tests")

		r, err := redis.StrLen("gr::one")
		if err != nil || r != 5 {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestStringsPipelined(t *testing.T) {
	test := func() {
		var s1, s2, s3, s4, s5, s6, s7 *RespString
		var i2, i3, i4, i5 *RespInt
		var f *RespFloat
		var b1 *RespBool

		err := redis.Pipelined(func(p *Pipeline) {
			s1 = p.Set("gr::father", "Hernan")
			b1 = p.SetNx("gr::it doesn't exist", "??")

			s2 = p.Get("gr::father")
			s3 = p.GetSet("gr::father", "Hernán Di Chello")

			k := KeyExpiration{2, Seconds}
			q := MustNotExist
			s4 = p.SetX("gr::A-Key-That-Not-Exists", "THE VALUE", &k, &q)

			s5 = p.Get("gr::i am sure this is not a key")
			s6 = p.Set("gr::number", "1")
			i2 = p.Incr("gr::number")
			i3 = p.IncrBy("gr::number", 2)
			f = p.IncrByFloat("gr::number:float", 0.5)
			i4 = p.Decr("gr::number")
			i5 = p.DecrBy("gr::number", 2)

			s7 = p.MSet("gr::one", "1", "gr::two", "2", "gr::three", "3")

			p.MSetNx("gr::four", "4")
			p.MGet("gr::one", "gr::two", "gr::three")
			p.Append("gr::two", "-dos")
			p.BitCount("gr::one")
			p.BitOp(AND, "gr::one", "gr::two")
			p.BitPos("gr::one", true)
			p.BitPos("gr::one", true, 0)
			p.BitPos("gr::one", true, 0, -1)
			p.GetBit("gr::one", 2)
			p.SetBit("gr::one", 1, true)
			p.GetRange("gr::father", 0, 2)
			p.SetRange("gr::one", 0, "2")
			p.StrLen("gr::one")

			p.PSetEx("gr::volatile::ms", 100, "ninja")
			p.SetEx("gr::volatile::s", 1, "ninja")
		})

		if err != nil {
			t.Fail()
		}
	}

	safeTestContext(test)

	print(".")
}

func TestStringsEnd(t *testing.T) {
	println("[OK]")
}