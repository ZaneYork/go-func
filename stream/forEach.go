package stream

import "runtime"

//ForEach 对任意数组操作，针对每个元素执行一遍指定的行为
func (s *Stream) ForEach(action func(interface{})) {
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
						action(value)
					}
				}
				ch2 <- 0
			}()
		}
		go func() {
			for _, v := range s.list {
				ch1 <- v
			}
			close(ch1)
		}()
		for i := 0; i < count; i++ {
			<-ch2
		}
	} else {
		for _, v := range s.list {
			action(v)
		}
	}
}
