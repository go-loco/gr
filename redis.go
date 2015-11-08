package gr

import (
	"errors"
	"strconv"
	"time"
)

type Redis struct {
	config Config
	pool   *pool
}

type Config struct {
	Address        string
	Port           int
	Database       string
	Password       string
	MinConnections int
	Timeout        int
}

func New() (*Redis, error) {
	return NewWithConfig(
		Config{
			Address:        "localhost",
			Port:           6379,
			MinConnections: 2,
		},
	)
}

func NewWithConfig(config Config) (*Redis, error) {

	var r *Redis
	pool, err := newPool(&config)

	if err == nil {
		r = &Redis{
			config: config,
			pool:   pool,
		}
	}

	return r, err
}

var (
	NilErr             = errors.New("Received a nil response.")
	NotConnectedErr    = errors.New("Client is not connected.")
	NotInitializedErr  = errors.New("Client is not initialized. Use redis.New()")
	PoolErr            = errors.New("Pool error")
	ParamErr           = errors.New("Error in parameters")
	NotEnoughParamsErr = errors.New("Not enough parameters")
	ParamsNotTuplesErr = errors.New("Parameters must be tuples (x,y)")
)

func (r *Redis) Pipelined(caller func(*Pipeline)) []error {

	p := &Pipeline{
		cmdsQueue: queue{},
		respQueue: queue{},
		redis:     r,
		err:       make([]error, 0, 3),
	}

	caller(p)

	if len(p.err) > 0 {
		return p.err
	}

	if err := p.execute(); err != nil {
		return []error{err}
	}

	return nil

}

////////
//PUBSUB
///////

func (r *Redis) Publish(channel string, message string) (int64, error) {
	return r.writeReadInt(rPublish(channel, message))
}

func (r *Redis) Subscribe(caller func(*PubSub), channels ...string) {
	r.pubSubStart(false, caller, channels...)
}

func (r *Redis) PSubscribe(caller func(*PubSub), channels ...string) {
	r.pubSubStart(true, caller, channels...)
}

func (r *Redis) pubSubStart(withPattern bool, caller func(*PubSub), channels ...string) {

	ps, _ := newPubSub(r)

	if withPattern {
		go ps.pSubscribe(channels...)
	} else {
		go ps.subscribe(channels...)
	}

	caller(ps)
	ps.unSubscribe()

}

func scanGeneric(s []string, err error) (int64, []string, error) {
	c, err := strconv.ParseInt(string(s[0]), 10, 64)
	return c, s[1:], err
}

////////
//KEYS
///////

func (r *Redis) Del(keys ...string) (int64, error) {
	rs, err := rDel(keys...)
	if err != nil {
		return 0, err
	}

	return r.writeReadInt(rs)
}

func (r *Redis) Dump(key string) (string, error) {
	return r.writeReadStr(rDump(key))
}

func (r *Redis) Exists(key string) (bool, error) {
	return r.writeReadBool(rExists(key))
}

func (r *Redis) Expire(key string, seconds int) (bool, error) {
	return r.writeReadBool(rExpire(key, seconds))
}

func (r *Redis) ExpireAt(key string, date time.Time) (bool, error) {
	return r.writeReadBool(rExpireAt(key, date))
}

func (r *Redis) Keys(pattern string) ([]string, error) {
	return r.writeReadStrArray(rKeys(pattern))
}

func (r *Redis) Migrate(
	host string,
	port int,
	key string,
	destinationDB string,
	timeout int,
	copy bool,
	replace bool) (string, error) {

	return r.writeReadStr(rMigrate(host, port, key, destinationDB, timeout, copy, replace))
}

func (r *Redis) Move(key string, db string) (bool, error) {
	return r.writeReadBool(rMove(key, db))
}

func (r *Redis) ObjectEncoding(arguments ...string) (string, error) {
	return r.writeReadStr(rObjectEncoding(arguments...))
}

func (r *Redis) ObjectRefCount(arguments ...string) (int64, error) {
	return r.writeReadInt(rObjectRefCount(arguments...))
}

func (r *Redis) ObjectIdleTime(arguments ...string) (int64, error) {
	return r.writeReadInt(rObjectIdleTime(arguments...))
}

func (r *Redis) Persist(key string) (bool, error) {
	return r.writeReadBool(rPersist(key))
}

func (r *Redis) PExpire(key string, milliseconds int) (bool, error) {
	return r.writeReadBool(rPExpire(key, milliseconds))
}

func (r *Redis) PExpireAt(key string, date time.Time) (bool, error) {
	return r.writeReadBool(rPExpireAt(key, date))
}

func (r *Redis) PTTL(key string) (int64, error) {
	return r.writeReadInt(rPTTL(key))
}

func (r *Redis) RandomKey() (string, error) {
	return r.writeReadStr(rRandomKey())
}

func (r *Redis) Rename(key string, newKey string) (string, error) {
	return r.writeReadStr(rRename(key, newKey))
}

func (r *Redis) RenameNx(key string, newKey string) (bool, error) {
	return r.writeReadBool(rRenameNx(key, newKey))
}

func (r *Redis) Restore(key string, ttl int, value string, replace bool) (string, error) {
	return r.writeReadStr(rRestore(key, ttl, value, replace))
}

func (r *Redis) Scan(cursor int, scanParams *ScanParams) (int64, []string, error) {
	s, err := r.writeReadStrArray(rScan(cursor, scanParams))
	c, sc, err := scanGeneric(s, err)

	return c, sc, err
}

func (r *Redis) Sort(key string, sortParams *SortParams) ([]string, error) {
	return r.writeReadStrArray(rSort(key, sortParams))
}

func (r *Redis) SortStore(key string, destination string, sortParams *SortParams) (int64, error) {
	return r.writeReadInt(rSortStore(key, destination, sortParams))
}

func (r *Redis) TTL(key string) (int64, error) {
	return r.writeReadInt(rTTL(key))
}

func (r *Redis) Type(key string) (string, error) {
	return r.writeReadStr(rType(key))
}

func (r *Redis) Wait(numSlaves int, timeout int) (int64, error) {
	return r.writeReadInt(rWait(numSlaves, timeout))
}

//////////
//STRINGS
func (r *Redis) Append(key string, value string) (int64, error) {
	return r.writeReadInt(rAppend(key, value))
}

func (r *Redis) BitCount(key string) (int64, error) {
	return r.writeReadInt(rBitCount(key))
}

func (r *Redis) BitOp(bitOperation BitOperation, destKey string, keys ...string) (int64, error) {

	rs, err := rBitOp(bitOperation, destKey, keys...)
	if err != nil {
		return 0, err
	}

	return r.writeReadInt(rs)
}

func (r *Redis) BitPos(key string, bit bool, startEnd ...int) (int64, error) {
	return r.writeReadInt(rBitPos(key, bit, startEnd...))
}

func (r *Redis) Get(key string) (string, error) {
	return r.writeReadStr(rGet(key))
}

func (r *Redis) GetBit(key string, offset int) (int64, error) {
	return r.writeReadInt(rGetBit(key, offset))
}

func (r *Redis) GetRange(key string, start int, end int) (string, error) {
	return r.writeReadStr(rGetRange(key, start, end))
}

func (r *Redis) GetSet(key string, value string) (string, error) {
	return r.writeReadStr(rGetSet(key, value))
}

func (r *Redis) Set(key string, value string) (string, error) {
	return r.writeReadStr(rSet(key, value))
}

func (r *Redis) SetX(key string, value string, keyExpiration *KeyExpiration, keyExistance *KeyExistance) (string, error) {

	rs, err := rSetX(key, value, keyExpiration, keyExistance)
	if err != nil {
		return "", err
	}

	return r.writeReadStr(rs)
}

func (r *Redis) Incr(key string) (int64, error) {
	return r.writeReadInt(rIncr(key))
}

func (r *Redis) IncrBy(key string, increment int) (int64, error) {
	return r.writeReadInt(rIncrBy(key, increment))
}

func (r *Redis) IncrByFloat(key string, increment float64) (float64, error) {
	return r.writeReadFloat(rIncrByFloat(key, increment))
}

func (r *Redis) Decr(key string) (int64, error) {
	return r.writeReadInt(rDecr(key))
}

func (r *Redis) DecrBy(key string, decrement int) (int64, error) {
	return r.writeReadInt(rDecrBy(key, decrement))
}

func (r *Redis) MGet(keys ...string) ([]string, error) {

	rs, err := rMGet(keys...)
	if err != nil {
		return nil, err
	}

	return r.writeReadStrArray(rs)
}

func (r *Redis) MSet(keyValues ...string) (string, error) {

	rs, err := rMSet(keyValues...)
	if err != nil {
		return "", err
	}

	return r.writeReadStr(rs)
}

func (r *Redis) MSetNx(keyValues ...string) (bool, error) {

	rs, err := rMSetNx(keyValues...)
	if err != nil {
		return false, err
	}

	return r.writeReadBool(rs)
}

func (r *Redis) PSetEx(key string, milliseconds int, value string) (string, error) {
	return r.writeReadStr(rPSetEx(key, milliseconds, value))
}

func (r *Redis) SetEx(key string, seconds int, value string) (string, error) {
	return r.writeReadStr(rSetEx(key, seconds, value))
}

func (r *Redis) SetBit(key string, offset int, value bool) (int64, error) {
	return r.writeReadInt(rSetBit(key, offset, value))
}

func (r *Redis) SetNx(key string, value string) (bool, error) {
	return r.writeReadBool(rSetNx(key, value))
}

func (r *Redis) SetRange(key string, offset int, value string) (int64, error) {
	return r.writeReadInt(rSetRange(key, offset, value))
}

func (r *Redis) StrLen(key string) (int64, error) {
	return r.writeReadInt(rStrLen(key))
}

////////
//LISTS
///////

func (r *Redis) LLen(key string) (int64, error) {
	return r.writeReadInt(rLLen(key))
}

func (r *Redis) LIndex(key string, index int) (string, error) {
	return r.writeReadStr(rLIndex(key, index))
}

func (r *Redis) LPush(key string, values ...string) (int64, error) {
	rs, err := rLPush(key, values...)
	if err != nil {
		return 0, err
	}

	return r.writeReadInt(rs)
}

func (r *Redis) LPushX(key string, value string) (int64, error) {
	return r.writeReadInt(rLPushX(key, value))
}

func (r *Redis) RPush(key string, values ...string) (int64, error) {
	rs, err := rRPush(key, values...)
	if err != nil {
		return 0, err
	}

	return r.writeReadInt(rs)
}

func (r *Redis) RPushX(key string, value string) (int64, error) {
	return r.writeReadInt(rRPushX(key, value))
}

func (r *Redis) LPop(key string) (string, error) {
	return r.writeReadStr(rLPop(key))
}

func (r *Redis) RPop(key string) (string, error) {
	return r.writeReadStr(rRPop(key))
}

func (r *Redis) LSet(key string, index int, value string) (string, error) {
	return r.writeReadStr(rLSet(key, index, value))
}

func (r *Redis) LInsert(key string, location InsertLocation, pivot string, value string) (int64, error) {

	rs, err := rLInsert(key, location, pivot, value)
	if err != nil {
		return 0, err
	}

	return r.writeReadInt(rs)
}

func (r *Redis) LRange(key string, start int, stop int) ([]string, error) {
	return r.writeReadStrArray(rLRange(key, start, stop))
}

func (r *Redis) LRem(key string, count int, value string) (int64, error) {
	return r.writeReadInt(rLRem(key, count, value))
}

func (r *Redis) LTrim(key string, start int, stop int) (string, error) {
	return r.writeReadStr(rLTrim(key, start, stop))
}

func (r *Redis) RPopLPush(source string, destination string) (string, error) {
	return r.writeReadStr(rRPopLPush(source, destination))
}

func (r *Redis) BLPop(timeout int, keys ...string) ([]string, error) {
	rs, err := rBLPop(timeout, keys...)
	if err != nil {
		return nil, err
	}

	return r.writeReadStrArray(rs)
}

func (r *Redis) BRPop(timeout int, keys ...string) ([]string, error) {
	rs, err := rBRPop(timeout, keys...)
	if err != nil {
		return nil, err
	}

	return r.writeReadStrArray(rs)
}

func (r *Redis) BRPopLPush(source string, destination string, timeout int) (string, error) {
	return r.writeReadStr(rBRPopLPush(source, destination, timeout))
}

///////
//SETS
//////

func (r *Redis) SAdd(key string, values ...string) (int64, error) {
	rs, err := rSAdd(key, values...)
	if err != nil {
		return 0, err
	}

	return r.writeReadInt(rs)
}

func (r *Redis) SCard(key string) (int64, error) {
	return r.writeReadInt(rSCard(key))
}

func (r *Redis) SDiff(keys ...string) ([]string, error) {
	rs, err := rSDiff(keys...)
	if err != nil {
		return nil, err
	}

	return r.writeReadStrArray(rs)
}

func (r *Redis) SDiffStore(key string, keys ...string) (int64, error) {
	rs, err := rSDiffStore(key, keys...)
	if err != nil {
		return 0, err
	}

	return r.writeReadInt(rs)
}

func (r *Redis) SInter(keys ...string) ([]string, error) {
	rs, err := rSInter(keys...)
	if err != nil {
		return nil, err
	}

	return r.writeReadStrArray(rs)
}

func (r *Redis) SInterStore(key string, keys ...string) (int64, error) {
	rs, err := rSInterStore(key, keys...)
	if err != nil {
		return 0, err
	}

	return r.writeReadInt(rs)
}

func (r *Redis) SIsMember(key string, value string) (bool, error) {
	return r.writeReadBool(rSIsMember(key, value))
}

func (r *Redis) SMembers(key string) ([]string, error) {
	return r.writeReadStrArray(rSMembers(key))
}

func (r *Redis) SMove(sourceKey string, destinationKey string, value string) (bool, error) {
	return r.writeReadBool(rSMove(sourceKey, destinationKey, value))
}

func (r *Redis) SPop(key string) (string, error) {
	return r.writeReadStr(rSPop(key))
}

func (r *Redis) SRandMember(key string, count int) ([]string, error) {
	return r.writeReadStrArray(rSRandMember(key, count))
}

func (r *Redis) SRem(key string, values ...string) (int64, error) {
	rs, err := rSRem(key, values...)
	if err != nil {
		return 0, err
	}

	return r.writeReadInt(rs)
}

func (r *Redis) SScan(key string, cursor int, scanParams *ScanParams) (int64, []string, error) {
	s, err := r.writeReadStrArray(rSScan(key, cursor, scanParams))
	return scanGeneric(s, err)
}

func (r *Redis) SUnion(keys ...string) ([]string, error) {
	rs, err := rSUnion(keys...)
	if err != nil {
		return nil, err
	}

	return r.writeReadStrArray(rs)
}

func (r *Redis) SUnionStore(key string, keys ...string) (int64, error) {
	rs, err := rSUnionStore(key, keys...)
	if err != nil {
		return 0, err
	}

	return r.writeReadInt(rs)
}

////////
//HASHES
///////

func (r *Redis) HDel(key string, fields ...string) (int64, error) {
	rs, err := rHDel(key, fields...)
	if err != nil {
		return 0, err
	}

	return r.writeReadInt(rs)
}

func (r *Redis) HExists(key string, field string) (bool, error) {
	return r.writeReadBool(rHExists(key, field))
}

func (r *Redis) HGet(key string, field string) (string, error) {
	return r.writeReadStr(rHGet(key, field))
}

func (r *Redis) HGetAll(key string) ([]string, error) {
	return r.writeReadStrArray(rHGetAll(key))
}

func (r *Redis) HIncrBy(key string, field string, increment int) (int64, error) {
	return r.writeReadInt(rHIncrBy(key, field, increment))
}

func (r *Redis) HIncrByFloat(key string, field string, increment float64) (float64, error) {
	return r.writeReadFloat(rHIncrByFloat(key, field, increment))
}

func (r *Redis) HKeys(key string) ([]string, error) {
	return r.writeReadStrArray(rHKeys(key))
}

func (r *Redis) HLen(key string) (int64, error) {
	return r.writeReadInt(rHLen(key))
}

func (r *Redis) HMGet(key string, fields ...string) ([]string, error) {
	rs, err := rHMGet(key, fields...)
	if err != nil {
		return nil, err
	}

	return r.writeReadStrArray(rs)
}

func (r *Redis) HMSet(key string, fieldValues ...string) (string, error) {
	rs, err := rHMSet(key, fieldValues...)
	if err != nil {
		return "", err
	}

	return r.writeReadStr(rs)
}

func (r *Redis) HSet(key string, field string, value string) (bool, error) {
	return r.writeReadBool(rHSet(key, field, value))
}

func (r *Redis) HSetNx(key string, field string, value string) (bool, error) {
	return r.writeReadBool(rHSetNx(key, field, value))
}

/* .3.0.2
func (r *Redis) HStrLen(key string, field string) (int64, error) {
	return r.writeReadInt(rHStrLen(key, field))
}
*/

func (r *Redis) HVals(key string) ([]string, error) {
	return r.writeReadStrArray(rHVals(key))
}

func (r *Redis) HScan(key string, cursor int, scanParams *ScanParams) (int64, []string, error) {
	s, err := r.writeReadStrArray(rHScan(key, cursor, scanParams))
	return scanGeneric(s, err)
}

////////
//HYPERLOGLOG
///////

func (r *Redis) PFAdd(key string, elements ...string) (int64, error) {
	rs, err := rPFAdd(key, elements...)
	if err != nil {
		return 0, err
	}

	return r.writeReadInt(rs)
}

func (r *Redis) PFCount(keys ...string) (int64, error) {
	rs, err := rPFCount(keys...)
	if err != nil {
		return 0, err
	}

	return r.writeReadInt(rs)
}

func (r *Redis) PFMerge(destkey string, sourcekeys ...string) (string, error) {
	rs, err := rPFMerge(destkey, sourcekeys...)
	if err != nil {
		return "", err
	}

	return r.writeReadStr(rs)
}

/////
////

func (r *Redis) writeReadGeneric(cmds [][]byte) (result *redisResponse, err error) {

	conn := r.pool.get()
	defer r.pool.put(conn)

	result, err = writeRead(cmds, conn)
	return
}

func (r *Redis) writeReadStr(cmds [][]byte) (result string, err error) {
	result, err = readString(r.writeReadGeneric(cmds))
	return
}

func (r *Redis) writeReadStrArray(cmds [][]byte) (result []string, err error) {
	result, err = readStringArray(r.writeReadGeneric(cmds))
	return
}

func (r *Redis) writeReadInt(cmds [][]byte) (result int64, err error) {
	result, err = readInt64(r.writeReadGeneric(cmds))
	return
}

func (r *Redis) writeReadFloat(cmds [][]byte) (result float64, err error) {
	result, err = readFloat64(r.writeReadGeneric(cmds))
	return
}

func (r *Redis) writeReadBool(cmds [][]byte) (result bool, err error) {
	result, err = readBool(r.writeReadGeneric(cmds))
	return
}
