package gr

import "strconv"

func rSelect(index uint) [][]byte {
	return multiCompile("SELECT", strconv.FormatUint(uint64(index), 10))
}
