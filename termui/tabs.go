package main

import(
"github.com/gizak/termui"
"github.com/gizak/termui/extra"
)

func main() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()
	
	tab1 := extra.NewTab("  All  ")
	tab2 := extra.NewTab("  Pods  ")

	bc := termui.NewBarChart()
	data := []int{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}
	bclabels := []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	bc.BorderLabel = "Bar Chart"
	bc.Data = data
	bc.Height = 10
	bc.DataLabels = bclabels
	bc.TextColor = termui.ColorGreen
	bc.BarColor = termui.ColorRed
	bc.NumColor = termui.ColorYellow

	grid := termui.NewGrid()
	grid.X, grid.Y = 0, 0
	grid.BgColor = termui.ThemeAttr("bg")
	grid.Width = termui.TermWidth()
	grid.AddRows(
		termui.NewRow(
			termui.NewCol(12, 0, bc),
		),
	)
	tab2.AddBlocks(grid)

	tabpane := extra.NewTabpane()
	tabpane.Y = 1
	tabpane.Width = termui.TermWidth()
	tabpane.Border = true

	tabpane.SetTabs(*tab1, *tab2)	

	termui.Render(tabpane)	

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/1", func(termui.Event) {
		tabpane.SetActiveLeft()
		termui.Clear()
		termui.Render(tabpane)
	})

	termui.Handle("/sys/kbd/2", func(termui.Event) {
		tabpane.SetActiveRight()
		grid.Align()
		termui.Clear()
		termui.Render(tabpane)
	})

	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		grid.Width = termui.TermWidth()
		grid.Align()
		termui.Clear()
		termui.Render(tabpane)
	})

	termui.Loop()	
}