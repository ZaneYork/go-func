package stream

import "runtime"

//Map 对任意数组操作，将A包装成B
func (s *Stream) Map(mappper func(interface{}) interface{}) *Stream {
	ret := make([]interface{}, len(s.list))
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
						ch2 <- mappper(value)
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
			ret[i] = <-ch2
		}
	} else {
		for i, v := range s.list {
			ret[i] = mappper(v)
		}
	}
	s.list = ret
	return s
}
