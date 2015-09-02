package gr

import (
	"strconv"
	"testing"

	//gosexy "github.com/gosexy/redis"
)

/////
/////

//var gosexyRedisClient = gosexy.New()
//var _err = gosexyRedisClient.ConnectWithTimeout("localhost", 6379, time.Second*1)

/*
func BenchmarkGosexyRedisPing(b *testing.B) {
	var err error
	gosexyRedisClient.Del("hello")
	for i := 0; i < b.N; i++ {
		_, err = gosexyRedisClient.Ping()
		if err != nil {
			b.Fatalf(err.Error())
			break
		}
	}
}
*/

/*
func BenchmarkGosexyRedisSet(b *testing.B) {
	var err error
	gosexyRedisClient.Del("hello")

	for i := 0; i < b.N; i++ {
		_, err = gosexyRedisClient.Set("hello", 1)
		if err != nil {
			b.Fatalf(err.Error())
			break
		}
	}
}

func BenchmarkGosexyRedisGet(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		_, err = gosexyRedisClient.Get("hello")
		if err != nil {
			b.Fatalf(err.Error())
			break
		}
	}
}


func BenchmarkGosexyRedisIncr(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		_, err = gosexyRedisClient.Incr("hello")
		if err != nil {
			b.Fatalf(err.Error())
			break
		}
	}
}

func BenchmarkGosexyRedisLPush(b *testing.B) {
	var err error
	gosexyRedisClient.Del("hello")

	for i := 0; i < b.N; i++ {
		_, err = gosexyRedisClient.LPush("hello", i)
		if err != nil {
			b.Fatalf(err.Error())
			break
		}
	}
}

func BenchmarkGosexyRedisLRange10(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		_, err = gosexyRedisClient.LRange("hello", 0, 10)
		if err != nil {
			b.Fatalf(err.Error())
			break
		}
	}
}

/////
//////

func BenchmarkGosexyRedisLRange100(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		_, err = gosexyRedisClient.LRange("hello", 0, 100)
		if err != nil {
			b.Fatalf(err.Error())
			break
		}
	}
}

func BenchmarkGosexyRedisPing(b *testing.B) {
	var err error
	redis.Del("hello")
	for i := 0; i < b.N; i++ {
		_, err = gosexyRedisClient.Ping()
		if err != nil {
			b.Fatalf(err.Error())
			break
		}
	}
}
*/

func BenchmarkGrRedisSet(b *testing.B) {
	var err error
	redis.Del("hello")

	for i := 0; i < b.N; i++ {
		_, err = redis.Set("hello", "1")
		if err != nil {
			b.Fatalf(err.Error())
			break
		}
	}

}

func BenchmarkGrRedisGet(b *testing.B) {
	var err error

	for i := 0; i < b.N; i++ {
		_, err = redis.Get("hello")
		if err != nil {
			b.Fatalf(err.Error())
			break
		}
	}

}

func BenchmarkGrRedisIncr(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		_, err = redis.Incr("hello")
		if err != nil {
			b.Fatalf(err.Error())
			break
		}
	}
}

func BenchmarkGrRedisLPush(b *testing.B) {
	var err error
	redis.Del("hello")

	for i := 0; i < b.N; i++ {
		_, err = redis.LPush("hello", strconv.Itoa(i))
		if err != nil {
			b.Fatalf(err.Error())
			break
		}
	}
}

func BenchmarkGrRedisLRange10(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		_, err = redis.LRange("hello", 0, 10)
		if err != nil {
			b.Fatalf(err.Error())
			break
		}
	}
}

func BenchmarkGrRedisLRange100(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		_, err = redis.LRange("hello", 0, 100)
		if err != nil {
			b.Fatalf(err.Error())
			break
		}
	}
}

func BenchmarkPipelineGet(b *testing.B) {

	redis.Pipelined(func(p *Pipeline) {

		for i := 0; i < b.N; i++ {
			p.Get("family:father")
		}

	})

}

/////
//////
