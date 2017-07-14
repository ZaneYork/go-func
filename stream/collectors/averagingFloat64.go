package collectors

// AveragingFloat64 统计结果集平均值
type AveragingFloat64 struct {
	Mapper func(interface{}) float64
}

type aveMidResult struct {
	Sum   float64
	Count int
}

// Supplier 提供容器
func (ave *AveragingFloat64) Supplier() func() interface{} {
	return func() interface{} {
		return aveMidResult{Sum: 0, Count: 0}
	}
}

// Accumulator 处理函数
func (ave *AveragingFloat64) Accumulator() func(interface{}, interface{}) interface{} {
	return func(identity interface{}, element interface{}) interface{} {
		mid := identity.(aveMidResult)
		mid.Count++
		mid.Sum += ave.Mapper(element)
		return mid
	}
}

// Combiner 组装结果
func (ave *AveragingFloat64) Combiner() func(interface{}, interface{}) interface{} {
	return func(a interface{}, b interface{}) interface{} {
		midA := a.(aveMidResult)
		midB := b.(aveMidResult)
		midA.Sum += midB.Sum
		midA.Count += midB.Count
		return midA
	}
}

// Finisher 收尾处理
func (ave *AveragingFloat64) Finisher() func(interface{}) interface{} {
	return func(identity interface{}) interface{} {
		mid := identity.(aveMidResult)
		return mid.Sum / float64(mid.Count)
	}
}
