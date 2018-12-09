package main

import (
	"fmt"
)

func main() {
	Roll(410, 72059)
	Roll(410, 7205900)
}

type Marble struct {
	value      int
	prev, next *Marble
}

type List struct {
	current *Marble
	head    *Marble
	len     int
}

func (l *List) Insert(newNode *Marble) {
	curNext := l.current.next

	newNode.prev = curNext
	newNode.next = curNext.next
	curNext.next = newNode
	newNode.next.prev = newNode

	l.current = newNode
	l.len++
}

func Roll(nPlayers, numMarbles int) {
	circle := &List{}
	currentMarble := &Marble{value: 0}
	currentMarble.prev = currentMarble
	currentMarble.next = currentMarble

	var currentPlayer int
	scores := map[int]int{}

	circle.current = currentMarble
	circle.head = currentMarble
	circle.tail = currentMarble

	for marbleToInsert := 1; marbleToInsert <= numMarbles; marbleToInsert++ {
		if marbleToInsert%23 == 0 {
			currentPlayer = ((currentPlayer + 1) % nPlayers)

			toRemove := circle.current
			for i := 0; i < 7; i++ {
				toRemove = toRemove.prev
			}
			toRemove.prev.next = toRemove.next
			toRemove.next.prev = toRemove.prev
			circle.current = toRemove.next
			scores[currentPlayer] += (marbleToInsert + toRemove.value)
		} else {
			newNode := &Marble{value: marbleToInsert}
			circle.Insert(newNode)
		}
	}
	var max int
	for _, s := range scores {
		if s > max {
			max = s
		}
	}
	fmt.Printf("max = %+v\n", max)
}
