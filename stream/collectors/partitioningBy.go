package collectors

import . "github.com/ZaneYork/go-func/stream"

// PartitioningBy 按照指定的分区
type PartitioningBy struct {
	Predicate  func(interface{}) bool
	DownStream Collector
}

// Supplier 提供容器
func (p *PartitioningBy) Supplier() func() interface{} {
	return func() interface{} {
		return make(map[bool][]interface{})
	}
}

// Accumulator 处理函数
func (p *PartitioningBy) Accumulator() func(interface{}, interface{}) interface{} {
	return func(identity interface{}, element interface{}) interface{} {
		ret := identity.(map[bool][]interface{})
		k := p.Predicate(element)
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
func (p *PartitioningBy) Combiner() func(interface{}, interface{}) interface{} {
	return func(a interface{}, b interface{}) interface{} {
		mapA := a.(map[bool][]interface{})
		mapB := b.(map[bool][]interface{})
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
func (p *PartitioningBy) Finisher() func(interface{}) interface{} {
	return func(identity interface{}) interface{} {
		if p.DownStream != nil {
			result := identity.(map[bool][]interface{})
			downStreamResult := make(map[bool]interface{})
			for k, v := range result {
				downStreamResult[k] = NewStream(v).Collect(p.DownStream)
			}
			return downStreamResult
		}
		return identity
	}
}
