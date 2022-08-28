package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	hum "github.com/VxVxN/market_analyzer/internal/humanizer"
	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
	e "github.com/VxVxN/market_analyzer/pkg/error"
)

func (server *Server) ratioDataHandler(c *gin.Context) {
	emitter := c.Param("name")

	report, err := prepareReport(emitter, marketanalyzer.NormalMode, hum.NumbersMode, 2)
	if err != nil {
		e.NewError("Failed to prepare report", http.StatusInternalServerError, nil).JsonResponse(c)
		return
	}

	ratiosForChart := []string{
		string(marketanalyzer.PE),
		string(marketanalyzer.PS),
	}

	if err = renderChart(c.Writer, "Ratio data", emitter, report, ratiosForChart, true); err != nil {
		e.NewError("Failed to redner chart", http.StatusInternalServerError, nil).JsonResponse(c)
		return
	}
}
