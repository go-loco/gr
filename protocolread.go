package gr

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

func readString(rr *redisResponse, err error) (string, error) {

	if err != nil {
		return "", err
	} else if rr.resp == nil {
		return "", NilErr
	}

	return string(rr.resp), err
}

func readStringArray(rr *redisResponse, err error) ([]string, error) {

	if err == nil {

		str := make([]string, 0, rr.arraySize)
		nextResponse := rr.next

		for i := 0; i < rr.arraySize; i++ {
			if nextResponse.redisType == tArray {

				s, err := readStringArray(nextResponse, err)
				if err != nil {
					break
				}

				str = append(str, s...)

			} else {
				str = append(str, string(nextResponse.resp))
				nextResponse = nextResponse.next
			}

		}

		return str, err
	}

	return nil, err
}

func readInt64(rr *redisResponse, err error) (int64, error) {
	if err != nil || rr.resp == nil {
		return -1, err
	}

	return strconv.ParseInt(string(rr.resp), 10, 64)
}

func readFloat64(rr *redisResponse, err error) (float64, error) {
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(string(rr.resp), 64)
}

func readBool(rr *redisResponse, err error) (bool, error) {

	rs := false

	if err == nil && string(rr.resp) == "1" {
		rs = true
	}

	return rs, err
}

type redisResponse struct {
	resp      []byte
	redisType byte
	arraySize int
	next      *redisResponse
}

func read(reader *bufio.Reader) (reply *redisResponse, err error) {

	reply = new(redisResponse)
	resp, line, err := readLine(reader)

	if err == nil {

		switch resp {
		case tInteger:
			reply.resp = line
			reply.redisType = tInteger
		case tString:
			reply.resp = line
			reply.redisType = tString
		case tBulkString:
			reply.resp, err = readBulk(line, reader)
			reply.redisType = tBulkString
		case tError:
			err = errors.New(string(line))
		case tArray:
			var size int

			if size, err = readSize(line); err == nil {

				reply.redisType = tArray
				reply.arraySize = size
				tail := reply

				for i := 0; i < size; i++ {
					if tail.next, err = read(reader); err != nil {
						break
					}
					tail = tail.next
				}
			}

		} //end switch

	} // end if

	return
}

func readLine(reader *bufio.Reader) (resp byte, line []byte, err error) {

	line, err = reader.ReadSlice('\n')

	if err == nil {
		resp = line[0]
		line = line[1 : len(line)-2]
	}

	return
}

func readSize(line []byte) (int, error) {
	return strconv.Atoi(string(line))
}

func readBulk(line []byte, reader *bufio.Reader) (reply []byte, err error) {

	var size int
	if size, err = readSize(line); err == nil {
		// If the requested value does not exist
		// reply will be nil (not exsitance of the value)
		if size < 0 {
			return
		}

		buf := make([]byte, size+2)
		if _, err := io.ReadFull(reader, buf); err == nil {
			reply = buf[:size]
		}

	}

	return
}
