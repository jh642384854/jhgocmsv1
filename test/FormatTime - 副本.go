package utils

import (
	"database/sql/driver"
	"fmt"
	"jhgocms/constant"
	"time"
)
/**
	参考：
	http://www.axiaoxin.com/article/241/
	https://www.cnblogs.com/mrylong/p/11326792.html
 */
type FormatTime struct {
	time.Time
}
//这个是在进行json输出的时候会自动做处理
func (t FormatTime) MarshalJSON() ([]byte, error) {
	fmt.Println("FormatTime MarshalJSON()")
	formatted := fmt.Sprintf("\"%s\"", t.Format(constant.TimeFormat))
	return []byte(formatted), nil
}

//这里有个疑问，为什么这里必须要定义为Value()和Scan()这样名称的方法？
func (t FormatTime) Value() (driver.Value, error) {
	fmt.Println("FormatTime Value()")
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *FormatTime) Scan(v interface{}) error {
	fmt.Println("FormatTime Scan()")
	value, ok := v.(time.Time)
	//value := time.Unix(v.(int64), 0)
	if ok {
		*t = FormatTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
