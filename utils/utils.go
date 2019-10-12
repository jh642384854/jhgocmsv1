package utils

import (
	"fmt"
	"crypto/md5"
	"github.com/sirupsen/logrus"
)

/**
	Logger的初始化示例
 */
func NewLogrusLogger() *logrus.Logger  {
	logrusInstance := logrus.New()
	return logrusInstance
}
/**
	对字符串进行MD5加密处理
 */
func EncodePassword(password string) string  {
	if password != ""{
		data := []byte(password)
		md5hash := md5.Sum(data)
		md5str := fmt.Sprintf("%x",md5hash)
		return md5str
	}
	return ""
}
/**
	对字符串的切片元素去重处理
 */
func RemoveReplicaSliceString(slc []string) []string {
	result := make([]string, 0)
	tempMap := make(map[string]bool, len(slc))
	for _, e := range slc{
		if tempMap[e] == false{
			tempMap[e] = true
			result = append(result, e)
		}
	}
	return result
}