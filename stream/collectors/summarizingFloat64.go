package collectors

// SummarizingFloat64 对数据集进行统计
type SummarizingFloat64 struct {
	Mapper func(interface{}) float64
}

// Float64SummaryStatistics 数据集统计结果_Float64
type Float64SummaryStatistics struct {
	Min     float64
	Max     float64
	Average float64
	Sum     float64
	Count   int
}

// Combine 拼装两个结果集
func (summary *Float64SummaryStatistics) Combine(other *Float64SummaryStatistics) {
	if summary.Count == 0 {
		summary.Max = other.Max
		summary.Min = other.Max
	} else if other.Count != 0 {
		if summary.Max < other.Max {
			summary.Max = other.Max
		}
		if summary.Min > other.Min {
			summary.Min = other.Min
		}
	}
	summary.Sum += other.Sum
	summary.Count += other.Count
}

// Supplier 提供容器
func (summary *SummarizingFloat64) Supplier() func() interface{} {
	return func() interface{} {
		return Float64SummaryStatistics{Min: 0, Max: 0, Average: 0, Sum: 0, Count: 0}
	}
}

// Accumulator 处理函数
func (summary *SummarizingFloat64) Accumulator() func(interface{}, interface{}) interface{} {
	return func(identity interface{}, element interface{}) interface{} {
		mid := identity.(Float64SummaryStatistics)
		value := summary.Mapper(element)
		if mid.Count == 0 {
			mid.Max = value
			mid.Min = value
		} else {
			if mid.Max < value {
				mid.Max = value
			}
			if mid.Min > value {
				mid.Min = value
			}
		}
		mid.Sum += value
		mid.Count++
		return mid
	}
}

// Combiner 组装结果
func (summary *SummarizingFloat64) Combiner() func(interface{}, interface{}) interface{} {
	return func(a interface{}, b interface{}) interface{} {
		midA := a.(Float64SummaryStatistics)
		midB := b.(Float64SummaryStatistics)
		midA.Combine(&midB)
		return midA
	}
}

// Finisher 收尾处理
func (summary *SummarizingFloat64) Finisher() func(interface{}) interface{} {
	return func(identity interface{}) interface{} {
		mid := identity.(Float64SummaryStatistics)
		mid.Average = mid.Sum / float64(mid.Count)
		return mid
	}
}
