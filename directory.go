package pkg

import (
	"os"
)

// @title    PathExists
// @description   文件目录是否存在
// @auth                     （2020/04/05  20:22）
// @param     path            string
// @return    err             error

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// @title    checkFileIsExist
// @description   查看文件是否存在
// @auth                     （2020/04/05  20:22）
// @param     path            string
// @return    bool             bool
// 判断文件是否存在  存在返回 true 不存在返回false

func CheckFileIsExist(path string) bool {
	var exist = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
