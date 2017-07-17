package stream

import (
	"math"
	"runtime"
)

/**
*	Reference:
*	http://www.cnblogs.com/gw811/archive/2012/10/04/2711746.html
 */

// Comparator 比较器
type Comparator struct {
	CompareTo func(interface{}, interface{}) int
}

// Sorted 对数据流操作，升序排序
func (s *Stream) Sorted(comparator *Comparator) *Stream {
	return s.MSorted(comparator)
}

// QSorted 对数据流操作，快排，升序排序
func (s *Stream) QSorted(comparator *Comparator) *Stream {
	qsort(s.list, 0, len(s.list), comparator.CompareTo)
	s.isParallel = false
	return s
}

// MSorted 对数据流操作，归并排序，升序排序
func (s *Stream) MSorted(comparator *Comparator) *Stream {
	aux := make([]interface{}, len(s.list)) //辅助切片
	var ch chan int
	cores := runtime.NumCPU()
	maxDepth := int(math.Log2(float64(cores)) + 1)
	if s.isParallel {
		ch = make(chan int)
		for j := 0; j < cores; j++ {
			go func(idx int) {
				for i := idx; i < len(s.list); i += cores {
					aux[i] = s.list[i]
				}
				ch <- 0
			}(j)
		}
		waitCh(ch, cores)
		go mergeSort(aux, s.list, 0, len(s.list), 0, comparator.CompareTo, 0, maxDepth, ch)
		waitCh(ch, 1)
	} else {
		for i, v := range s.list {
			aux[i] = v
		}
		mergeSort(aux, s.list, 0, len(s.list), 0, comparator.CompareTo, 0, maxDepth, ch)
	}
	s.isParallel = false
	return s
}
func waitCh(ch chan int, num int) {
	if ch != nil {
		for i := 0; i < num; i++ {
			<-ch
		}
	}
}

/**
 * 对指定 int 型数组的指定范围按数字升序进行排序。
 */
func qsort(list []interface{}, fromIndex int, toIndex int, comp func(interface{}, interface{}) int) {
	qsort1(list, fromIndex, toIndex-fromIndex, comp)
}

func qsort1(list []interface{}, off int, len int, comp func(interface{}, interface{}) int) {
	/*
	 * 当待排序的数组中的元素个数小于 7 时，采用插入排序 。
	 *
	 * 尽管插入排序的时间复杂度为O(n^2),但是当数组元素较少时， 插入排序优于快速排序，因为这时快速排序的递归操作影响性能。
	 */
	if len < 7 {
		for i := off; i < len+off; i++ {
			for j := i; j > off && comp(list[j-1], list[j]) > 0; j-- {
				swap(list, j, j-1)
			}
		}
		return
	}
	/*
	 * 当待排序的数组中的元素个数大于 或等于7 时，采用快速排序 。
	 *
	 * Choose a partition element, v
	 * 选取一个划分元，V
	 *
	 * 较好的选择了划分元(基准元素)。能够将数组分成大致两个相等的部分，避免出现最坏的情况。例如当数组有序的的情况下，
	 * 选择第一个元素作为划分元，将使得算法的时间复杂度达到O(n^2).
	 */
	// 当数组大小为size=7时 ，取数组中间元素作为划分元。
	m := off + (len >> 1)
	// 当数组大小 7<size<=40时，取首、中、末 三个元素中间大小的元素作为划分元。
	if len > 7 {
		l := off
		n := off + len - 1
		/*
		 * 当数组大小  size>40 时 ，从待排数组中较均匀的选择9个元素，
		 * 选出一个伪中数做为划分元。
		 */
		if len > 40 {
			s := len / 8
			l = med3(list, l, l+s, l+2*s, comp)
			m = med3(list, m-s, m, m+s, comp)
			n = med3(list, n-2*s, n-s, n, comp)
		}
		// 取出中间大小的元素的位置。
		m = med3(list, l, m, n, comp) // Mid-size, med of 3
	}

	//得到划分元V
	v := list[m]

	// Establish Invariant: v* (<v)* (>v)* v*
	a := off
	b := a
	c := off + len - 1
	d := c
	for true {
		for b <= c && comp(list[b], v) <= 0 {
			if list[b] == v {
				swap(list, a, b)
				a++
			}
			b++
		}
		for c >= b && comp(list[c], v) >= 0 {
			if comp(list[c], v) == 0 {
				swap(list, c, d)
				d--
			}
			c--
		}
		if b > c {
			break
		}
		swap(list, b, c)
		b++
		c--
	}
	// Swap partition elements back to middle
	s := off + len
	n := s
	s = min(a-off, b-a)
	vecswap(list, off, b-s, s)
	s = min(d-c, n-d-1)
	vecswap(list, b, n-s, s)
	// Recursively sort non-partition-elements
	s = b - a
	if s > 1 {
		qsort1(list, off, s, comp)
	}
	s = d - c
	if s > 1 {
		qsort1(list, n-s, s, comp)
	}
}

func min(a int, b int) int {
	if a >= b {
		return b
	}
	return a
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

/**
 * Swaps x[a] with x[b].
 */
func swap(list []interface{}, a int, b int) {
	list[a], list[b] = list[b], list[a]
}

/**
 * Swaps x[a .. (a+n-1)] with x[b .. (b+n-1)].
 */
func vecswap(list []interface{}, a int, b int, n int) {
	for i := 0; i < n; i++ {
		swap(list, a, b)
		a++
		b++
	}
}

/**
 * Returns the index of the median of the three indexed integers.
 */
func med3(list []interface{}, a int, b int, c int, comp func(interface{}, interface{}) int) int {
	if comp(list[a], list[b]) < 0 {
		if comp(list[b], list[c]) < 0 {
			return b
		}
		if comp(list[a], list[c]) < 0 {
			return c
		}
		return a
	}
	if comp(list[b], list[c]) > 0 {
		return b
	}
	if comp(list[a], list[c]) > 0 {
		return c
	}
	return a
}

func rangeCopy(src []interface{}, low int, dest []interface{}, destLow int, length int) {
	for i := 0; i < length; i++ {
		dest[destLow+i] = src[low+i]
	}
}

/**
 * Src is the source array that starts at index 0
 * Dest is the (possibly larger) array destination with a possible offset
 * low is the index in dest to start sorting
 * high is the end index in dest to end sorting
 * off is the offset to generate corresponding low, high in src
 */
func mergeSort(src []interface{}, dest []interface{}, low int, high int, off int, comp func(interface{}, interface{}) int, depth int, maxDepth int, ch chan int) {
	defer func() {
		if ch != nil {
			ch <- 0
		}
	}()
	length := high - low

	// Insertion sort on smallest arrays
	if length < 7 {
		for i := low; i < high; i++ {
			for j := i; j > low && comp(dest[j-1], dest[j]) > 0; j-- {
				swap(dest, j, j-1)
			}
		}
		return
	}

	// Recursively sort halves of dest into src
	destLow := low
	destHigh := high
	low += off
	high += off
	/*
	 *  >>>：无符号右移运算符
	 *  expression1 >>> expresion2：expression1的各个位向右移expression2
	 *  指定的位数。右移后左边空出的位数用0来填充。移出右边的位被丢弃。
	 *  例如：-14>>>2；  结果为：1073741820
	 */
	mid := (low + high) >> 1
	if ch != nil {
		if depth < maxDepth {
			ch2 := make(chan int)
			go mergeSort(dest, src, low, mid, -off, comp, depth+1, maxDepth, ch2)
			go mergeSort(dest, src, mid, high, -off, comp, depth+1, maxDepth, ch2)
			waitCh(ch2, 2)
		} else {
			mergeSort(dest, src, low, mid, -off, comp, depth+1, maxDepth, nil)
			mergeSort(dest, src, mid, high, -off, comp, depth+1, maxDepth, nil)
		}
	} else {
		mergeSort(dest, src, low, mid, -off, comp, depth+1, maxDepth, ch)
		mergeSort(dest, src, mid, high, -off, comp, depth+1, maxDepth, ch)
	}

	// If list is already sorted, just copy from src to dest. This is an
	// optimization that results in faster sorts for nearly ordered lists.
	if comp(src[mid-1], src[mid]) <= 0 {
		rangeCopy(src, low, dest, destLow, length)
		return
	}

	// Merge sorted halves (now in src) into dest
	for i, p, q := destLow, low, mid; i < destHigh; i++ {
		if q >= high || p < mid && comp(src[p], src[q]) <= 0 {
			dest[i] = src[p]
			p++
		} else {
			dest[i] = src[q]
			q++
		}
	}
}
