package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/market_analyzer/internal/consts"
	hum "github.com/VxVxN/market_analyzer/internal/humanizer"
	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
	"github.com/VxVxN/market_analyzer/internal/parser/myfileparser"
	preparerpkg "github.com/VxVxN/market_analyzer/internal/preparer"
	e "github.com/VxVxN/market_analyzer/pkg/error"
	"github.com/VxVxN/market_analyzer/pkg/tools"
)

type emittersReportRequest struct {
	Emitter    string                    `json:"emitter"`
	PeriodMode marketanalyzer.PeriodMode `json:"period_mode"`
	NumberMode hum.NumberMode            `json:"number_mode"`
	Precision  int                       `json:"precision"`
}

func (server *Server) emittersReportHandler(c *gin.Context) {
	req := emittersReportRequest{
		PeriodMode: marketanalyzer.NormalMode,
		NumberMode: hum.NumbersWithPercentagesMode,
	}
	errObj := tools.UnmarshalRequest(c, &req)
	if errObj != nil {
		errObj.JsonResponse(c)
		return
	}

	parser := myfileparser.Init("data/emitters/" + req.Emitter + consts.CsvFileExtension)

	rawData, err := parser.Parse()
	if err != nil {
		e.NewError("Failed to parse file", http.StatusInternalServerError, err).JsonResponse(c)
		return
	}

	preparer := preparerpkg.Init(rawData)
	rawMarketData, err := preparer.Prepare()
	if err != nil {
		e.NewError("Failed to prepare data", http.StatusInternalServerError, err).JsonResponse(c)
		return
	}

	analyzer := marketanalyzer.Init(rawMarketData)
	analyzer.SetPeriodMode(req.PeriodMode)
	marketData, err := analyzer.Calculate()
	if err != nil {
		e.NewError("Failed to calculate market model", http.StatusInternalServerError, err).JsonResponse(c)
		return
	}

	humanizer := hum.Init(marketData)
	humanizer.SetPrecision(req.Precision)
	humanizer.SetNumbersMode(req.NumberMode)
	humanizer.SetFieldsForDisplay(
		[]marketanalyzer.RowName{
			// marketanalyzer.Sales,
			// marketanalyzer.Earnings,
		},
	)
	data := humanizer.Humanize()

	c.JSON(200, gin.H{
		"headers": data.Headers,
		"rows":    data.Rows,
	})
}
