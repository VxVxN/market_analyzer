package server

import (
	"io"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/market_analyzer/internal/parser/smartlabparser"
	"github.com/VxVxN/market_analyzer/internal/saver/csvsaver"
	e "github.com/VxVxN/market_analyzer/pkg/error"
	"github.com/VxVxN/market_analyzer/pkg/tools"
)

func (server *Server) emittersLoadHandler(c *gin.Context) {

	fileHeader, err := c.FormFile("emitter")
	if err != nil {
		e.NewError("Failed to get file from form", http.StatusInternalServerError, err).JsonResponse(c)
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		e.NewError("Failed to open file", http.StatusInternalServerError, err).JsonResponse(c)
		return
	}
	defer tools.Close(file, "can't close file, when load emitter")

	data, err := io.ReadAll(file)
	if err != nil {
		e.NewError("Failed to read file", http.StatusInternalServerError, err).JsonResponse(c)
		return
	}

	parser := smartlabparser.Init()
	parser.SetFileData(data)

	parsedData, err := parser.Parse()
	if err != nil {
		e.NewError("Failed to parse data", http.StatusInternalServerError, err).JsonResponse(c)
		return
	}

	saver := csvsaver.Init(path.Join("data/emitters", fileHeader.Filename), parsedData.Headers, parsedData.Rows)
	if err = saver.Save(); err != nil {
		e.NewError("Failed to save emitter", http.StatusInternalServerError, err).JsonResponse(c)
		return
	}

	c.JSON(200, gin.H{})
}
