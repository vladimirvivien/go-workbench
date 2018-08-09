package main

import (
	"fmt"
	"math"

	ui "github.com/gizak/termui"
)

func main() {
	// initialize termui
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	sinps := func() []float64 {
		n := 400
		ps := make([]float64, n)
		for i := range ps {
			ps[i] = 1 + math.Sin(float64(i)/5)
		}
		return ps
	}()

	lc := ui.NewLineChart()
	lc.BorderLabel = "braille-mode Line Chart"
	lc.Data = sinps
	lc.Height = 20
	lc.AxesColor = ui.ColorWhite
	lc.LineColor = ui.ColorYellow | ui.AttrBold

	sinpsint := (func() []int {
		ps := make([]int, len(sinps))
		for i, v := range sinps {
			ps[i] = int(100*v + 10)
		}
		return ps
	})()

	spark := ui.Sparkline{}
	spark.Height = 8
	spdata := sinpsint
	spark.Data = spdata[:100]
	spark.LineColor = ui.ColorCyan
	spark.TitleColor = ui.ColorWhite

	sp := ui.NewSparklines(spark)
	sp.Height = 11
	sp.BorderLabel = "Sparkline"

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, lc),
		),
	)

	ui.Body.Align()
	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/n", func(ui.Event) {
		ui.Clear()
		fmt.Println("n pressed")
		// ui.Body.AddRows(
		// 	ui.NewRow(
		// 		ui.NewCol(12, 0, sp),
		// 	),
		// )
		// ui.Body.Width = ui.TermWidth()
		// ui.Body.Align()
		// ui.Render(ui.Body)
	})

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)
		i := t.Count
		if i > 10 {
			ui.StopLoop()
			return
		}

		lc.Data = sinps[2*i:]
		ui.Render(ui.Body)
	})

	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})

	ui.Handle("/sys/kbd/c", func(ui.Event) {
		ui.Clear()

		sinps := (func() []float64 {
			n := 220
			ps := make([]float64, n)
			for i := range ps {
				ps[i] = 1 + math.Sin(float64(i)/5)
			}
			return ps
		})()

		lc0 := ui.NewLineChart()
		lc0.BorderLabel = "braille-mode Line Chart"
		lc0.Data = sinps
		lc0.Width = 50
		lc0.Height = 12
		lc0.X = 0
		lc0.Y = 0
		lc0.AxesColor = ui.ColorWhite
		lc0.LineColor = ui.ColorGreen | ui.AttrBold

		ui.Render(lc0)
	})

	ui.Loop()
}
