package gr

import "strconv"

type InsertLocation uint

const (
	Before InsertLocation = iota
	After
)

func rLLen(key string) [][]byte {
	return multiCompile("LLEN", key)
}

func rLIndex(key string, index int) [][]byte {
	return multiCompile("LINDEX", key, strconv.Itoa(index))
}

func rLPush(key string, values ...string) ([][]byte, error) {

	if len(values) < 1 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile(append([]string{"LPUSH", key}, values...)...), nil
}

func rLPushX(key string, value string) [][]byte {
	return multiCompile("LPUSHX", key, value)
}

func rLPop(key string) [][]byte {
	return multiCompile("LPOP", key)
}

func rRPush(key string, values ...string) ([][]byte, error) {
	if len(values) < 1 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile(append([]string{"RPUSH", key}, values...)...), nil
}

func rRPushX(key string, value string) [][]byte {
	return multiCompile("RPUSHX", key, value)
}

func rRPop(key string) [][]byte {
	return multiCompile("RPOP", key)
}

func rLSet(key string, index int, value string) [][]byte {
	return multiCompile("LSET", key, strconv.Itoa(index), value)
}

func rLInsert(key string, location InsertLocation, pivot string, value string) ([][]byte, error) {

	switch location {
	case Before:
		return multiCompile("LINSERT", key, "BEFORE", pivot, value), nil
	case After:
		return multiCompile("LINSERT", key, "AFTER", pivot, value), nil
	default:
		return nil, ParamErr
	}

}

func rLRange(key string, start int, stop int) [][]byte {
	return multiCompile("LRANGE", key, strconv.Itoa(start), strconv.Itoa(stop))
}

func rLRem(key string, count int, value string) [][]byte {
	return multiCompile("LREM", key, strconv.Itoa(count), value)
}

func rLTrim(key string, start int, stop int) [][]byte {
	return multiCompile("LTRIM", key, strconv.Itoa(start), strconv.Itoa(stop))
}

func rRPopLPush(source string, destination string) [][]byte {
	return multiCompile("RPOPLPUSH", source, destination)
}

func rBLPop(timeout int, keys ...string) ([][]byte, error) {

	if len(keys) < 1 {
		return nil, NotEnoughParamsErr
	}

	strT := strconv.Itoa(timeout)

	values := make([]string, len(keys)+len(strT))
	values = append(values, keys...)
	values = append(values, strT)

	return multiCompile(append([]string{"BLPOP"}, values...)...), nil
}

func rBRPop(timeout int, keys ...string) ([][]byte, error) {

	if len(keys) < 1 {
		return nil, NotEnoughParamsErr
	}

	strT := strconv.Itoa(timeout)

	values := make([]string, len(keys)+len(strT))
	values = append(values, keys...)
	values = append(values, strT)

	return multiCompile(append([]string{"BRPOP"}, values...)...), nil
}

func rBRPopLPush(source string, destination string, timeout int) [][]byte {
	return multiCompile("BRPOPLPUSH", source, destination, strconv.Itoa(timeout))
}
