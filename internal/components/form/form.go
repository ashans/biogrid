package form

import (
	"biogrid/internal/entities"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func NewForm(config entities.AlignConfig, onSubmit func(entities.AlignConfig)) *widget.Form {
	seq1 := newSequenceField(config.Seq1)
	seq2 := newSequenceField(config.Seq2)
	matchVal := newSchemeField(config.Match)
	mismatchVal := newSchemeField(config.Mismatch)
	gapVal := newSchemeField(config.Gap)
	modeSelector := newModeSelector()
	inpForm := widget.NewForm(
		widget.NewFormItem("Sequence 1", seq1),
		widget.NewFormItem("Sequence 2", seq2),
		widget.NewFormItem("Match", matchVal),
		widget.NewFormItem("Mismatch", mismatchVal),
		widget.NewFormItem("Gap", gapVal),
		widget.NewFormItem("Alignment mode", modeSelector),
	)

	inpForm.OnSubmit = func() {
		match, _ := strconv.Atoi(matchVal.Text)
		mismatch, _ := strconv.Atoi(mismatchVal.Text)
		gap, _ := strconv.Atoi(gapVal.Text)
		onSubmit(entities.AlignConfig{
			Seq1: seq1.Text,
			Seq2: seq2.Text,
			Mode: modeSelector.Selected,
			Scheme: entities.Scheme{
				Match:    match,
				Mismatch: mismatch,
				Gap:      gap,
			},
		})
	}

	return inpForm
}
