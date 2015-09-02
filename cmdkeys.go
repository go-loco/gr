package gr

import (
	"strconv"
	"time"
)

func rDel(keys ...string) ([][]byte, error) {
	if len(keys) < 1 {
		return nil, NotEnoughParamsErr
	}

	return multiCompile1("DEL", keys...), nil
}

func rDump(key string) [][]byte {
	return multiCompile("DUMP", key)
}

func rExists(key string) [][]byte {
	return multiCompile("EXISTS", key)
}

func rExpire(key string, seconds int) [][]byte {
	return multiCompile("EXPIRE", key, strconv.Itoa(seconds))
}

func rExpireAt(key string, date time.Time) [][]byte {
	return multiCompile("EXPIREAT", key, strconv.FormatInt(date.Unix(), 10))
}

func rKeys(pattern string) [][]byte {
	return multiCompile("KEYS", pattern)
}

func rMigrate(
	host string,
	port int,
	key string,
	destinationDB string,
	timeout int,
	copy bool,
	replace bool) [][]byte {

	cmds := make([]string, 0, 8)
	var a, b []string

	a = []string{"MIGRATE", host, strconv.Itoa(port), key, destinationDB, strconv.Itoa(timeout)}
	cmds = append(cmds, a...)

	switch {
	case copy && replace:
		b = []string{"COPY", "REPLACE"}
	case copy && !replace:
		b = []string{"COPY"}
	case !copy && replace:
		b = []string{"REPLACE"}
	default:
	}

	cmds = append(cmds, b...)
	return multiCompile(cmds...)
}

func rMove(key string, db string) [][]byte {
	return multiCompile("MOVE", key, db)
}

func rObjectEncoding(arguments ...string) [][]byte {
	return multiCompile2("OBJECT", "ENCODING", arguments...)
}

func rObjectRefCount(arguments ...string) [][]byte {
	return multiCompile2("OBJECT", "REFCOUNT", arguments...)
}

func rObjectIdleTime(arguments ...string) [][]byte {
	return multiCompile2("OBJECT", "IDLETIME", arguments...)
}

func rPersist(key string) [][]byte {
	return multiCompile("PERSIST", key)
}

func rPExpire(key string, milliseconds int) [][]byte {
	return multiCompile("PEXPIRE", key, strconv.Itoa(milliseconds))
}

func rPExpireAt(key string, date time.Time) [][]byte {
	return multiCompile("PEXPIREAT", key, strconv.FormatInt(date.Unix(), 10))
}

func rPTTL(key string) [][]byte {
	return multiCompile("PTTL", key)
}

func rRandomKey() [][]byte {
	return multiCompile("RANDOMKEY")
}

func rRename(key string, newKey string) [][]byte {
	return multiCompile("RENAME", key, newKey)
}

func rRenameNx(key string, newKey string) [][]byte {
	return multiCompile("RENAMENX", key, newKey)
}

func rRestore(key string, ttl int, value string, replace bool) [][]byte {

	if replace {
		return multiCompile("RESTORE", key, strconv.Itoa(ttl), value, "REPLACE")
	}

	return multiCompile("RESTORE", key, strconv.Itoa(ttl), value)
}

func rScan(cursor int, scanParams *ScanParams) [][]byte {

	if scanParams == nil {
		return multiCompile("SCAN", strconv.Itoa(cursor))
	}

	m := multiCompile2("SCAN", strconv.Itoa(cursor), scanParams.params...)
	//debugCmds(m)
	return m
}

func rSort(key string, sortingParams *SortParams) [][]byte {

	if sortingParams == nil {
		return multiCompile("SORT", key)
	}

	return multiCompile2("SORT", key, sortingParams.params...)
}

func rSortStore(key string, destination string, sortingParams *SortParams) [][]byte {

	if sortingParams == nil {
		return multiCompile("SORT", key, "STORE", destination)
	}

	return multiCompile4("SORT", key, "STORE", destination, sortingParams.params...)
}

func rTTL(key string) [][]byte {
	return multiCompile("TTL", key)
}

func rType(key string) [][]byte {
	return multiCompile("TYPE", key)
}

func rWait(numSlaves int, timeout int) [][]byte {
	return multiCompile("WAIT", strconv.Itoa(numSlaves), strconv.Itoa(timeout))
}

//APPEND PARAMETERS FOR SCAN AND SORT
func appendParam(params *[]string, pattern ...string) {
	for _, v := range pattern {
		*params = append(*params, v)
	}

}

//SCAN params
type ScanParams struct {
	params []string
}

func (sp *ScanParams) Match(pattern string) *ScanParams {
	appendParam(&sp.params, "MATCH", pattern)
	return sp
}

func (sp *ScanParams) Count(count int) *ScanParams {
	appendParam(&sp.params, "COUNT", strconv.Itoa(count))
	return sp
}

//SORT PARAMS
type SortParams struct {
	params []string
}

func (sp *SortParams) By(pattern string) *SortParams {
	appendParam(&sp.params, "BY", pattern)
	return sp
}

func (sp *SortParams) NoSort() *SortParams {
	appendParam(&sp.params, "BY", "NOSORT")
	return sp
}

func (sp *SortParams) Desc() *SortParams {
	appendParam(&sp.params, "DESC")
	return sp
}

func (sp *SortParams) Asc() *SortParams {
	appendParam(&sp.params, "ASC")
	return sp
}

func (sp *SortParams) Alpha() *SortParams {
	appendParam(&sp.params, "ALPHA")
	return sp
}

func (sp *SortParams) Get(pattern ...string) *SortParams {
	for _, v := range pattern {
		appendParam(&sp.params, "GET", v)
	}

	return sp
}

func (sp *SortParams) Limit(offset int, count int) *SortParams {
	appendParam(&sp.params, "LIMIT", strconv.Itoa(offset), strconv.Itoa(count))
	return sp
}
