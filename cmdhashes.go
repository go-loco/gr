package gr

import (
	"math"
	"strconv"
)

func rHDel(key string, fields ...string) ([][]byte, error) {
	if len(fields) < 1 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile(append([]string{"HDEL", key}, fields...)...), nil
}

func rHExists(key string, field string) [][]byte {
	return multiCompile("HEXISTS", key, field)
}

func rHGet(key string, field string) [][]byte {
	return multiCompile("HGET", key, field)
}

func rHGetAll(key string) [][]byte {
	return multiCompile("HGETALL", key)
}

func rHIncrBy(key string, field string, increment int) [][]byte {
	return multiCompile("HINCRBY", key, field, strconv.Itoa(increment))
}

func rHIncrByFloat(key string, field string, increment float64) [][]byte {
	return multiCompile("HINCRBYFLOAT", key, field, strconv.FormatFloat(increment, 'f', -1, 64))
}

func rHKeys(key string) [][]byte {
	return multiCompile("HKEYS", key)
}

func rHLen(key string) [][]byte {
	return multiCompile("HLEN", key)
}

func rHMGet(key string, fields ...string) ([][]byte, error) {
	if len(fields) < 1 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile(append([]string{"HMGET", key}, fields...)...), nil
}

func rHMSet(key string, fieldValues ...string) ([][]byte, error) {
	lkv := len(fieldValues)

	if lkv < 1 {
		return nil, NotEnoughParamsErr
	} else if math.Mod(float64(lkv), 2) != 0 {
		return nil, ParamsNotTuplesErr
	}

	return multiCompile(append([]string{"HMSET", key}, fieldValues...)...), nil
}

func rHSet(key string, field string, value string) [][]byte {
	return multiCompile("HSET", key, field, value)
}

func rHSetNx(key string, field string, value string) [][]byte {
	return multiCompile("HSETNX", key, field, value)
}

/* -> turn on in 3.2.0
func rHStrLen(key string, field string) [][]byte {
	return multiCompile("HSTRLEN", key, field)
}
*/

func rHVals(key string) [][]byte {
	return multiCompile("HVALS", key)
}

func rHScan(key string, cursor int, scanParams *ScanParams) [][]byte {

	if scanParams == nil {
		return multiCompile("HSCAN", key, strconv.Itoa(cursor))
	}

	return multiCompile(append([]string{"HSCAN", key, strconv.Itoa(cursor)}, scanParams.params...)...)
}
