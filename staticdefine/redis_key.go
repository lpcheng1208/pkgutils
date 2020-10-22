package staticdefine

import "fmt"

// user
const (
	KeyUserToken =  "titi#userToken:"
)

func GetUserTokenKey(userid interface{}) string {
	return fmt.Sprintf("%s%v", KeyUserToken, userid)
}