package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	e "github.com/VxVxN/market_analyzer/pkg/error"
)

type EmitterData struct {
	Name  string
	Links []Link
}

func (server *Server) emitterHandler(c *gin.Context) {
	emitter := c.Param("name")

	data := EmitterData{
		Name: emitter,
		Links: []Link{
			{
				Url:   "/emitter/" + emitter + "/common-data",
				Label: "Common data",
			},
			{
				Url:   "/emitter/" + emitter + "/ratio-data",
				Label: "Ratio data",
			},
			{
				Url:   "/emitter/" + emitter + "/note",
				Label: "Note",
			},
		},
	}

	if err := server.emitterTemplate.Execute(c.Writer, data); err != nil {
		e.NewError("Failed to execute emitter template", http.StatusInternalServerError, nil).JsonResponse(c)
		return
	}
}
