package day4

type DiagonalSearch struct {
}

/*
Diagonal Search is a little more tricky that the others,
We want to move across the diagonal lines, however we have to exclude all
whose length (number of dicrete points n) < len(ws.SearchTerm)

This seems like we've maybe got to do a little pre-processing. One thing we can
do is build up a slice containing the indexes of the diagonals [][]int

[

	[3,0], (i, j)
	[2,1],
	[1,2],
	[0,3]

]

We take the size of the diagonals len(diagonals) and remove anything whose size is
< len(ws.SearchTerm)

Now Iterate over the remaining lines with the sliding window, use this to extract out
a string compare it to the desired searchTerm.

Feels like I shouldn't have to do this preprocessing stage though... Feels pretty annoying,
I'd prefer to just do the work as part of the original loop.
*/
func (ds *DiagonalSearch) FindTermOccurrences(ws *WordSearch) (occurrences int) {
	var mappedSearchSpace []string

	for k := 0; k < len(ws.SearchSpace)-1; k++ {
		var sb strings.Builder
		for j := 0; j <= k; j++ {
			i := k - j
			sb.WriteByte(ws.SearchSpace[i][j])
		}
		mappedSearchSpace = append(mappedSearchSpace, sb.String())
	}

	for k := len(ws.SearchSpace) - 1; k > 0; k-- {
		var sb strings.Builder
		for j := 0; j <= k; j++ {
			i := k - j
			sb.WriteByte(ws.SearchSpace[i][j])
		}
		mappedSearchSpace = append(mappedSearchSpace, sb.String())
	}

	for k := 0; k < len(ws.SearchSpace)-1; k++ {
		var sb strings.Builder
		for j := len(ws.SearchSpace) - 1; j >= k; j-- {
			i := j - k
			sb.WriteByte(ws.SearchSpace[i][j])
		}
		mappedSearchSpace = append(mappedSearchSpace, sb.String())
	}

	for k := len(ws.SearchSpace) - 1; k > 0; k-- {
		var sb strings.Builder
		for j := k; j > 0; j-- {
			i := k - j
			sb.WriteByte(ws.SearchSpace[i][j])
		}
		mappedSearchSpace = append(mappedSearchSpace, sb.String())
	}

	mappedWordSearch := &WordSearch{SearchSpace: mappedSearchSpace, SearchTerm: ws.SearchTerm}
	hs := HorizontalSearch{}
	return hs.FindTermOccurrences(mappedWordSearch)
}
