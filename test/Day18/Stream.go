package Day18

type Stream struct {
	material string
	index    int
}

func NewStream(input string) *Stream {
	s := &Stream{
		material: input,
		index:    0,
	}
	return s
}

func (s *Stream) NextChar() string {
	res := s.material[s.index]
	s.index++
	return string(res)
}

func (s *Stream) GetChars(length int) string {
	res := ""
	for i := 0; i < length; i++ {
		res = res + s.NextChar()
	}
	return res
}

func (s *Stream) GetUntilNext(charLeft string, charRight string) string {
	res := ""
	counter := 1

	for true {
		temp := s.NextChar()
		if temp == charLeft {
			counter++
		}
		if temp == charRight {
			counter--
		}
		if counter == 0 {
			break
		}
		res = res + temp

	}
	return res
}

func (s *Stream) GetUntilEnd() string {
	res := ""
	for s.hasNext() {
		temp := s.NextChar()
		res = res + temp
	}
	return res
}

func (s *Stream) hasNext() bool {
	if s.index < len(s.material) {
		return true
	}
	return false

}
