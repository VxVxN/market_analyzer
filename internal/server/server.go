package server

import (
	"fmt"
	"text/template"

	"github.com/VxVxN/log"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine

	indexTemplate   *template.Template
	emitterTemplate *template.Template
	noteTemplate    *template.Template
}

func Init() (*Server, error) {
	gin.SetMode(gin.ReleaseMode)
	server := Server{router: gin.Default()}

	var err error
	if err = log.Init("market_analyzer.log", log.CommonLog, false); err != nil {
		return nil, fmt.Errorf("cannot init log: %v", err)
	}

	server.router.Static("/static/", "web/")

	server.indexTemplate, err = template.ParseFiles("web/templates/index.tmpl")
	if err != nil {
		return nil, fmt.Errorf("cannot init index template: %v", err)
	}

	server.emitterTemplate, err = template.ParseFiles("web/templates/emitter.tmpl")
	if err != nil {
		return nil, fmt.Errorf("cannot init emitter template: %v", err)
	}

	server.noteTemplate, err = template.ParseFiles("web/templates/note.tmpl")
	if err != nil {
		return nil, fmt.Errorf("cannot init note template: %v", err)
	}

	return &server, nil
}

func (server *Server) ListenAndServe(listen string) error {
	log.Info.Printf("Listening %s", listen)
	if err := server.router.Run(listen); err != nil {
		return err
	}
	return nil
}

func (server *Server) SetRoutes() {
	server.router.GET("/", server.indexHandler)
	server.router.GET("/emitter/:group/:name", server.emitterHandler)
	server.router.GET("/emitter/:group/:name/common-data", server.commonDataHandler)
	server.router.GET("/emitter/:group/:name/ratio-data", server.ratioDataHandler)

	server.router.GET("/emitter/:group/:name/note", server.noteHandler)
	server.router.POST("/emitter/:group/:name/note/save", server.noteSaveHandler)
}
