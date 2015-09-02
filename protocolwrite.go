package gr

import (
	"bufio"
	"net"
	"strconv"
)

func writeRead(cmds [][]byte, conn *net.TCPConn) (reply *redisResponse, err error) {

	size := 0
	for _, c := range cmds {
		size += len(c)
	}

	writer := bufio.NewWriterSize(conn, size)

	if err = write(cmds, writer); err == nil {
		if err = writer.Flush(); err == nil {
			reader := bufio.NewReader(conn)
			reply, err = read(reader)
		}
	}

	return
}

func write(cmds [][]byte, writer *bufio.Writer) (err error) {

	cmdsCount := len(cmds)
	strCmdsCount := strconv.Itoa(cmdsCount)

	writer.WriteRune(tArray)
	writer.WriteString(strCmdsCount)
	writer.Write(_crlf)

	for _, c := range cmds {
		writer.Write(c)
	}

	return
}
