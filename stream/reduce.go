package stream

//Reduce 对任意数组操作，对数组进行降维打击
func (s *Stream) Reduce(identity interface{}, accumulator func(interface{}, interface{}) interface{}) interface{} {
	for _, v := range s.list {
		identity = accumulator(identity, v)
	}
	return identity
}
