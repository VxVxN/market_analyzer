package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	hum "github.com/VxVxN/market_analyzer/internal/humanizer"
	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
	e "github.com/VxVxN/market_analyzer/pkg/error"
)

func (server *Server) ratioDataHandler(c *gin.Context) {
	group := c.Param("group")
	emitter := c.Param("name")

	report, err := prepareReport(group, emitter, marketanalyzer.NormalMode, hum.NumbersMode, 2)
	if err != nil {
		e.NewError("Failed to prepare report", http.StatusInternalServerError, nil).JsonResponse(c)
		return
	}

	ratiosForChart := []string{
		string(marketanalyzer.PE),
		string(marketanalyzer.PS),
	}

	_, err = c.Writer.WriteString(fmt.Sprintf("<p><a href=\"/emitter/%s/%s\">Back</a></p>", group, emitter))
	if err != nil {
		e.NewError("Failed to write string to writer", http.StatusInternalServerError, err).JsonResponse(c)
		return
	}

	if err = renderChart(c.Writer, "Ratio data", "Ratio chart", report, ratiosForChart, true); err != nil {
		e.NewError("Failed to redner chart", http.StatusInternalServerError, nil).JsonResponse(c)
		return
	}
}
