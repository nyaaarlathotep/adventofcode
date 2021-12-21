package Day21

type player struct {
	score int
	pos   int
}

func newPlayer(posStart int) *player {
	p := &player{
		score: 0,
		pos:   posStart,
	}
	return p
}

func (p *player) movePawn(counter *int) {
	p.pos = p.pos + nextDiceNum(counter) + nextDiceNum(counter) + nextDiceNum(counter)
	if p.pos >= 10 {
		p.pos = p.pos % 10
	}
	if p.pos == 0 {
		p.pos = 10
	}
	p.score = p.score + p.pos
}

func (p *player) win() bool {
	return p.score >= 1000
}
