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

## day ten

I find the \n is different in Windows and Linux.

It's easy to solve today's puzzle two. I'd say there's no other better way to learn a new language than this.

## day 11

It's hard to debug for recursive func. Finally.

## day 12

You just never read the puzzle clearly, right?

## day 13

OK, half of the puzzles have been done. Nice job boy.

## day 14

I guess it will be overtime again, before I start today's puzzle.

Oh, come on. 

I get it, just like day6. Getting the abstractly similar objects together as a count num is the key. 

## day 15

**Dijkstra**, right?

Again. How long will it take this time?

该函数执行完成耗时： 33m43.834441956s. lol, at least it works.

## day 16

My brain is burning.

Finally, I guess it's a **pipe** or so-called **queue**? The whole afternoon is wasted.

lol, I just rewrote it, and it simply worked.

Calm down. Maybe I need another day to figure it out. 
Why I passed all the unit tests but can't pass the real users? Damn!

fuuuuuuuuuuuuc, the bitSize! Why did it run successfully? There are no errors.

## day 17
> If you're going to fire a highly scientific probe out of a super cool probe launcher,
> 
> you might as well do it with style.

## day 18

I think it's tree. Hope I choose the right data struct.

Recursions are insane. I can't tell the bugs between them.

## day 20

Don't hesitate, just try it.

## day 21

> Now that you're warmed up, it's time to play the real game.

Nooooo, leave me alone.

## day 22

All right, it's such a hard work again. 

## day 23

Lmao, I finished the part one by my hand. Maybe the part two wil be solved by this, too. 

## day 24

Maybe I'd give it up. It's a nightmare.

https://github.com/dphilipson/advent-of-code-2021/blob/master/src/days/day24.rs .
What a clever move! I thought it could be simplified, but I didn't try to dig it out. 

## day 25

ok fine. Finally. I need 7 more stars. It's time to look back.