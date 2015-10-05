package gr

import "strconv"

func rSAdd(key string, values ...string) ([][]byte, error) {
	if len(values) < 1 {
		return nil, NotEnoughParamsErr
	}

	cmds := append([]string{"SADD", key}, values...)
	return multiCompile(cmds...), nil
}

func rSCard(key string) [][]byte {
	return multiCompile("SCARD", key)
}

func rSDiff(keys ...string) ([][]byte, error) {
	if len(keys) < 2 {
		return nil, NotEnoughParamsErr
	}

	cmds := append([]string{"SDIFF"}, keys...)
	return multiCompile(cmds...), nil
}

func rSDiffStore(key string, keys ...string) ([][]byte, error) {
	if len(keys) < 2 {
		return nil, NotEnoughParamsErr
	}

	cmds := append([]string{"SDIFFSTORE", key}, keys...)
	return multiCompile(cmds...), nil
}

func rSInter(keys ...string) ([][]byte, error) {
	if len(keys) < 2 {
		return nil, NotEnoughParamsErr
	}

	cmds := append([]string{"SINTER"}, keys...)
	return multiCompile(cmds...), nil
}

func rSInterStore(key string, keys ...string) ([][]byte, error) {
	if len(keys) < 2 {
		return nil, NotEnoughParamsErr
	}

	cmds := append([]string{"SINTERSTORE", key}, keys...)
	return multiCompile(cmds...), nil
}

func rSIsMember(key string, value string) [][]byte {
	return multiCompile("SISMEMBER", key, value)
}

func rSMembers(key string) [][]byte {
	return multiCompile("SMEMBERS", key)
}

func rSMove(sourceKey, destinationKey, value string) [][]byte {
	return multiCompile("SMOVE", sourceKey, destinationKey, value)
}

func rSPop(key string) [][]byte {
	return multiCompile("SPOP", key)
}

func rSRandMember(key string, count int) [][]byte {
	return multiCompile("SRANDMEMBER", key, strconv.Itoa(count))
}

func rSRem(key string, values ...string) ([][]byte, error) {
	if len(values) < 1 {
		return nil, NotEnoughParamsErr
	}

	cmds := append([]string{"SREM", key}, values...)
	return multiCompile(cmds...), nil
}

func rSScan(key string, cursor int, scanParams *ScanParams) [][]byte {
	if scanParams == nil {
		return multiCompile("SSCAN", key, strconv.Itoa(cursor))
	}

	cmds := append([]string{"SSCAN", key, strconv.Itoa(cursor)}, scanParams.params...)
	return multiCompile(cmds...)
}

func rSUnion(keys ...string) ([][]byte, error) {
	if len(keys) < 2 {
		return nil, NotEnoughParamsErr
	}

	cmds := append([]string{"SUNION"}, keys...)
	return multiCompile(cmds...), nil
}

func rSUnionStore(key string, keys ...string) ([][]byte, error) {
	if len(keys) < 2 {
		return nil, NotEnoughParamsErr
	}

	cmds := append([]string{"SUNIONSTORE", key}, keys...)
	return multiCompile(cmds...), nil
}