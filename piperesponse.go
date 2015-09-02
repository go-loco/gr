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
	result, _ := readString(rr, nil)
	rs.Value = result
}

//String Array
type RespStringArray struct {
	Value []string
	Error error
}

func (rs *RespStringArray) set(b *redisResponse) {
	result, _ := readStringArray(b, nil)
	rs.Value = result
}

//Int
type RespInt struct {
	Value int64
	Error error
}

func (rs *RespInt) set(b *redisResponse) {
	result, _ := readInt64(b, nil)
	rs.Value = result
}

//Float
type RespFloat struct {
	Value float64
	Error error
}

func (rs *RespFloat) set(b *redisResponse) {
	result, _ := readFloat64(b, nil)
	rs.Value = result
}

//Bool
type RespBool struct {
	Value bool
	Error error
}

func (rs *RespBool) set(b *redisResponse) {
	result, _ := readBool(b, nil)
	rs.Value = result
}
