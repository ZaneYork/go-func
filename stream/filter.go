package stream

import "runtime"

// Filter 对任意数组操作，将不合格的元素踢出去
func (s *Stream) Filter(predicate func(interface{}) bool) *Stream {
	ret := make([]interface{}, len(s.list))
	size := 0
	if s.isParallel {
		count := runtime.NumCPU()
		ch1 := make(chan interface{}, count)
		ch2 := make(chan interface{})
		for i := 0; i < count; i++ {
			go func() {
				for true {
					if value, ok := <-ch1; !ok {
						break
					} else {
						if predicate(value) {
							ch2 <- value
						} else {
							ch2 <- nil
						}
					}
				}
			}()
		}
		go func() {
			for _, v := range s.list {
				ch1 <- v
			}
			close(ch1)
		}()
		for i := 0; i < len(s.list); i++ {
			value := <-ch2
			if value != nil {
				ret[size] = value
				size++
			}
		}
	} else {
		for _, v := range s.list {
			if predicate(v) {
				ret[size] = v
				size++
			}
		}
	}
	s.list = ret[:size]
	return s
}
