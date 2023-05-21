package server

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io/fs"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/market_analyzer/internal/consts"
	e "github.com/VxVxN/market_analyzer/pkg/error"
)

type Group struct {
	Name  string
	Links []Link
}

type Link struct {
	Url   string
	Label string
}

func (server *Server) indexHandler(c *gin.Context) {
	emittersByGroup := make(map[string][]string)
	englishTitle := cases.Title(language.English)
	err := filepath.WalkDir("data/emitters", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		group := englishTitle.String(filepath.Base(filepath.Dir(path)))
		emitter := strings.Replace(filepath.Base(path), consts.CsvFileExtension, "", 1)
		if strings.HasSuffix(emitter, ".txt") {
			return nil
		}
		emittersByGroup[group] = append(emittersByGroup[group], emitter)
		return nil
	})
	if err != nil {
		e.NewError("Failed to walk directories", http.StatusInternalServerError, err).JsonResponse(c)
		return
	}

	groups := prepareGroups(emittersByGroup)

	if err = server.indexTemplate.Execute(c.Writer, groups); err != nil {
		e.NewError("Failed to execute index template", http.StatusInternalServerError, err).JsonResponse(c)
		return
	}
}

func prepareGroups(emittersByGroup map[string][]string) []Group {
	var groups []Group
	for group, links := range emittersByGroup {
		var resultLinks []Link
		for _, name := range links {
			resultLinks = append(resultLinks, Link{
				Url:   path.Join("emitter", strings.ToLower(group), strings.ToLower(name)),
				Label: name,
			})
		}
		groups = append(groups, Group{
			Name:  group,
			Links: resultLinks,
		})
	}
	return groups
}
