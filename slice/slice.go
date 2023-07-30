package slice

import "log"

type Slice[T Number] struct {
	slice []T
}

// Add 在切片任意位置插入元素
func (s *Slice[T]) Add(item T, idx int) {
	if s.slice == nil {
		s.slice = make([]T, 0, 64)
	}
	if idx < 0 || idx > len(s.slice) {
		panic("invalid index")
	}
	tmp := append([]T{}, s.slice[idx:]...)
	s.slice = append(s.slice[:idx], item)
	s.slice = append(s.slice, tmp...)
}

// AddFirst 在切片头部插入元素
func (s *Slice[T]) AddFirst(item T) {
	s.Add(item, 0)
}

// AddLast 在切片尾部插入元素
func (s *Slice[T]) AddLast(item T) {
	s.Add(item, len(s.slice))
}

// DelByIndex 通过下标删除元素
func (s *Slice[T]) DelByIndex(idx int) T {
	if s.slice == nil {
		log.Panicf("slice has no element\n")
	}

	s.slice = shrink[T](s.slice)

	if idx < 0 || idx > len(s.slice)-1 {
		panic("invalid index")
	}
	tmp := s.slice[idx]
	s.slice = append(s.slice[:idx], s.slice[idx+1:]...)
	return tmp
}

// DelItem 删除指定元素
func (s *Slice[T]) DelItem(item T) T {
	index := s.include(item)
	if index == -1 {
		log.Panicf("slice does not include item %v\n", item)
	}

	return s.DelByIndex(index)
}

// 判断切片是否包含元素
func (s *Slice[T]) include(item T) int {
	for i, val := range s.slice {
		if val == item {
			return i
		}
	}
	return -1
}

func calcCap(c, l int) (int, bool) {
	if c <= 256 {
		return c, false
	} else {
		if l < c/4 {
			return c / 2, true
		}
	}

	return c, false
}

func shrink[T Number](src []T) []T {
	n, isShrink := calcCap(cap(src), len(src))

	if !isShrink {
		return src
	}

	tmp := make([]T, 0, n)
	tmp = append(tmp, src...)
	return tmp
}
