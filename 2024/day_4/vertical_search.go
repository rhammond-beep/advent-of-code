package day4

type VerticalSearch struct {
}

func (vs *VerticalSearch) FindTermOccurrences(ws *WordSearch) (occurrences int) {
	var mappedSearchSpace []string

	// Map to a horizontal space
	for i := 0; i < len(ws.SearchSpace); i++ {
		var sb strings.Builder
		for j := 0; j < len(ws.SearchSpace); j++ {
			sb.WriteByte(ws.SearchSpace[j][i])
		}
		mappedSearchSpace = append(mappedSearchSpace, sb.String())
	}

	hs := &HorizontalSearch{}
	mappedSearch := &WordSearch{SearchSpace: mappedSearchSpace, SearchTerm: ws.SearchTerm}

	return hs.FindTermOccurrences(mappedSearch)
}
