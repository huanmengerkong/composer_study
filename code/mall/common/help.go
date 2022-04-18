package common

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func GetInt64(nums interface{}) int {
	var num int = 0
	fmt.Printf("这个类型是%s",reflect.TypeOf(nums))
	switch nums.(type) {
	case int64:
		num = int(nums.(int64))
	case string:
		num, _ = strconv.Atoi(nums.(string))
	case int32:
		num = int(nums.(int32))
	case json.Number:
		num64,_:=nums.(json.Number).Int64()
		num = int(num64)
	}
	return num
}
