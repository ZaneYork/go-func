package stream

// Stream 数据流
type Stream struct {
	isParallel bool
	list       []interface{}
}

// Parallel 转为并发数据流
func (s *Stream) Parallel() *Stream {
	s.isParallel = true
	return s
}

// Sequential 转为单线程数据流
func (s *Stream) Sequential() *Stream {
	s.isParallel = false
	return s
}

// NewStream 新建数据流
func NewStream(list []interface{}) *Stream {
	return &Stream{list: list, isParallel: false}
}

// NewParallelStream 新建并发数据流
func NewParallelStream(list []interface{}) *Stream {
	return &Stream{list: list, isParallel: true}
}

// Count 统计数量
func (s *Stream) Count() int {
	return len(s.list)
}

// ToList 转为列表
func (s *Stream) ToList() []interface{} {
	return s.list
}

// FindFirst 取得第一条
func (s *Stream) FindFirst() interface{} {
	if len(s.list) > 0 {
		return s.list[0]
	}
	return nil
}
