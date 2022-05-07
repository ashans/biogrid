package algo

import (
	"biogrid/internal/entities"
	"biogrid/internal/util/gutil"
)

type SmithWaterman struct{}

func (SmithWaterman) InitGrid(g gutil.Grid, config entities.AlignConfig) gutil.Grid {
	g = g.AddRow(zeroSeries(len(config.Seq1)), false)
	g = g.AddColumn(append([]string{"0"}, zeroSeries(len(config.Seq2))...), false)

	return g
}

func (SmithWaterman) CellDetermine(top, left, diagonal int, s1, s2 string, config entities.AlignConfig) int {
	topV := minZero(top + config.Gap)
	leftV := minZero(left + config.Gap)
	diagonalV := minZero(diagonal + matchValue(s1, s2, config))
	return max(topV, leftV, diagonalV)
}

func zeroSeries(n int) (s []string) {
	for i := 0; i < n; i++ {
		s = append(s, "0")
	}
	return s
}

func minZero(n int) int {
	if n < 0 {
		return 0
	}
	return n
}
