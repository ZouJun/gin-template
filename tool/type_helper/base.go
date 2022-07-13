package type_helper

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"reflect"
	"strconv"
	"unicode"
)

func Float64TwoDecimal(val float64) float64 {
	v, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", val), 64)
	return v
}

func Float64ToString(v float64) string {
	return strconv.FormatFloat(v, 'E', -1, 64)
}

func Int64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func StringToFloat64(v string) float64 {
	v1, _ := strconv.ParseFloat(v, 64)
	return v1
}

//直接保存2位小数
func TwoDecimal(v float64) float64 {
	return math.Trunc(v*1e2) * 1e-2
}

//structToMap
func StructToMap(structObj interface{}) map[string]interface{} {
	t := reflect.TypeOf(structObj)
	v := reflect.ValueOf(structObj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[Camel2Case(t.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}

//驼峰转下划线---------------------------
// 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}

// 驼峰式写法转为下划线写法
func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

//驼峰转下划线---------------------------
