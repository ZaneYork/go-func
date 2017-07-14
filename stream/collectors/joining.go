package collectors

import (
	"bytes"
	"fmt"
	"strconv"
)

// Joining 按照指定的连接成字符串
type Joining struct {
	Delimiter string
	Suffix    string
	Prefix    string
}

// Supplier 提供容器
func (j *Joining) Supplier() func() interface{} {
	return func() interface{} {
		var buffer bytes.Buffer
		return buffer
	}
}

// Accumulator 处理函数
func (j *Joining) Accumulator() func(interface{}, interface{}) interface{} {
	return func(identity interface{}, element interface{}) interface{} {
		buffer := identity.(bytes.Buffer)
		switch element.(type) {
		case string:
			buffer.WriteString(element.(string))
		case int:
			buffer.WriteString(strconv.Itoa(element.(int)))
		default:
			str2 := fmt.Sprintf("%v", element)
			buffer.WriteString(str2)
		}
		buffer.WriteString(j.Delimiter)
		return buffer
	}
}

// Combiner 组装结果
func (j *Joining) Combiner() func(interface{}, interface{}) interface{} {
	return func(a interface{}, b interface{}) interface{} {
		bufferA := a.(bytes.Buffer)
		bufferB := b.(bytes.Buffer)
		bufferA.WriteString(bufferB.String())
		return bufferA
	}
}

// Finisher 收尾处理
func (j *Joining) Finisher() func(interface{}) interface{} {
	return func(identity interface{}) interface{} {
		buffer := identity.(bytes.Buffer)
		str := buffer.String()
		if len(str) > 0 {
			str = str[:len(str)-len(j.Delimiter)]
		}
		ret := j.Prefix + str + j.Suffix
		return ret
	}
}
