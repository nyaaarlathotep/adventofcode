package main

import "test/DataStruct"

func day23PartTwo(input string) int {

	room1 := DataStruct.NewStack()
	room2 := DataStruct.NewStack()
	room3 := DataStruct.NewStack()
	room4 := DataStruct.NewStack()
	room1.Push("B")
	room1.Push("A")
	room2.Push("C")
	room2.Push("D")
	room3.Push("B")
	room3.Push("C")
	room4.Push("D")
	room4.Push("A")

	rooms := make([]DataStruct.Stack, 0)
	rooms = append(rooms, room1, room4, room2, room3)



	return 0
}
