package algo

import (
	"biogrid/internal/entities"
	"biogrid/internal/util/gutil"
	"strconv"
	"strings"
)

func Align(config entities.AlignConfig) gutil.Grid {

	g := gutil.NewGrid(len(config.Seq1), len(config.Seq2), "")
	g = globalAlign(config, g)
	g = g.AddRow(append([]string{""}, strings.Split(config.Seq1, "")...), false)
	g = g.AddColumn(append([]string{"", ""}, strings.Split(config.Seq2, "")...), false)

	return g
}

func globalAlign(config entities.AlignConfig, g gutil.Grid) gutil.Grid {
	g = g.AddRow(buildGapSeries(len(config.Seq1), config.Gap), false)
	g = g.AddColumn(append([]string{"0"}, buildGapSeries(len(config.Seq2), config.Gap)...), false)
	dim := g.Dim()
	for i := 1; i < dim.Height; i++ {
		for j := 1; j < dim.Width; j++ {
			top := g.GetIntAt(j, i-1) + config.Gap
			left := g.GetIntAt(j-1, i) + config.Gap
			prev := g.GetIntAt(j-1, i-1)
			digAdd := matchValue(string(config.Seq1[j-1]), string(config.Seq2[i-1]), config)
			dig := prev + digAdd
			g.SetIntAt(j, i, max(top, left, dig))
		}
	}

	return g
}

func buildGapSeries(n, v int) (s []string) {
	for i := 0; i < n; i++ {
		s = append(s, strconv.Itoa(v*(i+1)))
	}
	return s
}

func matchValue(a string, b string, config entities.AlignConfig) int {
	if a == b {
		return config.Match
	}
	return config.Mismatch
}

func max(a ...int) (r int) {
	r = a[0]
	for _, v := range a {
		if r < v {
			r = v
		}
	}
	return r
}
