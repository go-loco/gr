package gr

import "strconv"

func rSAdd(key string, values ...string) ([][]byte, error) {
	if len(values) < 1 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile2("SADD", key, values...), nil
}

func rSCard(key string) [][]byte {
	return multiCompile("SCARD", key)
}

func rSDiff(keys ...string) ([][]byte, error) {
	if len(keys) < 1 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile1("SDIFF", keys...), nil
}

func rSDiffStore(key string, keys ...string) ([][]byte, error) {
	if len(keys) < 1 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile2("SDIFFSTORE", key, keys...), nil
}

func rSInter(keys ...string) ([][]byte, error) {
	if len(keys) < 1 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile1("SINTER", keys...), nil
}

func rSInterStore(key string, keys ...string) ([][]byte, error) {
	if len(keys) < 1 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile2("SINTERSTORE", key, keys...), nil
}

func rSIsMember(value string) ([][]byte) {
	return multiCompile("SISMEMBER", value)
}

func rSMembers(key string) ([][]byte) {
	return multiCompile("SMEMBERS", key)
}

func rSMove(sourceKey, destinationKey, value string) ([][]byte) {
	return multiCompile3("SMOVE", sourceKey, destinationKey, value)
}

func rSPop(key string, count int) ([][]byte) {
	return multiCompile2("SPOP", key, strconv.Itoa(count))
}

func rSRandMember(key string, count int) ([][]byte) {
	return multiCompile2("SRANDMEMBER", key, strconv.Itoa(count))
}

func rSRem(key string, values ...string) ([][]byte, error) {
	if len(values) < 1 {
		return nil, NotEnoughParamsErr
	}
	
	return multiCompile2("SREM", key, values...), nil
}

func rSScan(key string, cursor int, scanParams *ScanParams) [][]byte {

	if scanParams == nil {
		return multiCompile("SSCAN", key, strconv.Itoa(cursor))
	}

	return multiCompile3("SSCAN", key, strconv.Itoa(cursor), scanParams.params...)
}

func rSUnion(keys ...string) ([][]byte, error) {
	if len(keys) < 1 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile1("SUNION", keys...), nil
}

func rSUnionStore(key string, keys ...string) ([][]byte, error) {
	if len(keys) < 1 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile2("SUNIONSTORE", key, keys...), nil
}