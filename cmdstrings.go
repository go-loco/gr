package gr

import (
	"math"
	"strconv"
)

//1 - Key Expiration for SET cmd
var keyExpirationArray = []string{"EX", "PX"}

type TimePeriod uint

const (
	Seconds TimePeriod = iota
	MilliSeconds
)

type KeyExpiration struct {
	Time     uint
	TimeType TimePeriod
}

//2 - Key Existance for SET cmd
var keyExistanceArray = []string{"NX", "XX"}

type KeyExistance uint

const (
	MustNotExist KeyExistance = iota
	MustExist
)

//
var bitOperationArray = []string{"AND", "OR", "XOR", "NOT"}

type BitOperation uint

const (
	AND BitOperation = iota
	OR
	XOR
	NOT
)

func rAppend(key string, value string) [][]byte {
	return multiCompile("APPEND", key, value)
}

func rBitCount(key string) [][]byte {
	return multiCompile("BITCOUNT", key)
}

func rBitOp(bitOperation BitOperation, destKey string, keys ...string) ([][]byte, error) {

	if len(keys) < 1 {
		return nil, NotEnoughParamsErr
	}

	if bitOperation > NOT {
		return nil, ParamErr
	}

	return multiCompile3("BITOP", bitOperationArray[bitOperation], destKey, keys...), nil
}

func rBitPos(key string, bit bool, startEnd ...int) [][]byte {

	b := "0"
	if bit {
		b = "1"
	}

	bitParamsLenght := len(startEnd)

	switch bitParamsLenght {
	case 0:
		return multiCompile("BITPOS", key, b)
	case 1:
		return multiCompile("BITPOS", key, b, strconv.Itoa(startEnd[0]))
	default:
		return multiCompile("BITPOS", key, b, strconv.Itoa(startEnd[0]), strconv.Itoa(startEnd[1]))
	}

}

func rGetBit(key string, offset int) [][]byte {
	return multiCompile("GETBIT", key, strconv.Itoa(offset))
}

func rIncr(key string) [][]byte {
	return multiCompile("INCR", key)
}

func rIncrBy(key string, increment int) [][]byte {
	return multiCompile("INCRBY", key, strconv.Itoa(increment))
}

func rIncrByFloat(key string, increment float64) [][]byte {
	return multiCompile("INCRBYFLOAT", key, strconv.FormatFloat(increment, 'f', -1, 64))
}

func rDecr(key string) [][]byte {
	return multiCompile("DECR", key)
}

func rDecrBy(key string, decrement int) [][]byte {
	return multiCompile("DECRBY", key, strconv.Itoa(decrement))
}

func rGet(key string) [][]byte {
	return multiCompile("GET", key)
}

func rGetSet(key string, value string) [][]byte {
	return multiCompile("GETSET", key, value)
}

func rGetRange(key string, start int, end int) [][]byte {
	return multiCompile("GETRANGE", key, strconv.Itoa(start), strconv.Itoa(end))
}

func rMGet(keys ...string) ([][]byte, error) {

	if len(keys) < 1 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile1("MGET", keys...), nil
}

func rSet(key string, value string) [][]byte {
	return multiCompile("SET", key, value)
}

func rSetX(key string, value string, keyExpiration *KeyExpiration, keyExistance *KeyExistance) ([][]byte, error) {

	cmds := make([]string, 0, 6)

	cmds = append(cmds, "SET")
	cmds = append(cmds, key)
	cmds = append(cmds, value)

	if keyExpiration != nil {

		if keyExpiration.TimeType > MilliSeconds {
			return nil, ParamErr
		}

		cmds = append(cmds, keyExpirationArray[keyExpiration.TimeType])
		cmds = append(cmds, strconv.FormatUint(uint64(keyExpiration.Time), 10))
	}

	if keyExistance != nil {
		cmds = append(cmds, keyExistanceArray[*keyExistance])
	}

	return multiCompile(cmds...), nil
}

func rMSet(keyValues ...string) ([][]byte, error) {

	lkv := len(keyValues)

	if lkv < 1 {
		return nil, NotEnoughParamsErr
	} else if math.Mod(float64(lkv), 2) != 0 {
		return nil, ParamsNotTuplesErr
	}

	return multiCompile1("MSET", keyValues...), nil
}

func rMSetNx(keyValues ...string) ([][]byte, error) {

	lkv := len(keyValues)

	if lkv < 1 {
		return nil, NotEnoughParamsErr
	} else if math.Mod(float64(lkv), 2) != 0 {
		return nil, ParamsNotTuplesErr
	}

	return multiCompile1("MSETNX", keyValues...), nil
}

func rPSetEx(key string, milliseconds int, value string) [][]byte {
	return multiCompile("PSETEX", key, strconv.Itoa(milliseconds), value)
}

func rSetEx(key string, seconds int, value string) [][]byte {
	return multiCompile("SETEX", key, strconv.Itoa(seconds), value)
}

func rSetBit(key string, offset int, value bool) [][]byte {

	v := "0"
	if value {
		v = "1"
	}

	return multiCompile("SETBIT", key, strconv.Itoa(offset), v)
}

func rSetNx(key string, value string) [][]byte {
	return multiCompile("SETNX", key, value)
}

func rSetRange(key string, offset int, value string) [][]byte {
	return multiCompile("SETRANGE", key, strconv.Itoa(offset), value)
}

func rStrLen(key string) [][]byte {
	return multiCompile("STRLEN", key)
}
