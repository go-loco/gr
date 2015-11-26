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

/*
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
*/
