package server

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"

	e "github.com/VxVxN/market_analyzer/pkg/error"
)

type EmitterData struct {
	Name  string
	Links []Link
}

func (server *Server) emitterHandler(c *gin.Context) {
	group := c.Param("group")
	emitter := c.Param("name")

	templatePath := path.Join("/emitter", strings.ToLower(group), emitter, "%s")

	data := EmitterData{
		Name: emitter,
		Links: []Link{
			{
				Url:   fmt.Sprintf(templatePath, "common-data"),
				Label: "Common data",
			},
			{
				Url:   fmt.Sprintf(templatePath, "ratio-data"),
				Label: "Ratio data",
			},
			{
				Url:   fmt.Sprintf(templatePath, "note"),
				Label: "Note",
			},
		},
	}

	if err := server.emitterTemplate.Execute(c.Writer, data); err != nil {
		e.NewError("Failed to execute emitter template", http.StatusInternalServerError, nil).JsonResponse(c)
		return
	}
}
