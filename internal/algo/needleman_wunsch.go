package algo

import (
	"biogrid/internal/entities"
	"biogrid/internal/util/gutil"
	"strconv"
)

type NeedlemanWunsch struct{}

func (NeedlemanWunsch) InitGrid(g gutil.Grid, config entities.AlignConfig) gutil.Grid {
	g = g.AddRow(incrementalGapSeries(len(config.Seq1), config.Gap), false)
	g = g.AddColumn(append([]string{"0"}, incrementalGapSeries(len(config.Seq2), config.Gap)...), false)

	return g
}

func (NeedlemanWunsch) CellDetermine(top, left, diagonal int, s1, s2 string, config entities.AlignConfig) int {
	topV := top + config.Gap
	leftV := left + config.Gap
	diagonalV := diagonal + matchValue(s1, s2, config)
	return max(topV, leftV, diagonalV)
}

func incrementalGapSeries(n, v int) (s []string) {
	for i := 0; i < n; i++ {
		s = append(s, strconv.Itoa(v*(i+1)))
	}
	return s
}
