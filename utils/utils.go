package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/chenhg5/collection"
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
/**
	判断切片中是否包含某个元素
 */
func SliceContainSomeOne(slc []string,target string) bool {
	return collection.Collect(slc).Contains(target)
}