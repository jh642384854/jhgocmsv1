package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"jhgocms/config"
	"log"
	"math/big"
	mathrand "math/rand"
	"os"
	"reflect"
	"time"

	"github.com/chenhg5/collection"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"gopkg.in/gomail.v2"
)

/**
	Logger的初始化示例
 */
func NewLogrusLogger() *logrus.Logger {
	logrusInstance := logrus.New()
	return logrusInstance
}

/**
	对字符串进行MD5加密处理
 */
func EncodePassword(password string) string {
	if password != "" {
		data := []byte(password)
		md5hash := md5.Sum(data)
		md5str := fmt.Sprintf("%x", md5hash)
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
	for _, e := range slc {
		if tempMap[e] == false {
			tempMap[e] = true
			result = append(result, e)
		}
	}
	return result
}
//将是slice的接口类型转换为slice类型
func ToSlice(arr interface{}) []interface{}  {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice{
		panic("toslice arr no slice")
	}
	l := v.Len()
	ret := make([]interface{},l)
	for i:= 0;i<l ;i++  {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

/**
	判断切片中是否包含某个元素
 */
func SliceContainSomeOne(slc []string, target string) bool {
	return collection.Collect(slc).Contains(target)
}

//得到附件上传的地址
func GetUploadPath() string {
	var basePath = config.SysConfig.UploadPath
	var timeNow = time.Now()
	var subPath = fmt.Sprintf("%d%s%d", timeNow.Year(), ParseMonths(timeNow.Month().String()), timeNow.Day())
	var allPath = basePath + "/" + subPath + "/"
	var aferoInstanc = afero.NewOsFs()
	hasExits, err := afero.DirExists(aferoInstanc, allPath)
	if err != nil {
		log.Fatal("afero.DirExists err:" + err.Error())
	}
	if !hasExits {
		if err := aferoInstanc.MkdirAll(allPath, os.ModePerm); err != nil {
			log.Fatal("afero.MkdirAll err:" + err.Error())
		}
	}
	return allPath
}

/**
	生成指定长度的字符串
 */
func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

/**
	生成6为的随机数字验证码
 */
func CreateCaptcha() string {
	return fmt.Sprintf("%06v", mathrand.New(mathrand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

/**
	发送邮件
	sendEmail
 */
func SendEmail(sendEmail, receiveEmail, title, body, smtpServer, password string, port int) (bool, error) {
	mail := gomail.NewMessage()
	mail.SetHeader("From", sendEmail)
	mail.SetHeader("To", receiveEmail)
	mail.SetHeader("Subject", title)
	mail.SetBody("text/html", body)
	//smtp.qq.com 的端口可以是587或是465
	d := gomail.NewDialer(smtpServer, port, sendEmail, password)
	if err := d.DialAndSend(mail); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
