package gr

type Response interface {
	set(*redisResponse)
}

//String
type RespString struct {
	Value string
	Error error
}

func (rs *RespString) set(rr *redisResponse) {
	rs.Value, rs.Error = readString(rr, nil)
}

//String Array
type RespStringArray struct {
	Value []string
	Error error
}

func (rs *RespStringArray) set(b *redisResponse) {
	rs.Value, rs.Error = readStringArray(b, nil)
}

//Int
type RespInt struct {
	Value int64
	Error error
}

func (rs *RespInt) set(b *redisResponse) {
	rs.Value, rs.Error = readInt64(b, nil)
}

//Float
type RespFloat struct {
	Value float64
	Error error
}

func (rs *RespFloat) set(b *redisResponse) {
	rs.Value, rs.Error = readFloat64(b, nil)
}

//Bool
type RespBool struct {
	Value bool
	Error error
}

func (rs *RespBool) set(b *redisResponse) {
	rs.Value, rs.Error = readBool(b, nil)
}
