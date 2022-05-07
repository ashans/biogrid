package form

import (
	"biogrid/internal/entities"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func newSequenceField(s string) *widget.Entry {
	seq := widget.NewEntry()
	seq.SetText(s)
	seq.Validator = seqValidator
	return seq
}

func newSchemeField(v int) *widget.Entry {
	field := widget.NewEntry()
	field.SetText(strconv.Itoa(v))
	field.Validator = intValidator
	return field
}

func newModeSelector() *widget.RadioGroup {
	rGroup := widget.NewRadioGroup([]string{entities.GlobalAlignment, entities.LocalAlignment}, func(v string) {})
	rGroup.SetSelected(entities.GlobalAlignment)

	return rGroup
}
