package main

import (
	"image"
	"image/color"
	"math"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type SourceLine struct {
	TopLeft    image.Point
	Width      int
	Text       string
	TextHeight unit.Sp
	Bold       bool
	Color      color.NRGBA
}

func (line SourceLine) Layout(th *material.Theme, gtx layout.Context) {
	gtx.Constraints.Min.X = 0
	gtx.Constraints.Max.X = math.MaxInt
	gtx.Constraints.Min.Y = 0
	gtx.Constraints.Max.Y = math.MaxInt

	defer op.Offset(line.TopLeft).Push(gtx.Ops).Pop()
	if line.Width > 0 {
		defer clip.Rect{Max: image.Pt(line.Width, gtx.Metric.Sp(line.TextHeight))}.Push(gtx.Ops).Pop()
	}

	font := text.Font{Variant: "Mono"}
	if line.Bold {
		font.Weight = text.Heavy
	}
	paint.ColorOp{Color: line.Color}.Add(gtx.Ops)
	widget.Label{MaxLines: 1}.Layout(gtx, th.Shaper, font, line.TextHeight, line.Text)
}