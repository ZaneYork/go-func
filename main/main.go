package main

import (
	. "github.com/go-func/stream"
)

// func NewStopWatch() func() time.Duration {
// 	var QPCTimer func() func() time.Duration
// 	QPCTimer = func() func() time.Duration {
// 		lib, _ := syscall.LoadLibrary("kernel32.dll")
// 		qpc, _ := syscall.GetProcAddress(lib, "QueryPerformanceCounter")
// 		qpf, _ := syscall.GetProcAddress(lib, "QueryPerformanceFrequency")
// 		if qpc == 0 || qpf == 0 {
// 			return nil
// 		}
// 		var freq, start uint64
// 		syscall.Syscall(qpf, 1, uintptr(unsafe.Pointer(&freq)), 0, 0)
// 		syscall.Syscall(qpc, 1, uintptr(unsafe.Pointer(&start)), 0, 0)
// 		if freq <= 0 {
// 			return nil
// 		}
// 		freqns := float64(freq) / 1e9
// 		return func() time.Duration {
// 			var now uint64
// 			syscall.Syscall(qpc, 1, uintptr(unsafe.Pointer(&now)), 0, 0)
// 			return time.Duration(float64(now-start) / freqns)
// 		}
// 	}
// 	var StopWatch func() time.Duration
// 	if StopWatch = QPCTimer(); StopWatch == nil {
// 		// Fallback implementation
// 		start := time.Now()
// 		StopWatch = func() time.Duration { return time.Since(start) }
// 	}
// 	return StopWatch
// }

func main() {
	list := make([]interface{}, 80000000)
	for i := 0; i < 80000000; i++ {
		list[i] = i * 2
	}
	// watch := NewStopWatch()
	for i := 0; i < 1; i++ {
		NewParallelStream(list).Sorted(&Comparator{CompareTo: func(a interface{}, b interface{}) int {
			aInt := a.(int)
			bInt := b.(int)
			if aInt < bInt {
				return 1
			} else if aInt > bInt {
				return -1
			}
			return 0
		}})
		/*
			.Filter(func(item interface{}) bool {
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
		*/
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
			.ForEach(func(item interface{}) {
				fmt.Println(item)
			})
		*/

	}
	// fmt.Println(watch())
	//fmt.Println(tmp)
}
