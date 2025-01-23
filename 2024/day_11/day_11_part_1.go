package day11

import (
	"fmt"
	"strconv"
)

/*
Plutonian Pebbles, every time i blink the number on the the stones change.
Other times, a stone might well split in two, causing all other stones to shift over a bit to make
room in their perfectly straight line.

After some observation, you notice that the stones each simultaneously change according to the first applicable rule within the list:

no matter how the stones change, their order is preserved, and they can stay on their perfectly straight line
*/
func SolveDay11Part1() {
	var pp Plutonian
	// puzzleInput := []int64{0, 1, 10, 99, 999}
	puzzleInput := []int64{70949, 6183, 4, 3825336, 613971, 0, 15, 182}
	pp = &Pebble{Line: puzzleInput}

	for i := 0; i < 75; i++ {
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
	for i := 0; i < len(p.Line); i++ {
		pebble := p.Line[i]
		if pebble == 0 {
			p.Line[i] = 1
		} else if isEven, stringRep := checkEvenDigits(pebble); isEven {
			n := len(stringRep)
			leftStone, _ := strconv.ParseInt(stringRep[:n/2], 10, 64)
			rightStone, _ := strconv.ParseInt(stringRep[n/2:], 10, 64)
			p.createSpaceAndInsertStones(leftStone, rightStone, i)
			i += 1 // Skip the subsequent iteration, to count for in-place insertion
		} else {
			p.Line[i] = pebble * 2024
		}
	}
}

func (p *Pebble) stones() int {
	return len(p.Line)
}

/*
I'm tempted to shift everything off to the right, but then the outer loop is going to be messed up
unless I can find some clever way of skipping an interation
*/
func (p *Pebble) createSpaceAndInsertStones(leftStone, rightstone int64, i int) {
	p.Line = append(p.Line[:i+1], p.Line[i:]...) // use the unpack operator to add in the rest of the array "effectively duplicaing the ith position
	p.Line[i] = leftStone
	p.Line[i+1] = rightstone
}

func checkEvenDigits(digits int64) (bool, string) {
	stringRep := strconv.Itoa(int(digits))
	return len(stringRep)%2 == 0, stringRep
}

/*
Define the contract for a mysterious plutonian entity
*/
type Plutonian interface {
	blink()      // Blink and you'll miss it
	stones() int // How many magic stones are there
}
