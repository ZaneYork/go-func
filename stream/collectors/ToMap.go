package collectors

// ToMap 按照指定的规则组装为Map
type ToMap struct {
	KeyMapper   func(interface{}) interface{}
	ValueMapper func(interface{}) interface{}
}

// Supplier 提供容器
func (tm *ToMap) Supplier() func() interface{} {
	return func() interface{} {
		return make(map[interface{}]interface{})
	}
}

// Accumulator 处理函数
func (tm *ToMap) Accumulator() func(interface{}, interface{}) interface{} {
	return func(identity interface{}, element interface{}) interface{} {
		ret := identity.(map[interface{}]interface{})
		k := tm.KeyMapper(element)
		v := tm.ValueMapper(element)
		ret[k] = v
		return ret
	}
}

// Combiner 组装结果
func (tm *ToMap) Combiner() func(interface{}, interface{}) interface{} {
	return func(a interface{}, b interface{}) interface{} {
		mapA := a.(map[interface{}]interface{})
		mapB := b.(map[interface{}]interface{})
		for k, v := range mapB {
			if _, ok := mapA[k]; !ok {
				mapA[k] = v
			}
		}
		return mapA
	}
}

// Finisher 收尾处理
func (tm *ToMap) Finisher() func(interface{}) interface{} {
	return func(identity interface{}) interface{} {
		return identity
	}
}
