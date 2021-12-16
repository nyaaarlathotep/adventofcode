package DataStruct

import "strconv"

type Pipe struct {
	index      int
	raw        string
	cacheIndex int
	cache      string
}

func NewPipe(input string) *Pipe {
	p := &Pipe{
		index:      1,
		raw:        input,
		cacheIndex: 0,
		cache:      getFourBits(input[0:1]),
	}
	return p
}

func (p *Pipe) NextChar() string {
	if p.cacheIndex == 4 {
		p.cacheIndex = 0
		p.cache = getFourBits(p.raw[p.index : p.index+1])
		p.index++
	}
	res := string(p.cache[p.cacheIndex])
	p.cacheIndex++
	return res
}

func getFourBits(str string) string {
	ten, _ := strconv.ParseInt(str, 16, 32)
	bits := strconv.FormatInt(ten, 2)

	zeroNum := 4 - len(bits)
	for i := 0; i < zeroNum; i++ {
		bits = "0" + bits
	}

	return bits
}

func (p *Pipe) GetChars(count int) string {
	res := ""
	for i := 0; i < count; i++ {
		res = res + p.NextChar()
	}
	return res
}
