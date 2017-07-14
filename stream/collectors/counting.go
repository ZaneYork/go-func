package collectors

// Counting 统计结果集数量
type Counting struct {
}

// Supplier 提供容器
func (c *Counting) Supplier() func() interface{} {
	return func() interface{} {
		return int(0)
	}
}

// Accumulator 处理函数
func (c *Counting) Accumulator() func(interface{}, interface{}) interface{} {
	return func(identity interface{}, element interface{}) interface{} {
		return identity.(int) + 1
	}
}

// Combiner 组装结果
func (c *Counting) Combiner() func(interface{}, interface{}) interface{} {
	return func(a interface{}, b interface{}) interface{} {
		return a.(int) + b.(int)
	}
}

// Finisher 收尾处理
func (c *Counting) Finisher() func(interface{}) interface{} {
	return func(identity interface{}) interface{} {
		return identity
	}
}
