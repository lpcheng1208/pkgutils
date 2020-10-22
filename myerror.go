package pkgutils

import (
	"errors"
	"fmt"
	"strings"
)

func MakeMyRespError(code int32, msg string) error {
	errStr := fmt.Sprintf("|%d|%s", code, msg)
	err := errors.New(errStr)
	return err
}

func ParseMyRespError(errStr string) (code int32, msg string) {
	splitSlice := strings.Split(errStr, "|")
	lenNum := len(splitSlice)
	if lenNum < 2{
		return
	}
	msg = splitSlice[lenNum - 1]
	s, _ := StringToInt(splitSlice[lenNum - 2])
	code = int32(s)
	return
}
