package day4

type HorizontalSearch struct {
}

/*
check each row for any occuring instances of the search term
*/
func (hs *HorizontalSearch) FindTermOccurrences(ws *WordSearch) (occurrences int) {
	ub := len(ws.SearchTerm)
	for _, row := range ws.SearchSpace {
		// For each row, split up into chunks based on the length of the target word
		for i := ub; i < len(row)+1; i += 1 {
			word := row[i-ub : i]
			if word == ws.SearchTerm {
				occurrences += 1
			}
		}
	}
	return
}
