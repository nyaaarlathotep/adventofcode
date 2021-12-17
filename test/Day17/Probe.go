package Day17

type Probe struct {
	posX      int
	posY      int
	velocityX int
	velocityY int
	MaxHeight int
	Hit       bool
}

func NewProbe(vx int, vy int) *Probe {
	p := &Probe{
		posX:      0,
		posY:      0,
		velocityX: vx,
		velocityY: vy,
		MaxHeight: 0,
		Hit:       false,
	}
	return p
}

func (p *Probe) fly(area *[][]int) {
	if p.posY > p.MaxHeight {
		p.MaxHeight = p.posY
	}

	p.posX += p.velocityX
	p.Hit = testHit(p.posX, p.posY, area)
	p.posY += p.velocityY
	p.Hit = testHit(p.posX, p.posY, area)

	if p.velocityX > 0 {
		p.velocityX--
	} else if p.velocityX < 0 {
		p.velocityX++
	}
	p.velocityY--
}

func testHit(x int, y int, area *[][]int) bool {
	if x >= (*area)[0][0] && x <= (*area)[0][1] && y >= (*area)[1][0] && y <= (*area)[1][1] {
		return true
	}
	return false
}

func (p *Probe) Out(area *[][]int) bool {
	if p.posX > (*area)[0][1] || p.posY < (*area)[1][1] {
		return true
	}
	return false
}
