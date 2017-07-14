package stream

import "runtime"

// Collector 收集器接口，对于所有的并发流，定义打包行为
type Collector interface {
	Supplier() func() interface{}
	Accumulator() func(interface{}, interface{}) interface{}
	Combiner() func(interface{}, interface{}) interface{}
	Finisher() func(interface{}) interface{}
}

// Collect 收集器，对于所有的并发流，打包
func (s *Stream) Collect(collector Collector) interface{} {
	identity := collector.Supplier()()
	if s.isParallel {
		count := runtime.NumCPU()
		ch1 := make(chan interface{}, count)
		ch2 := make(chan interface{})
		for i := 0; i < count; i++ {
			go func() {
				result := collector.Supplier()()
				for true {
					if value, ok := <-ch1; !ok {
						break
					} else {
						result = collector.Accumulator()(result, value)
					}
				}
				ch2 <- result
			}()
		}
		go func() {
			for _, v := range s.list {
				ch1 <- v
			}
			close(ch1)
		}()
		for i := 0; i < count; i++ {
			tmp := <-ch2
			identity = collector.Combiner()(identity, tmp)
		}
	} else {
		for _, v := range s.list {
			identity = collector.Accumulator()(identity, v)
		}
	}
	identity = collector.Finisher()(identity)
	return identity
}
