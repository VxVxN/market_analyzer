package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"

	"github.com/VxVxN/market_analyzer/internal/consts"
	"github.com/VxVxN/market_analyzer/internal/humanizer"
	hum "github.com/VxVxN/market_analyzer/internal/humanizer"
	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
	"github.com/VxVxN/market_analyzer/internal/parser/myfileparser"
	preparerpkg "github.com/VxVxN/market_analyzer/internal/preparer"
	e "github.com/VxVxN/market_analyzer/pkg/error"
	"github.com/VxVxN/market_analyzer/pkg/tools"
)

func (server *Server) commonDataHandler(c *gin.Context) {
	emitter := c.Param("name")

	report, err := prepareReport(emitter, marketanalyzer.NormalMode, hum.NumbersMode, 2)
	if err != nil {
		e.NewError("Failed to prepare report", http.StatusInternalServerError, nil).JsonResponse(c)
		return
	}

	commonValueForChart := []string{
		string(marketanalyzer.Sales),
		string(marketanalyzer.Earnings),
		string(marketanalyzer.Debts),
		string(marketanalyzer.MarketCap),
	}
	if err = renderChart(c.Writer, "Common data", emitter, report, commonValueForChart, false); err != nil {
		e.NewError("Failed to redner chart", http.StatusInternalServerError, nil).JsonResponse(c)
		return
	}
}

func prepareReport(emitter string, periodMode marketanalyzer.PeriodMode, numberMode hum.NumberMode, precisionFlag int) (*humanizer.ReadyData, error) {
	parser := myfileparser.Init("data/emitters/" + emitter + consts.CsvFileExtension)

	rawData, err := parser.Parse()
	if err != nil {
		return nil, fmt.Errorf("fail parse emitter: %v", err)
	}

	preparer := preparerpkg.Init(rawData)
	rawMarketData, err := preparer.Prepare()
	if err != nil {
		return nil, fmt.Errorf("fail prepare data: %v", err)
	}

	analyzer := marketanalyzer.Init(rawMarketData)
	analyzer.SetPeriodMode(periodMode)
	marketData, err := analyzer.Calculate()
	if err != nil {
		return nil, fmt.Errorf("fail market model: %v", err)
	}

	humanizer := hum.Init(marketData)
	humanizer.SetPrecision(precisionFlag)
	humanizer.SetNumbersMode(numberMode)
	humanizer.SetFieldsForDisplay(
		[]marketanalyzer.RowName{
			// marketanalyzer.Sales,
			// marketanalyzer.Earnings,
		},
	)
	return humanizer.Humanize(), nil
}

func renderChart(writer gin.ResponseWriter, pageTitle, emitter string, report *hum.ReadyData, ratiosForChart []string, isFloat bool) error {
	chart := charts.NewLine()
	chart.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:  emitter,
				Link:   "/emitter/" + emitter,
				Target: "self",
			},
		),
		charts.WithLegendOpts(
			opts.Legend{
				Show: true,
			},
		),
		charts.WithTooltipOpts(
			opts.Tooltip{
				Show:    true,
				Trigger: "axis",
				AxisPointer: &opts.AxisPointer{
					Type: "shadow",
					Snap: true,
				},
			},
		),
		charts.WithInitializationOpts(
			opts.Initialization{
				PageTitle: pageTitle,
				Width:     "100%",
			},
		),
	)
	chart.SetXAxis(report.Headers)

	for _, row := range report.Rows {
		if !tools.ContainStringInSlice(row[0], ratiosForChart) {
			continue
		}
		chart.AddSeries(row[0], prepareLineItems(row, isFloat))
	}

	chart.SetSeriesOptions(
		charts.WithLineChartOpts(
			opts.LineChart{
				ConnectNulls: true,
			},
		),
	)

	if err := chart.Render(writer); err != nil {
		return err
	}
	return nil
}

func prepareLineItems(row []string, isFloat bool) []opts.LineData {
	items := make([]opts.LineData, 0)
	rowWithoutTitle := row[1:]
	for _, rawValue := range rowWithoutTitle {
		var value string
		if rawValue == "-" {
			value = ""
		} else {
			if isFloat {
				value = rawValue
			} else {
				value = strings.Replace(rawValue, ".", "", -1)
			}
		}
		items = append(items, opts.LineData{Name: row[0], SymbolSize: 5, Value: value})
	}
	return items
}
