package gutil

import (
	"errors"
	"fmt"
	"strconv"
)

type Grid [][]string

type GridDimension struct {
	Width  int
	Height int
}

func NewGrid(w int, h int, v string) Grid {
	var g [][]string
	for i := 0; i < h; i++ {
		g = append(g, initSlice(w, v))
	}

	return g
}

func (g Grid) ToString() string {
	return fmt.Sprintf("%v", g)
}

func (g *Grid) GetAt(x, y int) string {
	return [][]string(*g)[y][x]
}

func (g *Grid) GetIntAt(x, y int) int {
	val := g.GetAt(x, y)
	num, _ := strconv.Atoi(val)

	return num
}

func (g *Grid) SetAt(x, y int, v string) {
	[][]string(*g)[y][x] = v
}

func (g *Grid) SetIntAt(x, y int, v int) {
	[][]string(*g)[y][x] = strconv.Itoa(v)
}

func (g *Grid) Dim() GridDimension {
	return GridDimension{
		len([][]string(*g)[0]),
		len([][]string(*g)),
	}
}

func (g *Grid) AddRow(row []string, toEnd bool) Grid {

	if len(row) != len([][]string(*g)[0]) {
		panic(errors.New("invalid length in row"))
	}
	cop := row[:]

	if toEnd {
		return append(*g, cop)
	}
	return append([][]string{cop}, *g...)
}
func (g *Grid) AddRowWithValue(v string, toEnd bool) Grid {

	return g.AddRow(initSlice(len([][]string(*g)[0]), v), toEnd)
}

func (g *Grid) AddColumn(col []string, toEnd bool) Grid {

	if len(col) != len([][]string(*g)) {
		panic(errors.New("invalid length in column"))
	}
	var gNew Grid
	for i, v := range col {
		if toEnd {
			gNew = append(gNew, append([][]string(*g)[i], v))
		} else {
			gNew = append(gNew, append([]string{v}, [][]string(*g)[i]...))
		}
	}

	return gNew
}

func (g *Grid) AddColumnWithValue(v string, toEnd bool) Grid {

	return g.AddColumn(initSlice(len([][]string(*g)), v), toEnd)
}

func (g *Grid) Flatten() []string {
	var f []string
	for _, v := range *g {
		f = append(f, v...)
	}

	return f
}

func initSlice(n int, v string) []string {
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, v)
	}
	return s
}
