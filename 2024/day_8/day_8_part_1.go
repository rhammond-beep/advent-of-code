package day_8

import (
	"fmt"
	"unicode"
)

type Point struct {
	I int
	J int
}

type AntenaPair struct {
	Type      rune
	AntenaOne Point
	AntenaTwo Point
}

func (a *AntenaPair) PrintString(antiNodesMapped int) {
	fmt.Printf("----\n type: %v\n A1: %v\n A2: %v\n antinodes_mapped: %v\n----", a.Type, a.AntenaOne, a.AntenaTwo, antiNodesMapped)
}

/*
We know that for every pair of antennas, there can be at most two corresponding antinodes
This feels like we can draw a line which intersects with both the points, then plot two new points on either side of the far
*/
func (a *AntenaPair) computeAntinodePoints() (Point, Point) {
	dist_i, dist_j := a.AntenaOne.CalculateDistance(&a.AntenaTwo)
	p1 := Point{I: a.AntenaOne.I - (dist_i * 2), J: a.AntenaOne.J - (dist_j * 2)}
	p2 := Point{I: a.AntenaOne.I + (dist_i * 1), J: a.AntenaOne.J + (dist_j * 1)}
	return p1, p2
}

func (p *Point) CalculateDistance(p2 *Point) (int, int) {
	diffI := float64(p.I) - float64(p2.I)
	diffJ := float64(p.J) - float64(p2.J)
	return int(diffI), int(diffJ)
}

/*
the most sensible solutions seems to follow these steps:
 1. Find the set of "Perfectly Inline" candidate antinode positions
 2. These antinode positions should be based on every possible pair of identical antenas
 3. each position to see if the constraints are satisfied (as perscribed by the "Calculate Distance" Method
*/
func SolveDay8Part1(cityMap []string) int {
	antinodes := 0

	antenaPairs := ComputeAntenaPairs(cityMap)
	antinodeMap := make(map[Point]rune)
	for _, pair := range antenaPairs {
		antinodesMapped := 0
		an1, an2 := pair.computeAntinodePoints()

		if an1.I < len(cityMap) && an1.J < len(cityMap) && an1.I > -1 && an2.J > -1 {
			if _, present := antinodeMap[an1]; !present {
				antinodeMap[an1] = '#'
				antinodesMapped += 1
				antinodes += 1
			}
		}

		if an2.I < len(cityMap) && an2.J < len(cityMap) && an2.I > -1 && an2.J > -1 {
			if _, present := antinodeMap[an2]; !present {
				antinodeMap[an2] = '#'
				antinodesMapped += 1
				antinodes += 1
			}
		}

		pair.PrintString(antinodesMapped)
	}

	return antinodes
}

func ComputeAntenaPairs(cityMap []string) []AntenaPair {
	antenaTypeMap := make(map[Point]rune)
	antenaLocationMap := make(map[rune][]Point)

	for i := 0; i < len(cityMap); i++ {
		for j := 0; j < len(cityMap[0]); j++ {
			if unicode.IsLetter(rune(cityMap[i][j])) || unicode.IsDigit(rune(cityMap[i][j])) {
				antena := rune(cityMap[i][j])
				antenaLocation := Point{I: i, J: j}

				antenaPoints := antenaLocationMap[antena]
				antenaPoints = append(antenaPoints, antenaLocation)
				antenaLocationMap[antena] = antenaPoints

				antenaTypeMap[antenaLocation] = antena
			}
		}
	}

	// Compute all pairs of antena, order doesn't matter!
	pairs := make([]AntenaPair, 0)
	for key, points := range antenaLocationMap {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				pairs = append(pairs, AntenaPair{Type: key, AntenaOne: points[i], AntenaTwo: points[j]})
			}

		}
	}

	return pairs
}
