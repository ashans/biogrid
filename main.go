package main

import (
	"biogrid/internal/entities"
	"biogrid/internal/util/gutil"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"os"
	"strconv"
	"strings"
)

func main() {
	runApp()
}

func runApp() {
	err := os.Setenv("FYNE_SCALE", "0.8")
	if err != nil {
		panic(err)
	}
	a := app.New()
	mainWindow := a.NewWindow("BioGrid")
	mainWindow.Resize(fyne.NewSize(600, 400))
	seqValidator := func(s string) error {
		if len(s) == 0 {
			return errors.New("should not be empty")
		}
		if len(s) > 6 {
			return errors.New("should not exceed 6")
		}
		return nil
	}
	intValidator := func(s string) error {
		_, err := strconv.Atoi(s)
		if err != nil {
			return errors.New("should only be integer")
		}
		return nil
	}
	seq1 := widget.NewEntry()
	seq1.SetText("ATGCT")
	seq1.Validator = seqValidator
	seq2 := widget.NewEntry()
	seq2.SetText("AGCT")
	seq2.Validator = seqValidator
	matchVal := widget.NewEntry()
	matchVal.SetText("1")
	matchVal.Validator = intValidator
	mismatchVal := widget.NewEntry()
	mismatchVal.Validator = intValidator
	mismatchVal.SetText("-1")
	gapVal := widget.NewEntry()
	gapVal.SetText("-2")
	gapVal.Validator = intValidator
	mode := ""
	rGroup := widget.NewRadioGroup([]string{"Global", "Local"}, func(v string) {
		mode = v
	})
	rGroup.SetSelected("Global")
	inpForm := widget.NewForm(
		widget.NewFormItem("Sequence 1", seq1),
		widget.NewFormItem("Sequence 2", seq2),
		widget.NewFormItem("Match", matchVal),
		widget.NewFormItem("Mismatch", mismatchVal),
		widget.NewFormItem("Gap", gapVal),
		widget.NewFormItem("Alignment mode", rGroup),
	)

	p := layout.NewPaddedLayout()
	inpForm.OnSubmit = func() {
		fmt.Printf("Input | S1 : %s | S2 : %s | Mode : %s", seq1.Text, seq2.Text, mode)
		match, _ := strconv.Atoi(matchVal.Text)
		mismatch, _ := strconv.Atoi(mismatchVal.Text)
		gap, _ := strconv.Atoi(gapVal.Text)
		grid := container.NewGridWithColumns(len(seq1.Text)+2, getGridContent(entities.AlignConfig{
			Seq1: seq1.Text,
			Seq2: seq2.Text,
			Mode: mode,
			Scheme: entities.Scheme{
				Match:    match,
				Mismatch: mismatch,
				Gap:      gap,
			},
		})...)

		mainWindow.SetContent(container.NewVBox(
			container.New(p, inpForm),
			container.NewCenter(grid),
		))
	}
	mainWindow.SetContent(container.NewVBox(
		container.New(p, inpForm),
	))
	mainWindow.ShowAndRun()
}

func getGridContent(config entities.AlignConfig) (objs []fyne.CanvasObject) {

	grid := getValues(config)
	values := grid.Flatten()
	for _, v := range values {
		objs = append(objs, widget.NewLabel(v))
	}
	return objs
}

func getValues(config entities.AlignConfig) gutil.Grid {

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
