package main

import (
	"github.com/wcharczuk/go-chart"
	"os"
	"sort"
	"strconv"
)

// Item2 receive a map to create bar chart->  (label)Key: int, (value)value: float64
func Item2(values map[int]float64) {

	//Sort of the Numbers in the map
	keys := make([]int, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	//Adding the values to the Bar Values List
	var charValues []chart.Value
	for _, k := range keys {
		valV := values[k]
		valL := strconv.Itoa(k)
		charValues = append(charValues, chart.Value{Label: valL, Value: valV})
	}

	graph := chart.BarChart{
		Title: "Number Aparition Bar Chart",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		Bars:     charValues,
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}