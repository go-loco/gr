package gr_test

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/xzip/gr"
)

func TestStringsBegin(t *testing.T) {
	log.Println("[Testing Strings]")
}

func TestSet(t *testing.T) {
	test := func() {
		r, err := redis.Set("gr::father", "Darth Vader")

		if err != nil || r != "OK" {
			t.Fail()
		}
	}

	safeTestContext(test)
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
}

func TestGet(t *testing.T) {
	test := func() {
		redis.Set("gr::father", "Darth Vader")

		r, err := redis.Get("gr::father")
		if err != nil || r != "Darth Vader" {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSetGet(t *testing.T) {
	test := func() {
		redis.Set("gr::father", "Darth Vader")

		r, err := redis.GetSet("gr::father", "Anakin")
		if err != nil || r != "Darth Vader" {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSetX(t *testing.T) {
	test := func() {
		k := gr.KeyExpiration{2, gr.Seconds}
		q := gr.MustNotExist

		r, err := redis.SetX("gr::A-Key-That-Not-Exists", "THE VALUE", &k, &q)

		if err != nil || r != "OK" {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSetXFail(t *testing.T) {
	test := func() {
		k := gr.KeyExpiration{2, 3}
		q := gr.MustExist

		_, err := redis.SetX("gr::A-Key-That-Not-Exists", "THE VALUE", &k, &q)

		if err != gr.ParamErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestGetNil(t *testing.T) {
	test := func() {
		_, err := redis.Get("gr::i am sure this is not a key")

		if err != gr.NilErr {
			t.Fail()
		}
	}

	safeTestContext(test)
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
}

func TestIncrByfloat(t *testing.T) {
	test := func() {
		r, err := redis.IncrByFloat("gr::number:float", 0.5)
		if err != nil || r != 0.5 {
			t.Fail()
		}
	}

	safeTestContext(test)
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
}

func TestMSetFail(t *testing.T) {
	test := func() {
		_, err := redis.MSet()
		if err != gr.NotEnoughParamsErr {
			t.Fail()
			fmt.Println(err)
		}

		_, err = redis.MSet("foo")
		if err != gr.ParamsNotTuplesErr {
			t.Fail()
			fmt.Println(err)
		}
	}

	safeTestContext(test)
}

func TestMSetNx(t *testing.T) {
	test := func() {
		r, err := redis.MSetNx("gr::four", "4")
		if err != nil || !r {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestMSetNxFail(t *testing.T) {
	test := func() {
		_, err := redis.MSetNx()
		if err != gr.NotEnoughParamsErr {
			t.Fail()
			fmt.Println(err)
		}

		_, err = redis.MSetNx("foo")
		if err != gr.ParamsNotTuplesErr {
			t.Fail()
			fmt.Println(err)
		}
	}

	safeTestContext(test)
}

func TestMGet(t *testing.T) {
	test := func() {

		testCase := []string{"gr::one", "1", "gr::two", "2", "gr::three", "3"}
		testResult := []string{"1", "2", "3"}

		redis.MSet(testCase...)

		r, err := redis.MGet("gr::one", "gr::two", "gr::three")
		if err != nil {
			t.Fail()
		}

		if !reflect.DeepEqual(testResult, r) {
			t.Fail()
		}

	}

	safeTestContext(test)
}

func TestMGetFail(t *testing.T) {
	test := func() {
		_, err := redis.MGet()
		if err != gr.NotEnoughParamsErr {
			t.Fail()
		}
	}

	safeTestContext(test)
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
}

func TestBitOp(t *testing.T) {
	test := func() {
		redis.MSet([]string{"gr::one", "1", "gr::two", "2-dos", "gr::three", "3"}...)

		_, err := redis.BitOp(gr.AND, "gr::one", "gr::two")
		if err != nil {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestBitOpFail(t *testing.T) {
	test := func() {
		_, err := redis.BitOp(gr.AND, "gr::one")
		if err != gr.NotEnoughParamsErr {
			t.Fail()
		}

		_, err = redis.BitOp(10, "gr::one", "gr::two")
		if err != gr.ParamErr {
			t.Fail()
		}
	}

	safeTestContext(test)
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
}

func TestGetRange(t *testing.T) {
	test := func() {
		redis.Set("gr::father", "Darth Vader")

		r, err := redis.GetRange("gr::father", 0, 2)
		if err != nil || r != "Dar" {
			t.Fail()
		}
	}

	safeTestContext(test)
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
		if err != gr.NilErr {
			t.Fail()
		}
	}

	safeTestContext(test)
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
		if err != gr.NilErr {
			t.Fail()
		}
	}

	safeTestContext(test)
}

func TestSetBit(t *testing.T) {
	test := func() {
		r, err := redis.SetBit("gr::one", 1, true)
		if err != nil || r != 0 {
			t.Fail()
		}
	}

	safeTestContext(test)
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
}

func TestStringsPipelinedFailed(t *testing.T) {

	safeTestContext(func() {

		err := redis.Pipelined(func(p *gr.Pipeline) {
			k := gr.KeyExpiration{2, 3}
			q := gr.MustExist
			p.SetX("gr::A-Key-That-Not-Exists", "THE VALUE", &k, &q)

			p.BitOp(10, "gr::one", "gr::two")
		})

		for _, e := range err {
			if e != gr.ParamErr {
				t.Fail()
			}
		}

		err = redis.Pipelined(func(p *gr.Pipeline) {
			p.BitOp(gr.AND, "gr::one")
			p.MGet()
			p.MSet()
			p.MSetNx()
		})

		for _, e := range err {
			if e != gr.NotEnoughParamsErr {
				t.Fail()
			}
		}

	})
}

func TestStringsPipelinedTimeout(t *testing.T) {

	safeTestContext(func() {

		var ninja *gr.RespString

		//PSetEx
		err := redis.Pipelined(func(p *gr.Pipeline) {
			p.PSetEx("gr::volatile::ms", 100, "ninja")
			ninja = p.Get("gr::volatile::ms")
		})

		if err != nil {
			t.Fail()
		}

		if ninja.Error != nil || ninja.Value != "ninja" {
			t.Fail()
		}

		time.Sleep(200 * time.Millisecond)

		_, e := redis.Get("gr::volatile::ms")
		if e != gr.NilErr {
			t.Fail()
		}

		//SetEx
		err = redis.Pipelined(func(p *gr.Pipeline) {
			p.SetEx("gr::volatile", 1, "ninja")
			ninja = p.Get("gr::volatile")
		})

		if err != nil {
			t.Fail()
		}

		if ninja.Error != nil || ninja.Value != "ninja" {
			t.Fail()
		}

		time.Sleep(1100 * time.Millisecond)

		_, e = redis.Get("gr::volatile")
		if e != gr.NilErr {
			t.Fail()
		}

	})
}

func TestStringsPipelined(t *testing.T) {

	safeTestContext(func() {

		var s [8]*gr.RespString
		var i [14]*gr.RespInt
		var f *gr.RespFloat
		var b [2]*gr.RespBool
		var sa *gr.RespStringArray

		err := redis.Pipelined(func(p *gr.Pipeline) {
			s[0] = p.Set("gr::father", "Darth Vader")
			b[0] = p.SetNx("gr::it doesn't exist", "??")

			s[1] = p.Get("gr::father")
			s[2] = p.GetSet("gr::father", "Anakin")

			k := gr.KeyExpiration{2, gr.Seconds}
			q := gr.MustNotExist
			s[3] = p.SetX("gr::A-Key-That-Not-Exists", "THE VALUE", &k, &q)

			s[4] = p.Get("gr::i am sure this is not a key")

			s[5] = p.Set("gr::number", "1")
			i[0] = p.Incr("gr::number")
			i[1] = p.IncrBy("gr::number", 2)
			f = p.IncrByFloat("gr::number:float", 0.5)
			i[2] = p.Decr("gr::number")
			i[3] = p.DecrBy("gr::number", 2)

			s[6] = p.MSet("gr::one", "1", "gr::two", "2", "gr::three", "3")

			b[1] = p.MSetNx("gr::four", "4")
			sa = p.MGet("gr::one", "gr::two", "gr::three")
			i[4] = p.Append("gr::two", "-dos")
			i[5] = p.BitCount("gr::one")
			i[6] = p.BitOp(gr.AND, "gr::one", "gr::two")
			i[7] = p.BitPos("gr::one", true)
			i[8] = p.BitPos("gr::one", true, 0)
			i[9] = p.BitPos("gr::one", true, 0, -1)
			i[10] = p.GetBit("gr::one", 2)
			i[11] = p.SetBit("gr::one", 1, true)

			p.Set("gr::father", "Darth Vader")
			s[7] = p.GetRange("gr::father", 0, 2)

			i[12] = p.SetRange("gr::one", 0, "2")
			i[13] = p.StrLen("gr::one")

		})

		if err != nil {
			t.Fail()
		}

		if s[0].Error != nil || s[0].Value != "OK" {
			t.Fail()
		}

		if b[0].Error != nil || !b[0].Value {
			t.Fail()
		}

		if s[1].Error != nil || s[1].Value != "Darth Vader" {
			t.Fail()
		}

		if s[2].Error != nil || s[2].Value != "Darth Vader" {
			t.Fail()
		}

		if s[3].Error != nil || s[3].Value != "OK" {
			t.Fail()
		}

		if s[4].Error != gr.NilErr {
			t.Fail()
		}

		if s[5].Error != nil || s[5].Value != "OK" {
			t.Fail()
		}

		if i[0].Error != nil || i[0].Value != 2 {
			t.Fail()
		}

		if i[1].Error != nil || i[1].Value != 4 {
			t.Fail()
		}

		if f.Error != nil || f.Value != 0.5 {
			t.Fail()
		}

		if i[2].Error != nil || i[2].Value != 3 {
			t.Fail()
		}

		if i[3].Error != nil || i[3].Value != 1 {
			t.Fail()
		}

		if s[6].Error != nil || s[6].Value != "OK" {
			t.Fail()
		}

		if b[1].Error != nil || !b[1].Value {
			t.Fail()
		}

		if sa.Error != nil || !reflect.DeepEqual([]string{"1", "2", "3"}, sa.Value) {
			t.Fail()
		}

		if i[4].Error != nil || i[4].Value != 5 {
			t.Fail()
		}

		if i[5].Error != nil || i[5].Value != 3 {
			t.Fail()
		}

		if i[6].Error != nil {
			t.Fail()
		}

		if i[7].Error != nil || i[8].Error != nil || i[9].Error != nil {
			t.Fail()
		}

		if i[10].Error != nil || i[10].Value != 1 {
			t.Fail()
		}

		if i[11].Error != nil || i[11].Value != 0 {
			t.Fail()
		}

		if s[7].Error != nil || s[7].Value != "Dar" {
			t.Fail()
		}

		if i[12].Error != nil || i[12].Value != 5 {
			t.Fail()
		}

		if i[13].Error != nil || i[13].Value != 5 {
			t.Fail()
		}

	})
}

func TestStringsEnd(t *testing.T) {
	println("[OK]")
}
