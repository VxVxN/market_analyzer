package server

import (
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/market_analyzer/internal/consts"
	e "github.com/VxVxN/market_analyzer/pkg/error"
)

func (server *Server) emittersListHandler(c *gin.Context) {
	var emitters []string
	err := filepath.WalkDir("data/emitters", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		emitter := strings.Replace(filepath.Base(path), consts.CsvFileExtension, "", 1)
		emitters = append(emitters, emitter)
		return nil
	})
	if err != nil {
		e.NewError("Failed to walk directories", http.StatusInternalServerError, nil).JsonResponse(c)
		return
	}
	c.JSON(200, gin.H{
		"emitters": emitters,
	})
}
