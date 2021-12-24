package Day23

func day23PartTwo(input string) int {

	roomLeft := [2]int{}
	room1 := NewStack(2)
	room2 := NewStack(2)
	room3 := NewStack(2)
	room4 := NewStack(2)
	roomRight := [2]int{}
	roomShed := [3]int{}

	room1.Push(2)
	room1.Push(1)
	room2.Push(3)
	room2.Push(4)
	room3.Push(2)
	room3.Push(3)
	room4.Push(4)
	room4.Push(1)

	NextStep(room1, room2, room3, room4, roomLeft, roomRight, roomShed, 0)

	return 0
}

func NextStep(room1 Stack, room2 Stack, room3 Stack, room4 Stack,
	roomLeft [2]int, roomRight [2]int, roomShed [3]int, energy int) {
	if roomLeft[1] != 0 {
		switch roomLeft[1] {
		case 1:
			if room1.IsEmpty() {
				room1.Push(1)
				energy = energy + 10
			}
			break
		}
	}

	if !room1.IsEmpty() {
		//num := room1.Pop()
		if roomLeft[1] == 0 {

		}
	}

}

func findHome(num int, room1 *Stack, room2 *Stack, room3 *Stack, room4 *Stack) {

}
