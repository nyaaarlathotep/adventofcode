# advent of code

advent of code 2021 of nyaaar

May I try to learn some GO by this chance. Hope I can hold on for a little longer.

![Alt](https://repobeats.axiom.co/api/embed/9774c0942d69815f520f3eea28a9af134d5ed00b.svg "Repobeats analytics image")

## day four

It's getting harder, I'm stuck in my foolishness.

## day five

Processing 45° and 135° is different. I forget it, what a shame.

Besides, I try to guess the part two's question, but I failed.

## day six

Damn, I found why it's difficult. My heap! You bastard. Maybe I'd try some recursion.

Oh, it cost too much time now. Maybe I'd try another way.

该函数执行完成耗时： 25m28.9910549s. The result is right though.

Oh, thanks a lot reddit. I'm such a fool. 该函数执行完成耗时： 21.378µs It goes well this time.

## day seven

I mistake today's puzzle, god! I thought there may be numbers don't show up which is much more confusing.

该函数执行完成耗时： 2.7331ms

here's my hint:
```go
// return whether string big contains string small. 
// It allows a number of misCount letters missed in the big string. 
func strContain(small string, big string, misCount int) bool {
	sSlice := strings.Split(small, "")
	bSlice := strings.Split(big, "")
	for _, s := range sSlice {
		get := false
		for _, ss := range bSlice {
			if s == ss {
				get = true
			}
		}
		if !get {
			misCount--
		}
		if misCount < 0 {
			return false
		}
	}
	return true
}
```

## day eight

It's a little difficult to understand the puzzle today. What does a basin really mean? It can be more accurate.

It's a pleasure to see how to build a data structure of set in GO. How ingenious!