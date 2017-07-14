package main

import (
	. "github.com/go-func/stream"
	"github.com/go-func/stream/collectors"
	"fmt"
)

func main() {
	list := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		list[i] = i * 2
	}
	tmp := NewParallelStream(list).Filter(func(item interface{}) bool {
		if item.(int) < 100 {
			return false
		}
		return true
	}).Map(func(item interface{}) interface{} {
		value := item.(int)
		value /= 2
		return value
	}).Collect(&collectors.GroupingBy{Classifier: func(item interface{}) interface{} {
		return item.(int)%2 == 0
	}, DownStream: &collectors.PartitioningBy{Predicate: func(item interface{}) bool {
		return item.(int)%3 == 0
	}, DownStream: &collectors.SummarizingFloat64{Mapper: func(item interface{}) float64 {
		return float64(item.(int))
	}}}})
	/*
		.Collect(&collectors.PartitioningBy{Predicate: func(item interface{}) bool {
			return item.(int)%3 == 0
		}})
		.Collect(&collectors.GroupingBy{Classifier: func(item interface{}) interface{} {
			return item.(int)%3 == 0
		}})
		.Collect(&collectors.Joining{Prefix: "[", Suffix: "]", Delimiter: ","})
		.Collect(&collectors.Counting{})
		.Collect(&collectors.AveragingFloat64{Mapper: func(item interface{}) float64 {
			return float64(item.(int))
		}})
		.Collect(&collectors.ToMap{KeyMapper: func(item interface{}) interface{} {
			return item.(int) * 2
		}, ValueMapper: func(item interface{}) interface{} {
			return item
		}})
		.Collect(&collectors.SummarizingFloat64{Mapper: func(item interface{}) float64 {
			return float64(item.(int))
		}})
	*/
	fmt.Println(tmp)
}
