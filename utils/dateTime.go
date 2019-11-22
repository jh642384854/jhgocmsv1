package utils

import (
	"github.com/chenhg5/collection"
	"jhgocms/constant"
	"time"
)

/**
	将月份的英文单词的字符串转换为01-12这样的字符串
 */
func ParseMonths(month string) string {
	var monthMaps = map[string]interface{}{"January": "01", "February": "02", "March": "03", "April": "04", "May": "05", "June": "06", "July": "07", "August": "08", "September": "09", "October": "10", "November": "11", "December": "12"}
	if (collection.Collect(monthMaps).Has(month)) {
		return monthMaps[month].(string)
	}
	return ""
}

/**
	获取当前时间的时间戳
 */
func GetNowTimestamp() int64 {
	return time.Now().Unix()
}

/**
	将指定时间戳转换为日期格式
 */
func ConverTimesstampToDatetime(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(constant.TimeFormat)
}