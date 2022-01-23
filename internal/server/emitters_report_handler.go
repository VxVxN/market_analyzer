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

/**
 * @api {post} /emitters/report Return report about emitter
 * @apiName emittersReportHandler
 * @apiGroup emitters
 *
 * @apiParam {String} emitter Mandatory emitter name
 * @apiParam {String="normal","first","second","third","four",
"year"} period_mode="normal" Optional period mod that indicates for which period the report will be displayed.
 * @apiParam {String="num_with_percent","number",
"percent"} number_mode="num_with_percent" Optional number mod that specifies which format to output numbers in.
 * @apiParam {Number} precision Optional field specifies how many decimal places will be displayed in the report.
 *
 * @apiParamExample {json} Request-Example:
 *     {
 *       "emitter": "yndx,
 *       "period_mode": "normal",
 *       "number_mode": "num_with_percent",
 *       "precision": 3
 *     }
 *
 * @apiSuccessExample {json} Success-Response:
 *	HTTP/1.1 200 OK
 *	{
 *		"headers": [
 *			"#",
 *			"2016/4",
 *			"2017/4",
 *			"2018/4",
 *			"2019/4",
 *			"2020/2",
 *			"2020/3",
 *			"2020/4",
 *			"2021/1",
 *			"2021/2",
 *			"2021/3"
 *		],
 *		"rows": [
 *			[
 *				"market_cap",
 *				"-",
 *				"66.000.000.000",
 *				"66.000.000.000(+0%)",
 *				"66.000.000.000(+0%)",
 *				"-",
 *				"-",
 *				"66.000.000.000(+0%)",
 *				"66.000.000.000(+0%)",
 *				"66.000.000.000(+0%)",
 *				"66.000.000.000(+0%)"
 *			],
 *			[
 *				"sales",
 *				"-",
 *				"-",
 *				"-",
 *				"-",
 *				"1.600.000.000",
 *				"990.000.000(-38%)",
 *				"3.930.000.000(+297%)",
 *				"-",
 *				"1.970.000.000(-50%)",
 *				"1.290.000.000(-35%)"
 *			],
 *			[
 *				"earnings",
 *				"-",
 *				"-",
 *				"-",
 *				"-",
 *				"-162.000.000",
 *				"256.000.000(+258%)",
 *				"1.660.000.000(+548%)",
 *				"-",
 *				"-180.000.000(-111%)",
 *				"245.000.000(+236%)"
 *			],
 *			[
 *				"debts",
 *				"-",
 *				"-",
 *				"1.480.000.000",
 *				"1.320.000.000(-11%)",
 *				"-",
 *				"-",
 *				"1.400.000.000(+6%)",
 *				"-",
 *				"-",
 *				"2.700.000.000(+93%)"
 *			],
 *			[
 *				"p/e",
 *				"-",
 *				"-",
 *				"-",
 *				"-",
 *				"-",
 *				"-",
 *				"40",
 *				"-",
 *				"-367",
 *				"269"
 *			],
 *			[
 *				"p/s",
 *				"-",
 *				"-",
 *				"-",
 *				"-",
 *				"-",
 *				"-",
 *				"17",
 *				"-",
 *				"34",
 *				"51"
 *			]
 *		]
 *	}
 *
 * @apiErrorExample Error-Response:
 *		HTTP/1.1 500 Internal Server Error
 *		{
 *			"message":"Failed to parse file"
 *		}
*/

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
