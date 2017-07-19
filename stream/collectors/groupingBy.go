package collectors

import . "github.com/ZaneYork/go-func/stream"

// GroupingBy 按照指定的规则分组
type GroupingBy struct {
	Classifier func(interface{}) interface{}
	DownStream Collector
}

// Supplier 提供容器
func (g *GroupingBy) Supplier() func() interface{} {
	return func() interface{} {
		return make(map[interface{}][]interface{})
	}
}

// Accumulator 处理函数
func (g *GroupingBy) Accumulator() func(interface{}, interface{}) interface{} {
	return func(identity interface{}, element interface{}) interface{} {
		ret := identity.(map[interface{}][]interface{})
		k := g.Classifier(element)
		if value, ok := ret[k]; ok {
			ret[k] = append(value, element)
		} else {
			list := make([]interface{}, 1)
			list[0] = element
			ret[k] = list
		}
		return ret
	}
}

// Combiner 组装结果
func (g *GroupingBy) Combiner() func(interface{}, interface{}) interface{} {
	return func(a interface{}, b interface{}) interface{} {
		mapA := a.(map[interface{}][]interface{})
		mapB := b.(map[interface{}][]interface{})
		for k, v := range mapB {
			if value, ok := mapA[k]; ok {
				mapA[k] = append(value, v...)
			} else {
				mapA[k] = v
			}
		}
		return mapA
	}
}

// Finisher 收尾处理
func (g *GroupingBy) Finisher() func(interface{}) interface{} {
	return func(identity interface{}) interface{} {
		if g.DownStream != nil {
			result := identity.(map[interface{}][]interface{})
			downStreamResult := make(map[interface{}]interface{})
			for k, v := range result {
				downStreamResult[k] = NewStream(v).Collect(g.DownStream)
			}
			return downStreamResult
		}
		return identity
	}
}
