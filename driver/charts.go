package driver

import (
	"sort"

	"github.com/vicanso/go-charts/v2"
)

func chartDefaults(opt *charts.ChartOption) {
	opt.Theme = "dark"
	opt.Height = 400
	opt.BackgroundColor = charts.Color{
		R: 24,
		G: 28,
		B: 37,
		A: 255,
	}
}

func ChartLine(values []float64, labels []string, unit string) ([]byte, error) {
	p, err := charts.LineRender(
		[][]float64{values},
		charts.SVGTypeOption(),
		func(opt *charts.ChartOption) {
			chartDefaults(opt)
			opt.XAxis = charts.NewXAxisOption(labels)
			opt.SymbolShow = charts.FalseFlag()
			opt.Legend = charts.LegendOption{
				Data: []string{unit},
			}
			opt.LineStrokeWidth = 2
		},
	)
	if err != nil {
		return nil, err
	}
	return p.Bytes()
}

func ChartPie(values []string) ([]byte, error) {
	m := make(map[string]float64)
	for _, v := range values {
		c, _ := m[v]
		m[v] = c + 1
	}
	counts := make([]float64, len(m))
	labels := make([]string, len(m))
	i := 0
	for k := range m {
		labels[i] = k
		i += 1
	}
	sort.Strings(labels)
	for i, label := range labels {
		counts[i] = m[label]
	}
	p, err := charts.PieRender(
		counts,
		charts.SVGTypeOption(),
		charts.PieSeriesShowLabel(),
		func(opt *charts.ChartOption) {
			chartDefaults(opt)
			f := false
			opt.Legend = charts.LegendOption{
				Orient: charts.OrientVertical,
				Data:   labels,
				Show:   &f,
			}
		},
	)
	if err != nil {
		return nil, err
	}
	return p.Bytes()
}
