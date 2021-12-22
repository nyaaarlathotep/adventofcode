package Day21

func Day21(input string) int {

	playerA := newPlayer(3)
	playerB := newPlayer(10)
	counter := 0
	counterP := &counter
	rollCounter := 0
	for true {
		playerA.movePawn(counterP)
		rollCounter = rollCounter + 3
		if playerA.win() {
			break
		}
		playerB.movePawn(counterP)
		rollCounter = rollCounter + 3

		if playerB.win() {
			break
		}
	}
	res := 0
	if playerA.win() {
		res = playerB.score
	} else {
		res = playerA.score
	}
	return res * rollCounter
}

func PartTwo(input string) int {

	posA := 3
	posB := 10
	isA := true
	roll := 0
	cache = make(map[state]winState)
	var winA, winB int
	for i := 1; i <= 3; i++ {
		s := state{isA, roll, i, 0, 0, posA , posB}
		win := nextTurn(s)
		winA += win.AWinCounter
		winB += win.BWinCounter
	}
	if winA > winB {
		return winA
	} else {
		return winB
	}

}

var cache map[state]winState

func nextTurn(s state) winState {
	if res, found := cache[s]; found {
		return res
	}

	sStart := s
	if s.isA {
		s.posA = 1 + ((s.posA + s.dice - 1) % 10)
		if s.roll == 2 {
			s.scoreA += s.posA
		}
		if s.scoreA >= 21 {
			cache[sStart] = winState{1, 0}
			return winState{1, 0}
		}
	} else {
		s.posB = 1 + ((s.posB + s.dice - 1) % 10)
		if s.roll == 2 {
			s.scoreB += s.posB
		}
		if s.scoreB >= 21 {
			cache[sStart] = winState{0, 1}
			return winState{0, 1}
		}
	}

	if s.roll == 2 {
		s.isA = !s.isA
	}
	s.roll = (s.roll + 1) % 3
	var winA, winB int
	for i := 1; i <= 3; i++ {
		s.dice = i
		win := nextTurn(s)
		winA += win.AWinCounter
		winB += win.BWinCounter
	}

	cache[sStart] = winState{winA, winB}
	return winState{winA, winB}

}

func nextDiceNum(counter *int) int {
	*counter = *counter + 1
	if *counter == 101 {
		*counter = 1
	}
	return *counter
}

type state struct {
	isA    bool
	roll   int
	dice   int
	scoreA int
	scoreB int
	posA   int
	posB   int
}

type winState struct {
	AWinCounter int
	BWinCounter int
}
