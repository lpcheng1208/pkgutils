package tokenhelper

import (
	"errors"
	"fmt"
	"github.com/tal-tech/go-zero/core/stringx"
	"time"
)

func MakeToken(userId string) string {
	gid := stringx.RandId()
	nowTime := time.Now().UnixNano()/1e6
	return fmt.Sprintf("%s%s%v", userId, gid, nowTime)
}

func ParseTokenToUserId(token string) (string, error) {
	userId := token[:32]
	if userId == ""{
		return "", errors.New("token err")
	}
	return userId, nil
}