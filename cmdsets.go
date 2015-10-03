package gr

import "strconv"

func rSAdd(key string, values ...string) ([][]byte, error) {
	if len(values) < 1 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile(append([]string{"SADD", key}, values...)...), nil
}

func rSCard(key string) [][]byte {
	return multiCompile("SCARD", key)
}

func rSDiff(keys ...string) ([][]byte, error) {
	if len(keys) < 2 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile(append([]string{"SDIFF"}, keys...)...), nil
}

func rSDiffStore(key string, keys ...string) ([][]byte, error) {
	if len(keys) < 2 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile(append([]string{"SDIFFSTORE", key}, keys...)...), nil
}

func rSInter(keys ...string) ([][]byte, error) {
	if len(keys) < 2 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile(append([]string{"SINTER"}, keys...)...), nil
}

func rSInterStore(key string, keys ...string) ([][]byte, error) {
	if len(keys) < 2 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile(append([]string{"SINTERSTORE", key}, keys...)...), nil
}

func rSIsMember(key string, value string) ([][]byte) {
	return multiCompile("SISMEMBER", key, value)
}

func rSMembers(key string) ([][]byte) {
	return multiCompile("SMEMBERS", key)
}

func rSMove(sourceKey, destinationKey, value string) ([][]byte) {
	return multiCompile("SMOVE", sourceKey, destinationKey, value)
}

func rSPop(key string) ([][]byte) {
	return multiCompile("SPOP", key)
}

func rSRandMember(key string, count int) ([][]byte) {
	return multiCompile("SRANDMEMBER", key, strconv.Itoa(count))
}

func rSRem(key string, values ...string) ([][]byte, error) {
	if len(values) < 1 {
		return nil, NotEnoughParamsErr
	}
	
	return multiCompile(append([]string{"SREM", key}, values...)...), nil
}

func rSScan(key string, cursor int, scanParams *ScanParams) [][]byte {
	if scanParams == nil {
		return multiCompile("SSCAN", key, strconv.Itoa(cursor))
	}

	return multiCompile(append([]string{"SSCAN", key, strconv.Itoa(cursor)}, scanParams.params...)...)
}

func rSUnion(keys ...string) ([][]byte, error) {
	if len(keys) < 2 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile(append([]string{"SUNION"}, keys...)...), nil
}

func rSUnionStore(key string, keys ...string) ([][]byte, error) {
	if len(keys) < 2 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile(append([]string{"SUNIONSTORE", key}, keys...)...), nil
}