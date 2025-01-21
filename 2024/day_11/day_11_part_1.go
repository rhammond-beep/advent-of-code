package day11

import "fmt"

/*
Plutonian Pebbles, every time i blink the number on the the stones change.
Other times, a stone might well split in two, causing all other stones to shift over a bit to make
room in their perfectly straight line.

After some observation, you notice that the stones each simultaneously change according to the first applicable rule within the list:

no matter how the stones change, their order is preserved, and they can stay on their perfectly straight line
*/
func SolveDay11Part1() {
	var pp Plutonian
	puzzleInput := []int64{0, 1, 10, 99, 999}
	pp = &Pebble{Line: puzzleInput}

	for i := 0; i < 1; i++ {
		pp.blink()
	}

	fmt.Println(pp.stones())
}

type Pebble struct {
	Line []int64
}

/*
- If the stone is engraved with the number 0, it gets replaced by a 1
- If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
- If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
*/
func (p *Pebble) blink() {
	for i, pebble := range p.Line {
		if pebble == 0 {
			p.Line[i] = 1
		}
		if true { // even number of digits

		} else {
			p.Line[i] = pebble * 2024
		}
	}
}

func (p *Pebble) stones() int {
	return len(p.Line)
}

/*
Define the contract for a mysterious plutonian entity
*/
type Plutonian interface {
	blink()
	stones() int
}
