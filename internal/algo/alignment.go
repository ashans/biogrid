package algo

import (
	"biogrid/internal/entities"
	"biogrid/internal/util/gutil"
	"errors"
	"strings"
)

type Aligner struct {
	global AlignAlgorithm
	local  AlignAlgorithm
}

func NewAligner() *Aligner {
	return &Aligner{
		global: NeedlemanWunsch{},
		local:  SmithWaterman{},
	}
}

func (a *Aligner) Align(config entities.AlignConfig) gutil.Grid {

	g := gutil.NewGrid(len(config.Seq1), len(config.Seq2), "")

	algorithm := a.selectAlgorithm(config)
	g = align(config, g, algorithm)
	g = g.AddRow(append([]string{""}, strings.Split(config.Seq1, "")...), false)
	g = g.AddColumn(append([]string{"", ""}, strings.Split(config.Seq2, "")...), false)

	return g
}

func (a *Aligner) selectAlgorithm(config entities.AlignConfig) AlignAlgorithm {
	switch config.Mode {
	case entities.GlobalAlignment:
		return a.global
	case entities.LocalAlignment:
		return a.local
	default:
		panic(errors.New("invalid mode to select algorithm"))
	}
}

func align(config entities.AlignConfig, g gutil.Grid, algo AlignAlgorithm) gutil.Grid {
	g = algo.InitGrid(g, config)
	dim := g.Dim()
	for i := 1; i < dim.Height; i++ {
		for j := 1; j < dim.Width; j++ {
			g.SetIntAt(j, i, algo.CellDetermine(g.GetIntAt(j, i-1), g.GetIntAt(j-1, i), g.GetIntAt(j-1, i-1), string(config.Seq1[j-1]), string(config.Seq2[i-1]), config))
		}
	}

	return g
}
