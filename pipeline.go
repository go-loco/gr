package gr

import (
	"bufio"
	"net"
	"time"
)

type Pipeline struct {
	cmdsSize  int
	cmdsQueue queue
	respQueue queue
	redis     *Redis
	err       error
}

func newPipeline(r *Redis) *Pipeline {
	return &Pipeline{
		cmdsQueue: queue{},
		respQueue: queue{},
		redis:     r,
	}
}

////////
//KEYS
///////

func (p *Pipeline) Del(keys ...string) *RespInt {
	return p.enqueueIntErr(rDel(keys...))
}

func (p *Pipeline) Dump(key string) *RespString {
	return p.enqueueStr(rDump(key))
}

func (p *Pipeline) Exists(key string) *RespBool {
	return p.enqueueBool(rExists(key))
}

func (p *Pipeline) Expire(key string, seconds int) *RespBool {
	return p.enqueueBool(rExpire(key, seconds))
}

func (p *Pipeline) ExpireAt(key string, date time.Time) *RespBool {
	return p.enqueueBool(rExpireAt(key, date))
}

func (p *Pipeline) Keys(pattern string) *RespStringArray {
	return p.enqueueStrArray(rKeys(pattern))
}

func (p *Pipeline) Migrate(
	host string,
	port int,
	key string,
	destinationDB string,
	timeout int,
	copy bool,
	replace bool) *RespString {

	return p.enqueueStr(rMigrate(host, port, key, destinationDB, timeout, copy, replace))
}

func (p *Pipeline) Move(key string, db string) *RespBool {
	return p.enqueueBool(rMove(key, db))
}

func (p *Pipeline) ObjectEncoding(arguments ...string) *RespString {
	return p.enqueueStr(rObjectEncoding(arguments...))
}

func (p *Pipeline) ObjectRefCount(arguments ...string) *RespInt {
	return p.enqueueInt(rObjectRefCount(arguments...))
}

func (p *Pipeline) ObjectIdleTime(arguments ...string) *RespInt {
	return p.enqueueInt(rObjectIdleTime(arguments...))
}

func (p *Pipeline) Persist(key string) *RespBool {
	return p.enqueueBool(rPersist(key))
}

func (p *Pipeline) PExpire(key string, milliseconds int) *RespBool {
	return p.enqueueBool(rPExpire(key, milliseconds))
}

func (p *Pipeline) PExpireAt(key string, date time.Time) *RespBool {
	return p.enqueueBool(rPExpireAt(key, date))
}

func (p *Pipeline) PTTL(key string) *RespInt {
	return p.enqueueInt(rPTTL(key))
}

func (p *Pipeline) RandomKey() *RespString {
	return p.enqueueStr(rRandomKey())
}

func (p *Pipeline) Rename(key string, newKey string) *RespString {
	return p.enqueueStr(rRename(key, newKey))
}

func (p *Pipeline) RenameNx(key string, newKey string) *RespBool {
	return p.enqueueBool(rRenameNx(key, newKey))
}

func (p *Pipeline) Restore(key string, ttl int, value string, replace bool) *RespString {
	return p.enqueueStr(rRestore(key, ttl, value, replace))
}

/*
func (p *Pipeline) Scan(cursor int, scanParams *ScanParams) (int64, []string, error) {
	s, err := r.writeReadStrArray(rScan(cursor, scanParams))
	c, sc, err := scanGeneric(s, err)

	return c, sc, err
}
*/

func (p *Pipeline) Sort(key string, sortParams *SortParams) *RespStringArray {
	return p.enqueueStrArray(rSort(key, sortParams))
}

func (p *Pipeline) SortStore(key string, destination string, sortParams *SortParams) *RespInt {
	return p.enqueueInt(rSortStore(key, destination, sortParams))
}

func (p *Pipeline) TTL(key string) *RespInt {
	return p.enqueueInt(rTTL(key))
}

func (p *Pipeline) Type(key string) *RespString {
	return p.enqueueStr(rType(key))
}

func (p *Pipeline) Wait(numSlaves int, timeout int) *RespInt {
	return p.enqueueInt(rWait(numSlaves, timeout))
}

//////////
//STRINGS
func (p *Pipeline) Append(key string, value string) *RespInt {
	return p.enqueueInt(rAppend(key, value))
}

func (p *Pipeline) BitCount(key string) *RespInt {
	return p.enqueueInt(rBitCount(key))
}

func (p *Pipeline) BitOp(bitOperation BitOperation, destKey string, keys ...string) *RespInt {
	return p.enqueueIntErr(rBitOp(bitOperation, destKey, keys...))
}

//TODO: agregar verificacion start end
func (p *Pipeline) BitPos(key string, bit bool, startEnd ...int) *RespInt {
	return p.enqueueInt(rBitPos(key, bit, startEnd...))
}

func (p *Pipeline) Get(key string) *RespString {
	return p.enqueueStr(rGet(key))
}

func (p *Pipeline) GetBit(key string, offset int) *RespInt {
	return p.enqueueInt(rGetBit(key, offset))
}

func (p *Pipeline) GetRange(key string, start int, end int) *RespString {
	return p.enqueueStr(rGetRange(key, start, end))
}

func (p *Pipeline) GetSet(key string, value string) *RespString {
	return p.enqueueStr(rGetSet(key, value))
}

func (p *Pipeline) Set(key string, value string) *RespString {
	return p.enqueueStr(rSet(key, value))
}

//TODO agregar verificacion
func (p *Pipeline) SetX(key string, value string, keyExpiration *KeyExpiration, keyExistance *KeyExistance) *RespString {
	return p.enqueueStrErr(rSetX(key, value, keyExpiration, keyExistance))
}

func (p *Pipeline) Incr(key string) *RespInt {
	return p.enqueueInt(rIncr(key))
}

func (p *Pipeline) IncrBy(key string, increment int) *RespInt {
	return p.enqueueInt(rIncrBy(key, increment))
}

func (p *Pipeline) IncrByFloat(key string, increment float64) *RespFloat {
	return p.enqueueFloat(rIncrByFloat(key, increment))
}

func (p *Pipeline) Decr(key string) *RespInt {
	return p.enqueueInt(rDecr(key))
}

func (p *Pipeline) DecrBy(key string, decrement int) *RespInt {
	return p.enqueueInt(rDecrBy(key, decrement))
}

func (p *Pipeline) MGet(keys ...string) *RespStringArray {
	return p.enqueueStrArrayErr(rMGet(keys...))
}

func (p *Pipeline) MSet(keyValues ...string) *RespString {
	return p.enqueueStrErr(rMSet(keyValues...))
}

func (p *Pipeline) MSetNx(keyValues ...string) *RespBool {
	return p.enqueueBoolErr(rMSetNx(keyValues...))
}

func (p *Pipeline) PSetEx(key string, milliseconds int, value string) *RespString {
	return p.enqueueStr(rPSetEx(key, milliseconds, value))
}

func (p *Pipeline) SetEx(key string, seconds int, value string) *RespString {
	return p.enqueueStr(rSetEx(key, seconds, value))
}

func (p *Pipeline) SetBit(key string, offset int, value bool) *RespInt {
	return p.enqueueInt(rSetBit(key, offset, value))
}

func (p *Pipeline) SetNx(key string, value string) *RespBool {
	return p.enqueueBool(rSetNx(key, value))
}

func (p *Pipeline) SetRange(key string, offset int, value string) *RespInt {
	return p.enqueueInt(rSetRange(key, offset, value))
}

func (p *Pipeline) StrLen(key string) *RespInt {
	return p.enqueueInt(rStrLen(key))
}

////////
//LISTS
///////

func (p *Pipeline) LLen(key string) *RespInt {
	return p.enqueueInt(rLLen(key))
}

func (p *Pipeline) LIndex(key string, index int) *RespString {
	return p.enqueueStr(rLIndex(key, index))
}

func (p *Pipeline) LPush(key string, values ...string) *RespInt {
	return p.enqueueIntErr(rLPush(key, values...))
}

func (p *Pipeline) LPushX(key string, value string) *RespInt {
	return p.enqueueInt(rLPushX(key, value))
}

func (p *Pipeline) RPush(key string, values ...string) *RespInt {
	return p.enqueueIntErr(rRPush(key, values...))
}

func (p *Pipeline) RPushX(key string, value string) *RespInt {
	return p.enqueueInt(rRPushX(key, value))
}

func (p *Pipeline) LPop(key string) *RespString {
	return p.enqueueStr(rLPop(key))
}

func (p *Pipeline) RPop(key string) *RespString {
	return p.enqueueStr(rRPop(key))
}

func (p *Pipeline) LSet(key string, index int, value string) *RespString {
	return p.enqueueStr(rLSet(key, index, value))
}

func (p *Pipeline) LInsert(key string, location InsertLocation, pivot string, value string) *RespInt {
	return p.enqueueIntErr(rLInsert(key, location, pivot, value))
}

func (p *Pipeline) LRange(key string, start int, stop int) *RespStringArray {
	return p.enqueueStrArray(rLRange(key, start, stop))
}

func (p *Pipeline) LRem(key string, count int, value string) *RespInt {
	return p.enqueueInt(rLRem(key, count, value))
}

func (p *Pipeline) LTrim(key string, start int, stop int) *RespString {
	return p.enqueueStr(rLTrim(key, start, stop))
}

func (p *Pipeline) RPopLPush(source string, destination string) *RespString {
	return p.enqueueStr(rRPopLPush(source, destination))
}

func (p *Pipeline) BLPop(timeout int, keys ...string) *RespStringArray {
	return p.enqueueStrArrayErr(rBLPop(timeout, keys...))
}

func (p *Pipeline) BRPop(timeout int, keys ...string) *RespStringArray {
	return p.enqueueStrArrayErr(rBRPop(timeout, keys...))
}

func (p *Pipeline) BRPopLPush(source string, destination string, timeout int) *RespString {
	return p.enqueueStr(rBRPopLPush(source, destination, timeout))
}

///////
//SETS
//////

func (p *Pipeline) SAdd(key string, values ...string) *RespInt {
	return p.enqueueIntErr(rSAdd(key, values...))
}

func (p *Pipeline) SCard(key string) *RespInt {
	return p.enqueueInt(rSCard(key))
}

func (p *Pipeline) SDiff(keys ...string) *RespStringArray {
	return p.enqueueStrArrayErr(rSDiff(keys...))
}

func (p *Pipeline) SDiffStore(key string, keys ...string) *RespInt {
	return p.enqueueIntErr(rSDiffStore(key, keys...))
}

func (p *Pipeline) SInter(keys ...string) *RespStringArray {
	return p.enqueueStrArrayErr(rSInter(keys...))
}

func (p *Pipeline) SInterStore(key string, keys ...string) *RespInt {
	return p.enqueueIntErr(rSInterStore(key, keys...))
}

func (p *Pipeline) SIsMember(key string, value string) *RespBool {
	return p.enqueueBool(rSIsMember(key, value))
}

func (p *Pipeline) SMembers(key string) *RespStringArray {
	return p.enqueueStrArray(rSMembers(key))
}

func (p *Pipeline) SMove(sourceKey string, destinationKey string, value string) *RespBool {
	return p.enqueueBool(rSMove(sourceKey, destinationKey, value))
}

func (p *Pipeline) SPop(key string) *RespString {
	return p.enqueueStr(rSPop(key))
}

func (p *Pipeline) SRandMember(key string, count int) *RespStringArray {
	return p.enqueueStrArray(rSRandMember(key, count))
}

func (p *Pipeline) SRem(key string, values ...string) *RespInt {
	return p.enqueueIntErr(rSRem(key, values...))
}

func (p *Pipeline) SUnion(keys ...string) *RespStringArray {
	return p.enqueueStrArrayErr(rSUnion(keys...))
}

func (p *Pipeline) SUnionStore(key string, keys ...string) *RespInt {
	return p.enqueueIntErr(rSUnionStore(key, keys...))
}

////////
//HASHES
///////

func (p *Pipeline) HDel(key string, fields ...string) *RespInt {
	return p.enqueueIntErr(rHDel(key, fields...))
}

func (p *Pipeline) HExists(key string, field string) *RespBool {
	return p.enqueueBool(rHExists(key, field))
}

func (p *Pipeline) HGet(key string, field string) *RespString {
	return p.enqueueStr(rHGet(key, field))
}

func (p *Pipeline) HGetAll(key string) *RespStringArray {
	return p.enqueueStrArray(rHGetAll(key))
}

func (p *Pipeline) HIncrBy(key string, field string, increment int) *RespInt {
	return p.enqueueInt(rHIncrBy(key, field, increment))
}

func (p *Pipeline) HIncrByFloat(key string, field string, increment float64) *RespFloat {
	return p.enqueueFloat(rHIncrByFloat(key, field, increment))
}

func (p *Pipeline) HKeys(key string) *RespStringArray {
	return p.enqueueStrArray(rHKeys(key))
}

func (p *Pipeline) HLen(key string) *RespInt {
	return p.enqueueInt(rHLen(key))
}

func (p *Pipeline) HMGet(key string, fields ...string) *RespStringArray {
	return p.enqueueStrArrayErr(rHMGet(key, fields...))
}

func (p *Pipeline) HMSet(key string, fieldValues ...string) *RespString {
	return p.enqueueStrErr(rHMSet(key, fieldValues...))
}

func (p *Pipeline) HSet(key string, field string, value string) *RespBool {
	return p.enqueueBool(rHSet(key, field, value))
}

func (p *Pipeline) HSetNx(key string, field string, value string) *RespBool {
	return p.enqueueBool(rHSetNx(key, field, value))
}

/* Since 3.2.0
func (p *Pipeline) HStrLen(key string, field string) *RespInt {
	return p.enqueueInt(rHStrLen(key, field))
}
*/

func (p *Pipeline) HVals(key string) *RespStringArray {
	return p.enqueueStrArray(rHVals(key))
}

//HScan

/////////////
//HYPERLOGLOG
/////////////

func (p *Pipeline) PFAdd(key string, elements ...string) *RespInt {
	return p.enqueueIntErr(rPFAdd(key, elements...))
}

func (p *Pipeline) PFCount(keys ...string) *RespInt {
	return p.enqueueIntErr(rPFCount(keys...))
}

func (p *Pipeline) PFMerge(destkey string, sourcekeys ...string) *RespString {
	return p.enqueueStrErr(rPFMerge(destkey, sourcekeys...))
}

////////
///////

/////////////
//CONNECTION
////////////

func (p *Pipeline) Select(index uint) *RespString {
	return p.enqueueStr(rSelect(index))
}

///////
///////

func (p *Pipeline) enqueueResp(cmds [][]byte, rs pipelineResponse, err error) {

	p.cmdsQueue.enqueue(cmds)
	p.respQueue.enqueue(rs)

	for _, c := range cmds {
		p.cmdsSize += len(c)
	}

	if err != nil {
		rs.setErr(err)
		p.err = PipelineInputErr
	}

}

func (p *Pipeline) enqueueStr(cmds [][]byte) *RespString {
	rs := new(RespString)
	p.enqueueResp(cmds, rs, nil)

	return rs
}

func (p *Pipeline) enqueueStrArray(cmds [][]byte) *RespStringArray {
	rs := new(RespStringArray)
	p.enqueueResp(cmds, rs, nil)

	return rs
}

func (p *Pipeline) enqueueInt(cmds [][]byte) *RespInt {
	rs := new(RespInt)
	p.enqueueResp(cmds, rs, nil)

	return rs
}

func (p *Pipeline) enqueueFloat(cmds [][]byte) *RespFloat {
	rs := new(RespFloat)
	p.enqueueResp(cmds, rs, nil)

	return rs
}

func (p *Pipeline) enqueueBool(cmds [][]byte) *RespBool {
	rs := new(RespBool)
	p.enqueueResp(cmds, rs, nil)

	return rs
}

func (p *Pipeline) enqueueStrErr(cmds [][]byte, err error) *RespString {
	rs := new(RespString)
	p.enqueueResp(cmds, rs, err)

	return rs
}

func (p *Pipeline) enqueueStrArrayErr(cmds [][]byte, err error) *RespStringArray {
	rs := new(RespStringArray)
	p.enqueueResp(cmds, rs, err)

	return rs
}

func (p *Pipeline) enqueueIntErr(cmds [][]byte, err error) *RespInt {
	rs := new(RespInt)
	p.enqueueResp(cmds, rs, err)

	return rs
}

func (p *Pipeline) enqueueFloatErr(cmds [][]byte, err error) *RespFloat {
	rs := new(RespFloat)
	p.enqueueResp(cmds, rs, err)

	return rs
}

func (p *Pipeline) enqueueBoolErr(cmds [][]byte, err error) *RespBool {
	rs := new(RespBool)
	p.enqueueResp(cmds, rs, err)

	return rs
}

func (p *Pipeline) execute(multi bool) (err error) {

	///////////////
	//Input errors = return!
	if p.err != nil {
		return p.err
	}

	//Get a connection
	conn := p.redis.pool.get()
	defer p.cleanConn(conn)

	//MULTI = call "EXEC" command
	if multi {
		p.enqueueStrArray(rExec())
	}

	////////////////
	//WRITE COMMANDS
	if err = p.writeExecute(conn); err != nil {
		return
	}

	if multi {
		err = p.readExecuteMulti(conn)
	} else {
		err = p.readExecute(conn)
	}

	return
}

func (p *Pipeline) writeExecute(conn *net.TCPConn) (err error) {

	cmdDeq := func() [][]byte {
		if c, ok := p.cmdsQueue.dequeue().([][]byte); ok {
			return c
		}
		return nil
	}

	writer := bufio.NewWriterSize(conn, p.cmdsSize)
	for cmd := cmdDeq(); cmd != nil; cmd = cmdDeq() {
		err = write(cmd, writer)
		if err != nil {
			return
		}
	}

	err = writer.Flush()

	return

}

func (p *Pipeline) respDeq() pipelineResponse {
	if c, ok := p.respQueue.dequeue().(pipelineResponse); ok {
		return c
	}
	return nil
}

func (p *Pipeline) readExecute(conn *net.TCPConn) (err error) {

	reader := bufio.NewReader(conn)

	for pipeResp := p.respDeq(); pipeResp != nil; pipeResp = p.respDeq() {

		rr, err := read(reader)
		if err == nil {
			pipeResp.read(rr)
		}

	}

	return
}

func (p *Pipeline) readExecuteMulti(conn *net.TCPConn) (err error) {

	var multiResp *redisResponse
	reader := bufio.NewReader(conn)

	if q0, err := readString(read(reader)); q0 != "OK" || err != nil {
		return err
	}

	for i := 1; i < p.respQueue.size-1; i++ {
		if q, err := readString(read(reader)); q != "QUEUED" || err != nil {
			return err
		}
	}

	//Read response from "exec" command from multi
	if multiResp, err = read(reader); err != nil {
		return
	}

	//Dequeue de first element which is "multi"
	p.respDeq()

	//Dequeue all elements but not the last one (which is "exec")
	for i := 0; i < p.respQueue.size-2; i++ {
		p.respDeq().read(multiResp.elements[i])
	}

	return
}

func (p *Pipeline) cleanConn(conn *net.TCPConn) {

	go func() {
		m := multiCompile("SELECT", "0")
		writeRead(m, conn)

		m = multiCompile("DISCARD")
		writeRead(m, conn)

		p.redis.pool.put(conn)
	}()

}
