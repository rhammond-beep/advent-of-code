package day_8

import (
	"fmt"
	"math"
	"unicode"
	// helper "github.com/rhammond-beep/advent-of-code-go-helper"
)

type Point struct {
	I int
	J int
}

/*
Calculate the difference between two points
*/
func (p *Point) CalculateDistance(p2 *Point) int {
	diff := math.Abs(float64(p.I)-float64(p2.I)) + math.Abs(float64(p.J)-float64(p2.J))
	return int(math.Floor(diff))
}

/*
takes in two points and checks to see if they're double the diagonal distance
*/
func isCandidatePointDoubleTheDistance(candidateLocation, point1, point2 *Point) bool {
	return (candidateLocation.CalculateDistance(point1) * 2) == candidateLocation.CalculateDistance(point2)
}

func SolveDay8Part1() {
	// cityMap := helper.ReadChallengeInput("./day_8_input.txt")
	cityMap := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	} // Expected Answer Is 14

	fmt.Println(calculateAntinodes(cityMap))
}

/*
Each antenna is tuned to a specific frequency, is indicated
as a single lowercase, uppercase letter, or digit.

an antinode occurs at any point that is perfectly in line with two antennas of the same frequency - but only when one of the antennas is
twice as far away as the other.

This means that for any pair of antennas with the the sane
*/
func calculateAntinodes(cityMap []string) (antinodes int) {

	antenaMap := make(map[Point]rune)

	for i := 0; i < len(cityMap); i++ { // assume square array
		for j := 0; j < len(cityMap); j++ {
			if unicode.IsLetter(rune(cityMap[i][j])) {
				antenaMap[Point{I: i, J: j}] = rune(cityMap[i][j])
			}
		}
	}

	antinode := &Point{I: 1, J: 1}
	point1 := &Point{I: 2, J: 2}
	point2 := &Point{I: 3, J: 3}

	fmt.Printf("%v\n", isCandidatePointDoubleTheDistance(antinode, point1, point2))

	// check to see if each valid locatio

	return
}
