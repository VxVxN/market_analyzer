package commands

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/spf13/cobra"

	hum "github.com/VxVxN/market_analyzer/internal/humanizer"
	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
	"github.com/VxVxN/market_analyzer/pkg/tools"
)

func InitWebCmd() *cobra.Command {
	var precisionFlag int
	var periodFlag string

	commonValueForChart := []string{
		string(marketanalyzer.Sales),
		string(marketanalyzer.Earnings),
		string(marketanalyzer.Debts),
		string(marketanalyzer.MarketCap),
	}

	ratiosForChart := []string{
		string(marketanalyzer.PE),
		string(marketanalyzer.PS),
	}

	cmd := &cobra.Command{
		Use:   "web",
		Short: "todo",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			emitter := args[0]
			report := prepareReport(emitter, periodFlag, hum.NumbersMode.String(), precisionFlag)

			commonDataPage := "Common data"
			ratioPage := "Ratio data"
			commonDataUrl := renderChart(commonDataPage, report, ratiosForChart, emitter, false)
			ratioUrl := renderChart(ratioPage, report, commonValueForChart, emitter, true)

			f, _ := os.Create("data/saved_charts/index.html")

			links := []Link{
				{
					Url:   commonDataUrl,
					Label: commonDataPage,
				},
				{
					Url:   ratioUrl,
					Label: ratioPage,
				},
			}

			indexTemplate, err := template.ParseFiles("data/templates/index.tmpl")
			if err != nil {
				log.Fatalln(err)
			}
			if err := indexTemplate.Execute(f, links); err != nil {
				log.Fatalln(err)
			}
		},
	}

	cmd.Flags().IntVarP(&precisionFlag, "precision", "", 2, "Sets the number of digits after the dot for percentages")
	cmd.Flags().StringVarP(&periodFlag, "period", "", marketanalyzer.NormalMode.String(), "Sets the period mode")

	return cmd
}

type Link struct {
	Url   string
	Label string
}

func renderChart(pageTitle string, report *hum.ReadyData, ratiosForChart []string, emitter string, isFloat bool) string {
	chart := charts.NewLine()
	chart.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:  "Home",
				Link:   "index.html",
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
		if tools.ContainStringInSlice(row[0], ratiosForChart) {
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

	filename := strings.ToLower(pageTitle)
	filename = strings.Replace(filename, " ", "-", -1)
	filename = fmt.Sprintf("%s-%s.html", emitter, filename)

	path := fmt.Sprintf("data/saved_charts/%s", filename)
	f, _ := os.Create(path)
	if err := chart.Render(f); err != nil {
		log.Fatalln(err)
	}
	return filename
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
