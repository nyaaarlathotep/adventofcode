package DataStruct

type exists struct{}

type Set struct {
	m map[interface{}]exists
}

func NewSet(items ...interface{}) *Set {
	s := &Set{}
	s.m = make(map[interface{}]exists)
	s.Add(items...)
	return s
}

// Add 是的，只用基本类型，也就是说，能用==比较的类型
func (s *Set) Add(items ...interface{}) {
	for _, item := range items {
		s.m[item] = exists{}
	}
}

func (s *Set) Remove(item interface{}) {
	delete(s.m, item)
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) List() [][2]int {
	var list [][2]int
	for item := range s.m {
		list = append(list, item.([2]int))
	}
	return list
}
