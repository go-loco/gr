package gr

import "strconv"

var _crlf = []byte("\r\n")

const (
	tString     = '+'
	tBulkString = '$'
	tInteger    = ':'
	tArray      = '*'
	tError      = '-'
)

func compileBulkString(cmd string) []byte {

	cmdLength := len(cmd)
	strCmdLength := strconv.Itoa(cmdLength)
	arrayLength := 1 + len(strCmdLength) + 4 + cmdLength

	c := make([]byte, 0, arrayLength)

	c = append(c, tBulkString)
	c = append(c, strCmdLength...)
	c = append(c, _crlf...)
	c = append(c, cmd...)
	c = append(c, _crlf...)

	return c
}

func multiCompile(values ...string) [][]byte {
	r := make([][]byte, len(values))

	for i, v := range values {
		r[i] = compileBulkString(v)
	}

	return r
}

func multiCompile1(v1 string, values ...string) [][]byte {
	r := make([][]byte, len(values)+1)
	r[0] = compileBulkString(v1)

	for i, v := range values {
		r[i+1] = compileBulkString(v)
	}

	return r
}

func multiCompile2(v1 string, v2 string, values ...string) [][]byte {
	r := make([][]byte, len(values)+2)
	r[0] = compileBulkString(v1)
	r[1] = compileBulkString(v2)

	for i, v := range values {
		r[i+2] = compileBulkString(v)
	}

	return r
}

func multiCompile3(v1 string, v2 string, v3 string, values ...string) [][]byte {
	r := make([][]byte, len(values)+3)
	r[0] = compileBulkString(v1)
	r[1] = compileBulkString(v2)
	r[2] = compileBulkString(v3)

	for i, v := range values {
		r[i+3] = compileBulkString(v)
	}

	return r
}

func multiCompile4(v1 string, v2 string, v3 string, v4 string, values ...string) [][]byte {
	r := make([][]byte, len(values)+4)
	r[0] = compileBulkString(v1)
	r[1] = compileBulkString(v2)
	r[2] = compileBulkString(v3)
	r[3] = compileBulkString(v4)

	for i, v := range values {
		r[i+4] = compileBulkString(v)
	}

	return r
}

func debugCmds(cmds [][]byte) {

	println("")
	println("********")
	println("DEBUG")
	println("********")

	for _, c := range cmds {
		print(string(c))
	}

	println("********")
}
