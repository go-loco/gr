package gr

type pipelineResponse interface {
	read(*redisResponse)
	setErr(error)
}

//String
type RespString struct {
	Value string
	Error error
}

func (rs *RespString) read(rr *redisResponse) {
	rs.Value, rs.Error = readString(rr, nil)
}

func (rs *RespString) setErr(err error) {
	rs.Error = err
}

//String Array
type RespStringArray struct {
	Value []string
	Error error
}

func (rs *RespStringArray) read(b *redisResponse) {
	rs.Value, rs.Error = readStringArray(b, nil)
}

func (rs *RespStringArray) setErr(err error) {
	rs.Error = err
}

//Int
type RespInt struct {
	Value int64
	Error error
}

func (rs *RespInt) read(b *redisResponse) {
	rs.Value, rs.Error = readInt64(b, nil)
}

func (rs *RespInt) setErr(err error) {
	rs.Error = err
}

//Float
type RespFloat struct {
	Value float64
	Error error
}

func (rs *RespFloat) read(b *redisResponse) {
	rs.Value, rs.Error = readFloat64(b, nil)
}

func (rs *RespFloat) setErr(err error) {
	rs.Error = err
}

//Bool
type RespBool struct {
	Value bool
	Error error
}

func (rs *RespBool) read(b *redisResponse) {
	rs.Value, rs.Error = readBool(b, nil)
}

func (rs *RespBool) setErr(err error) {
	rs.Error = err
}
