package algo

import (
	"biogrid/internal/entities"
	"biogrid/internal/util/gutil"
)

type AlignAlgorithm interface {
	InitGrid(gutil.Grid, entities.AlignConfig) gutil.Grid
	CellDetermine(top, left, diagonal int, s1, s2 string, config entities.AlignConfig) int
}
