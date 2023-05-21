package server

import (
	"github.com/VxVxN/market_analyzer/internal/consts"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"

	e "github.com/VxVxN/market_analyzer/pkg/error"
)

type NoteData struct {
	Group string
	Name  string
	Text  string
}

func (server *Server) noteHandler(c *gin.Context) {
	group := c.Param("group")
	emitter := c.Param("name")

	noteText, _ := os.ReadFile(path.Join("data", "emitters", group, emitter+consts.TxtFileExtension))

	data := NoteData{
		Group: group,
		Name:  emitter,
		Text:  string(noteText),
	}

	if err := server.noteTemplate.Execute(c.Writer, data); err != nil {
		e.NewError("Failed to execute note template", http.StatusInternalServerError, nil).JsonResponse(c)
		return
	}
}
