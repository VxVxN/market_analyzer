package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	e "github.com/VxVxN/market_analyzer/pkg/error"
)

type NoteData struct {
	Name string
	Text string
}

func (server *Server) noteHandler(c *gin.Context) {
	emitter := c.Param("name")

	noteText, _ := os.ReadFile("data/emitters/" + emitter + ".txt")

	data := NoteData{
		Name: emitter,
		Text: string(noteText),
	}

	if err := server.noteTemplate.Execute(c.Writer, data); err != nil {
		e.NewError("Failed to execute note template", http.StatusInternalServerError, nil).JsonResponse(c)
		return
	}
}
