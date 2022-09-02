package server

import (
	"net/http"
	"os"

	"github.com/VxVxN/log"
	"github.com/gin-gonic/gin"

	e "github.com/VxVxN/market_analyzer/pkg/error"
	"github.com/VxVxN/market_analyzer/pkg/tools"
)

type RequestNoteSave struct {
	Text string `json:"text" binding:"required"`
}

func (server *Server) noteSaveHandler(c *gin.Context) {
	emitter := c.Param("name")

	var req RequestNoteSave
	if errObj := tools.UnmarshalRequest(c, &req); errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	flags := os.O_TRUNC | os.O_WRONLY | os.O_CREATE

	file, err := os.OpenFile("data/emitters/"+emitter+".txt", flags, 0644)
	if err != nil {
		e.NewError("Failed to create note file", http.StatusInternalServerError, nil).JsonResponse(c)
		return
	}
	defer tools.Close(file, "Failed to close file when save note")

	if _, err = file.WriteString(req.Text); err != nil {
		e.NewError("Failed to write note to file", http.StatusInternalServerError, nil).JsonResponse(c)
		return
	}
}
