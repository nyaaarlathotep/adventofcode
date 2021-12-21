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

func nextDiceNum(counter *int) int {
	*counter = *counter + 1
	if *counter == 101 {
		*counter = 1
	}
	return *counter
}
